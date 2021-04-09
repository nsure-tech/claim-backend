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

type PendingExecutor struct {
	workerCh chan *models.Pending
}

func (p *PendingExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	go p.runInspector(ctx)
	go p.runProcess(wg)
}

func NewPendingExecutor() *PendingExecutor {
	return &PendingExecutor{
		workerCh: make(chan *models.Pending, common.PendingWorkChanNum),
	}
}

func (p *PendingExecutor) runProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case pending, ok := <-p.workerCh:
			if ok {
				err := service.ExecutePending(pending)
				if err != nil {
					log.GetLog().Error("service ExecutePending error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("pending runProcess select end.")
}

func (p *PendingExecutor) runInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.PendingInspectorTime * time.Second):
			pends, err := service.GetUnsettledPends(common.PendingWorkUnsettleCount)
			if err != nil {
				log.GetLog().Error("service GetUnsettledPends error!", zap.Error(err))
				continue
			}
			for _, pend := range pends {
				p.workerCh <- pend
			}
		case <-ctx.Done():
			close(p.workerCh)
			log.GetLog().Warn("close pending workerCh")
			return
		}
	}
}
