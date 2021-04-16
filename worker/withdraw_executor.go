package worker

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"nsure/vote/common"
	"nsure/vote/config"
	"nsure/vote/contract"
	"nsure/vote/log"
	"nsure/vote/service"
	"nsure/vote/utils"
	"sync"
	"time"
)

type WithdrawExecutor struct {
	workerCh     chan uint64
	client       *ethclient.Client
	nSure        *contract.Contract
	start        uint64
	contrAddress string
	chainStatus  bool
}

func (w *WithdrawExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	if blockNumber, err := w.client.BlockNumber(context.Background()); err != nil {
		log.GetLog().Error("BlockNumber error!", zap.Error(err))
	} else {
		if err = w.filterAndStore(blockNumber); err != nil {
			log.GetLog().Error("filterAndStore error!", zap.Error(err))
		}
	}

	go w.runInspector(ctx)
	go w.runProcess(wg)
	go w.runWithdrawBack(ctx)
}

func NewWithdrawExecutor() *WithdrawExecutor {
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.RawUrl)
	if err != nil {
		log.GetLog().Error("eth client dial", zap.Error(err))
	}
	contrAddress := ethcommon.HexToAddress(conf.ContractTreasuryAddress)
	nSureContract, err := contract.NewContract(contrAddress, client)
	if err != nil {
		log.GetLog().Error("contract.NewStore", zap.Error(err))
	}
	start, err := service.GetConfig(common.BlockNumberClaim)
	if err != nil {
		log.GetLog().Error("service.GetConfig", zap.Error(err))
	}
	return &WithdrawExecutor{
		workerCh:     make(chan uint64),
		client:       client,
		nSure:        nSureContract,
		start:        utils.StringToUint64(start),
		contrAddress: contrAddress.String(),
		chainStatus:  false,
	}
}

func (w *WithdrawExecutor) filterAndStore(blockNumber uint64) error {
	end := blockNumber - common.ConfirmBlock
	if w.start >= end {
		return nil
	}
	opts := &bind.FilterOpts{
		Start:   w.start + 1,
		End:     &end,
		Context: context.Background(),
	}

	var from []ethcommon.Address

	claim, err := w.nSure.FilterClaim(opts, from)

	if err != nil {
		log.GetLog().Error("FilterClaim", zap.Error(err))
		return err
	}

	for b := claim.Next(); b; b = claim.Next() {
		log.GetLog().Debug("claim withdraw", zap.String("user", claim.Event.User.String()))
		log.GetLog().Debug("claim withdraw", zap.Any("currency", claim.Event.Currency))
		log.GetLog().Debug("claim withdraw", zap.Any("amount", claim.Event.Amount))
		log.GetLog().Debug("claim withdraw", zap.Any("nonce", claim.Event.Nonce))
		log.GetLog().Debug("claim withdraw", zap.Any("raw", claim.Event.Raw))
		raw, _ := json.Marshal(claim.Event.Raw)
		amount := decimal.NewFromBigInt(claim.Event.Amount, 0)
		nonce := claim.Event.Nonce.Uint64()

		err = service.AddChainClaim(claim.Event.User.String(), claim.Event.Currency.String(),
			amount, nonce, string(raw))
		if err != nil {
			log.GetLog().Error("service AddChainClaim", zap.Error(err))
		}
	}

	w.start = end
	service.UpdateConfig(common.BlockNumberClaim, utils.U64ToA(w.start))
	return nil
}

func (w *WithdrawExecutor) runProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case blockNumber, ok := <-w.workerCh:
			if ok {
				err := w.filterAndStore(blockNumber)
				if err != nil {
					log.GetLog().Error("withdraw filterAndStore error!", zap.Error(err))
				}
			} else {

				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("withdraw runProcess select end.")
}

func (w *WithdrawExecutor) runInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.WithdrawInspectorTime * time.Second):
			blockNumber, err := w.client.BlockNumber(context.Background())
			if err != nil {
				w.chainStatus = false
				log.GetLog().Error("BlockNumber error!", zap.Error(err))
				continue
			}
			w.chainStatus = true
			w.workerCh <- blockNumber

		case <-ctx.Done():
			close(w.workerCh)
			log.GetLog().Warn("close withdraw workerCh")
			return
		}
	}
}

func (w *WithdrawExecutor) runWithdrawBack(ctx context.Context) {
	chainStatues := false
	for {
		select {
		case <-time.After(common.WithdrawBackInspectorTime * time.Second):
			if !w.chainStatus {
				chainStatues = false
				continue
			}
			if !chainStatues {
				chainStatues = true
				time.Sleep(time.Duration(common.WithdrawChainMinute) * time.Minute)
			}
			withdraws, err := service.GetWithdrawsByEndAt(common.WithdrawBackMinute)
			if err != nil {
				log.GetLog().Error("service GetWithdrawsByEndAt", zap.Error(err))
				continue
			}
			for _, withdraw := range withdraws {
				err = service.ExecuteWithdraw(withdraw)
				if err != nil {
					log.GetLog().Error("service ExecuteWithdraw", zap.Error(err))
				}
			}

		case <-ctx.Done():
			log.GetLog().Warn("close withdraw workerCh")
			return
		}
	}
}
