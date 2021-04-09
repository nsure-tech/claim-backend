package worker

import (
	"context"
	"go.uber.org/zap"
	"nsure/vote/common"
	"nsure/vote/log"
	"nsure/vote/models"
	"nsure/vote/service"
	"sync"
	"time"
)

type BillExecutor struct {
	workerCh chan *models.Bill
}

func (b *BillExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	go b.runInspector(ctx)
	go b.runProcess(wg)
}

func NewBillExecutor() *BillExecutor {
	return &BillExecutor{
		workerCh: make(chan *models.Bill, common.BillWorkChanNum),
	}
}

func (b *BillExecutor) runProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case bill, ok := <-b.workerCh:
			if ok {
				err := service.ExecuteBill(bill)
				if err != nil {
					log.GetLog().Error("service ExecuteBill error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("bill runProcess select end.")
}

func (b *BillExecutor) runInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.BillInspectorTime * time.Second):
			bills, err := service.GetUnsettledBills(common.BillWorkUnsettleCount)
			if err != nil {
				log.GetLog().Error("service GetUnsettledBills error!", zap.Error(err))
				continue
			}
			for _, bill := range bills {
				b.workerCh <- bill
			}
		case <-ctx.Done():
			close(b.workerCh)
			log.GetLog().Warn("close bill workerCh")
			return
		}
	}
}
