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

type BillWaitExecutor struct {
	workerCh chan *models.WaitBill
}

func (b *BillWaitExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	go b.runInspector(ctx)
	go b.runProcess(wg)
}

func NewBillWaitExecutor() *BillWaitExecutor {
	return &BillWaitExecutor{
		workerCh: make(chan *models.WaitBill, common.BillWorkChanNum),
	}
}

func (b *BillWaitExecutor) runProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case bill, ok := <-b.workerCh:
			if ok {
				err := service.ExecuteWaitBill(bill)
				if err != nil {
					log.GetLog().Error("service ExecuteWaitBill error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("bill runProcess select end.")
}

func (b *BillWaitExecutor) runInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.BillWaitInspectorTime * time.Second):
			bills, err := service.GetUnsettledWaitBills(common.BillWorkUnsettleCount)
			if err != nil {
				log.GetLog().Error("service GetUnsettledWaitBills error!", zap.Error(err))
				continue
			}
			for _, bill := range bills {
				b.workerCh <- bill
			}
		case <-ctx.Done():
			close(b.workerCh)
			log.GetLog().Warn("close wait bill workerCh")
			return
		}
	}
}
