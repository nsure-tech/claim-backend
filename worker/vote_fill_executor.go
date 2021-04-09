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

type VoteFillExecutor struct {
	workerCh chan int64
}

func (f *VoteFillExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	go f.runInspector(ctx)
	go f.runProcess(wg)
}

func NewVoteFillExecutor() *VoteFillExecutor {
	return &VoteFillExecutor{
		workerCh: make(chan int64, common.VoteFillWorkChanNum),
	}
}

func (f *VoteFillExecutor) runProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case fill, ok := <-f.workerCh:
			if ok {
				err := service.ExecuteVoteFill(fill)
				if err != nil {
					log.GetLog().Error("service ExecuteFill error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("fill runProcess select end.")
}

func (f *VoteFillExecutor) runInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.VoteFillInspectorTime * time.Second):
			statuses := []common.ClaimStatus{common.ClaimStatusPass, common.ClaimStatusDeny}
			claims, err := service.GetClaimClose(statuses)
			if err != nil {
				log.GetLog().Error("service GetUnsettledFills error!", zap.Error(err))
				continue
			}
			for _, claim := range claims {
				if claim.Status == common.ClaimStatusPass || claim.Status == common.ClaimStatusDeny {
					f.workerCh <- claim.Id
				}

			}
		case <-ctx.Done():
			close(f.workerCh)
			log.GetLog().Warn("close fill workerCh")
			return
		}
	}
}
