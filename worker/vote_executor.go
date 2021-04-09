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

type VoteExecutor struct {
	workerVoteCh  chan int64
	workerClaimCh chan int64
}

func (v *VoteExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	go v.runApplyInspector(ctx)
	go v.runApplyProcess(wg)
	//go v.runClaimInspector(ctx)
	//go v.runClaimProcess(wg)
}

func NewVoteExecutor() *VoteExecutor {
	return &VoteExecutor{
		workerVoteCh:  make(chan int64, common.ApplyWorkChanNum),
		workerClaimCh: make(chan int64, common.ApplyWorkChanNum),
	}
}

func (v *VoteExecutor) runApplyProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case claimId, ok := <-v.workerVoteCh:
			if ok {
				err := service.ExecuteVote(claimId)
				if err != nil {
					log.GetLog().Error("service ExecuteVote error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("apply runProcess select end.")
}

func (v *VoteExecutor) runApplyInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.VoteInspectorTime * time.Second):
			claims, err := service.GetVoteByVoteNum(common.ArbiterMaxNum)
			if err != nil {
				log.GetLog().Error("service GetApplyByApplyNum error!", zap.Error(err))
				continue
			}
			for _, claim := range claims {
				v.workerVoteCh <- claim.ClaimId
			}
		case <-ctx.Done():
			close(v.workerVoteCh)
			log.GetLog().Warn("close vote workerVoteCh")
			return
		}
	}
}

func (v *VoteExecutor) runClaimProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case claimId, ok := <-v.workerClaimCh:
			if ok {
				err := service.ExecuteVote(claimId)
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

func (v *VoteExecutor) runClaimInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.ClaimInspectorTime * time.Second):
			claims, err := service.GetClaimByEndVote()
			if err != nil {
				log.GetLog().Error("service GetClaimByEndVote error!", zap.Error(err))
				continue
			}
			for _, claim := range claims {
				v.workerClaimCh <- claim.Id
			}
		case <-ctx.Done():
			close(v.workerClaimCh)
			log.GetLog().Warn("close vote workerClaimCh")
			return
		}
	}
}
