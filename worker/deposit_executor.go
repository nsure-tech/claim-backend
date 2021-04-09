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

type DepositExecutor struct {
	workerCh     chan uint64
	client       *ethclient.Client
	nSure        *contract.Contract
	start        uint64
	contrAddress string
}

func (d *DepositExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	if blockNumber, err := d.client.BlockNumber(context.Background()); err != nil {
		log.GetLog().Error("BlockNumber error!", zap.Error(err))
	} else {
		err = d.filterAndStore(blockNumber)
	}

	go d.runInspector(ctx)
	go d.runProcess(wg)
}

func NewDepositExecutor() *DepositExecutor {
	conf := config.GetConfig()
	client, err := ethclient.Dial(conf.RawUrl)
	if err != nil {
		log.GetLog().Error("eth client Dial", zap.Error(err))
	}
	contrAddress := ethcommon.HexToAddress(conf.ContractTreasuryAddress)
	nSureContract, err := contract.NewContract(contrAddress, client)
	if err != nil {
		log.GetLog().Error("contract.NewStore", zap.Error(err))
	}
	start, err := service.GetConfig(common.BlockNumberDeposit)
	if err != nil {
		log.GetLog().Error("service.GetConfig", zap.Error(err))
	}
	return &DepositExecutor{
		workerCh:     make(chan uint64),
		client:       client,
		nSure:        nSureContract,
		start:        utils.StringToUint64(start),
		contrAddress: contrAddress.String(),
	}
}

func (d *DepositExecutor) filterAndStore(blockNumber uint64) error {
	end := blockNumber - common.ConfirmBlock
	if d.start >= end {
		return nil
	}
	opts := &bind.FilterOpts{
		Start:   d.start + 1,
		End:     &end,
		Context: context.Background(),
	}

	var from []ethcommon.Address

	deposit, err := d.nSure.FilterDeposit(opts, from)

	if err != nil {
		log.GetLog().Error("FilterDeposit", zap.Error(err))
		return err
	}

	for b := deposit.Next(); b; b = deposit.Next() {
		log.GetLog().Debug("deposit", zap.String("user", deposit.Event.User.String()))
		log.GetLog().Debug("deposit", zap.Any("amount", deposit.Event.Amount))
		log.GetLog().Debug("deposit", zap.Any("raw", deposit.Event.Raw))
		raw, _ := json.Marshal(deposit.Event.Raw)
		amount := decimal.NewFromBigInt(deposit.Event.Amount, 0)

		err := service.DepositTreasury(common.CurrencyNSure, deposit.Event.User.String(), d.contrAddress,
			string(raw), common.TransferStatusDeposit, amount)
		if err != nil {
			log.GetLog().Error("deposit DepositTreasury", zap.Error(err))
		}
	}

	d.start = end
	service.UpdateConfig(common.BlockNumberDeposit, utils.U64ToA(d.start))
	return nil
}

func (d *DepositExecutor) runProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case blockNumber, ok := <-d.workerCh:
			if ok {
				err := d.filterAndStore(blockNumber)
				if err != nil {
					log.GetLog().Error("deposit filterAndStore error!", zap.Error(err))
				}
			} else {

				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("deposit runProcess select end.")
}

func (d *DepositExecutor) runInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.DepositInspectorTime * time.Second):
			blockNumber, err := d.client.BlockNumber(context.Background())
			if err != nil {
				log.GetLog().Error("BlockNumber error!", zap.Error(err))
				continue
			}
			d.workerCh <- blockNumber

		case <-ctx.Done():
			close(d.workerCh)
			log.GetLog().Warn("close deposit workerCh")
			return
		}
	}
}
