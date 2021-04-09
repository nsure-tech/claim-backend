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

type ChaFillExecutor struct {
	workerSuccessCh chan int64
	workerFailCh    chan int64
}

func (c *ChaFillExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	go c.runSuccessInspector(ctx)
	go c.runSuccessProcess(wg)
	go c.runFailInspector(ctx)
	go c.runFailProcess(wg)
}

func NewChaFillExecutor() *ChaFillExecutor {
	return &ChaFillExecutor{
		workerSuccessCh: make(chan int64, common.ChallengeWorkChanNum),
		workerFailCh:    make(chan int64, common.ChallengeWorkChanNum),
	}
}

func (c *ChaFillExecutor) runSuccessProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case claimId, ok := <-c.workerSuccessCh:
			if ok {
				err := service.ExecuteChallengeFillSuccess(claimId)
				if err != nil {
					log.GetLog().Error("service ExecuteChallengeSuccess error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("challenge runSuccessProcess select end.")
}

func (c *ChaFillExecutor) runSuccessInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.ChallengeInspectorTime * time.Second):
			chaFills, err := service.GetChallengeFillByStatus(common.ChallengeStatusSuccess, common.ChallengeWorkUnsettleCount)
			if err != nil {
				log.GetLog().Error("service GetChallengeByStatus error!", zap.Error(err))
				continue
			}
			for _, chaFill := range chaFills {
				c.workerSuccessCh <- chaFill.ClaimId
			}
		case <-ctx.Done():
			close(c.workerSuccessCh)
			log.GetLog().Warn("close challenge workerSuccessCh")
			return
		}
	}
}

func (c *ChaFillExecutor) runFailProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case claimId, ok := <-c.workerFailCh:
			if ok {
				err := service.ExecuteChallengeFillFail(claimId)
				if err != nil {
					log.GetLog().Error("service ExecuteChallengeFail error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("challenges runFailProcess select end.")
}

func (c *ChaFillExecutor) runFailInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.ChallengeInspectorTime * time.Second):
			chaFills, err := service.GetChallengeFillByStatus(common.ChallengeStatusFail, common.ChallengeWorkUnsettleCount)
			if err != nil {
				log.GetLog().Error("service GetChallengeByStatus error!", zap.Error(err))
				continue
			}
			for _, chaFill := range chaFills {
				c.workerFailCh <- chaFill.ClaimId
			}
		case <-ctx.Done():
			close(c.workerFailCh)
			log.GetLog().Warn("close challenge workerFailCh")
			return
		}
	}
}
