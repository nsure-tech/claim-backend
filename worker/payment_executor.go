package worker

import (
	"context"
	"go.uber.org/zap"
	"nsure/vote/common"
	"nsure/vote/log"
	"nsure/vote/service"
	"sync"
	"time"
)

type PaymentExecutor struct {
	workerCh chan int64
}

func (p *PaymentExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	go p.runInspector(ctx)
	go p.runProcess(wg)
}

func NewPaymentExecutor() *PaymentExecutor {
	return &PaymentExecutor{
		workerCh: make(chan int64),
	}
}

func (p *PaymentExecutor) runProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case paymentId, ok := <-p.workerCh:
			if ok {
				err := service.ExecutePayment(paymentId)
				if err != nil {
					log.GetLog().Error("service ExecutePayment error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("claim payment runProcess select end.")
}

func (p *PaymentExecutor) runInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.PaymentInspectorTime * time.Second):
			payments, err := service.GetPaymentByEnd(common.PaymentMinute)
			if err != nil {
				log.GetLog().Error("service GetPaymentByEnd error!", zap.Error(err))
				continue
			}
			for _, payment := range payments {
				p.workerCh <- payment.Id
			}
		case <-ctx.Done():
			close(p.workerCh)
			log.GetLog().Warn("close claim payment workerCh")
			return
		}
	}
}
