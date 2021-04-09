package process

import (
	"context"
	"nsure/vote/config"
	"nsure/vote/log"
	"nsure/vote/rest"
	"nsure/vote/worker"
	"sync"
)

func StartEngine(ctx context.Context, wg *sync.WaitGroup) {
	worker.NewBillExecutor().Start(ctx, wg)
	worker.NewPendingExecutor().Start(ctx, wg)

	worker.NewApplyExecutor().Start(ctx, wg)
	worker.NewVoteExecutor().Start(ctx, wg)

	worker.NewChallengeExecutor().Start(ctx, wg)
	worker.NewChaFillExecutor().Start(ctx, wg)

	worker.NewDepositExecutor().Start(ctx, wg)
	worker.NewWithdrawExecutor().Start(ctx, wg)

	worker.NewVoteFillExecutor().Start(ctx, wg)

	worker.NewRewardExecutor().Start(ctx, wg)

	worker.NewPaymentExecutor().Start(ctx, wg)

	config := config.GetConfig()
	httpServer := rest.NewHttpServer(config.RestServer.Addr)
	go httpServer.Start()

	log.GetLog().Info("rest server ok")
}
