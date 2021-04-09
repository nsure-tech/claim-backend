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

type RewardExecutor struct {
	rewardCh    chan int64
	rewardChaCh chan int64
}

func (r *RewardExecutor) Start(ctx context.Context, wg *sync.WaitGroup) {
	go r.runRewardInspector(ctx)
	go r.runRewardProcess(wg)
	//go r.runRewardChaInspector(ctx)
	//go r.runRewardChaProcess(wg)
}

func NewRewardExecutor() *RewardExecutor {
	return &RewardExecutor{
		rewardCh:    make(chan int64),
		rewardChaCh: make(chan int64),
	}
}

func (r *RewardExecutor) runRewardProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending reward runProcess end.")

	for {
		select {
		case rewardId, ok := <-r.rewardCh:
			if ok {
				err := service.ExecuteReward(rewardId)
				if err != nil {
					log.GetLog().Error("service ExecuteReward error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("reward runProcess select end.")
}

func (r *RewardExecutor) runRewardInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.RewardInspectorTime * time.Second):
			rewards, err := service.GetRewardByEnd(common.RewardMinute)
			if err != nil {
				log.GetLog().Error("service GetRewardByEnd error!", zap.Error(err))
				continue
			}
			for _, reward := range rewards {
				r.rewardCh <- reward.Id
			}
		case <-ctx.Done():
			close(r.rewardCh)
			log.GetLog().Warn("close reward workerCh")
			return
		}
	}
}

func (r *RewardExecutor) runRewardChaProcess(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer log.GetLog().Warn("pending reward runProcess end.")

	for {
		select {
		case rewardChaId, ok := <-r.rewardChaCh:
			if ok {
				err := service.ExecuteRewardCha(rewardChaId)
				if err != nil {
					log.GetLog().Error("service ExecuteRewardCha error!", zap.Error(err))
				}
			} else {
				goto endFor
			}

		}
	}
endFor:
	log.GetLog().Debug("rewardCha runProcess select end.")
}

func (r *RewardExecutor) runRewardChaInspector(ctx context.Context) {
	for {
		select {
		case <-time.After(common.RewardChaInspectorTime * time.Second):
			rewardChas, err := service.GetRewardChaByEnd(common.RewardChaMinute)
			if err != nil {
				log.GetLog().Error("service GetRewardChaByEnd error!", zap.Error(err))
				continue
			}
			for _, rewardCha := range rewardChas {
				r.rewardChaCh <- rewardCha.Id
			}
		case <-ctx.Done():
			close(r.rewardChaCh)
			log.GetLog().Warn("close rewardCha workerCh")
			return
		}
	}
}
