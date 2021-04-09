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

type ApplyExecutor struct {
	workerApplyCh chan int64
	workerClaimCh chan int64
}

func (a *ApplyExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	go a.runApplyInspector(ctx)
	go a.runApplyProcess(wg)
	//go a.runClaimInspector(ctx)
	//go a.runClaimProcess(wg)
}

func NewApplyExecutor() *ApplyExecutor {
	return &ApplyExecutor{
		workerApplyCh: make(chan int64, common.ApplyWorkChanNum),
		workerClaimCh: make(chan int64, common.ApplyWorkChanNum),
	}
}

func (a *ApplyExecutor) runApplyProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case claimId, ok := <-a.workerApplyCh:
			if ok {
				err := service.ExecuteApply(claimId)
				if err != nil {
					log.GetLog().Error("service ExecuteApply error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("apply runProcess select end.")
}

func (a *ApplyExecutor) runApplyInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.ApplyInspectorTime * time.Second):
			applies, err := service.GetApplyByApplyNum(common.ApplyMaxNum, nil)
			if err != nil {
				log.GetLog().Error("service GetApplyByApplyNum error!", zap.Error(err))
				continue
			}
			for _, apply := range applies {
				a.workerApplyCh <- apply.ClaimId
			}
		case <-ctx.Done():
			close(a.workerApplyCh)
			log.GetLog().Warn("close apply workerApplyCh")
			return
		}
	}
}

func (a *ApplyExecutor) runClaimProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case claimId, ok := <-a.workerClaimCh:
			if ok {
				err := service.ExecuteClaim(claimId)
				if err != nil {
					log.GetLog().Error("service ExecuteClaim error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("claim runProcess select end.")
}

func (a *ApplyExecutor) runClaimInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.ClaimInspectorTime * time.Second):
			claims, err := service.GetClaimByEndApply()
			if err != nil {
				log.GetLog().Error("service GetClaimByEndApply error!", zap.Error(err))
				continue
			}
			for _, claim := range claims {
				a.workerClaimCh <- claim.Id
			}
		case <-ctx.Done():
			close(a.workerClaimCh)
			log.GetLog().Warn("close apply workerClaimCh")
			return
		}
	}
}
