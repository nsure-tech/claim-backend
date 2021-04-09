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

type ChallengeExecutor struct {
	workerSuccessCh chan int64
	workerFailCh    chan int64
}

func (c *ChallengeExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	go c.runSuccessInspector(ctx)
	go c.runSuccessProcess(wg)
	go c.runFailInspector(ctx)
	go c.runFailProcess(wg)
}

func NewChallengeExecutor() *ChallengeExecutor {
	return &ChallengeExecutor{
		workerSuccessCh: make(chan int64, common.ChallengeWorkChanNum),
		workerFailCh:    make(chan int64, common.ChallengeWorkChanNum),
	}
}

func (c *ChallengeExecutor) runSuccessProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case claimId, ok := <-c.workerSuccessCh:
			if ok {
				err := service.ExecuteChallengeSuccess(claimId)
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

func (c *ChallengeExecutor) runSuccessInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.ChallengeInspectorTime * time.Second):
			challenges, err := service.GetChallengeByStatus(common.ChallengeStatusSuccess, common.ChallengeWorkUnsettleCount)
			if err != nil {
				log.GetLog().Error("service GetChallengeByStatus error!", zap.Error(err))
				continue
			}
			for _, challenge := range challenges {
				c.workerSuccessCh <- challenge.ClaimId
			}
		case <-ctx.Done():
			close(c.workerSuccessCh)
			log.GetLog().Warn("close challenge workerSuccessCh")
			return
		}
	}
}

func (c *ChallengeExecutor) runFailProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending runProcess end.")

	for {
		select {
		case claimId, ok := <-c.workerFailCh:
			if ok {
				err := service.ExecuteChallengeFail(claimId)
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

func (c *ChallengeExecutor) runFailInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.ChallengeInspectorTime * time.Second):
			challenges, err := service.GetChallengeByStatus(common.ChallengeStatusFail, common.ChallengeWorkUnsettleCount)
			if err != nil {
				log.GetLog().Error("service GetChallengeByStatus error!", zap.Error(err))
				continue
			}
			for _, challenge := range challenges {
				c.workerFailCh <- challenge.ClaimId
			}
		case <-ctx.Done():
			close(c.workerFailCh)
			log.GetLog().Warn("close challenge workerFailCh")
			return
		}
	}
}
