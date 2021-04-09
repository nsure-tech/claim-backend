package main

import (
	"context"
	"nsure/vote/config"
	"nsure/vote/log"
	"nsure/vote/process"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	config.GetConfig()
	logger := log.GetLog()

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	process.StartEngine(ctx, wg)

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	// Handle shutdown
	logger.Info("*********************************Shutdown signal received*********************************")
	cancel() // Signal cancellation to context.Context
	logger.Info("*********************************Shutdown wait*********************************")
	wg.Wait() // Block here until are workers are done
	logger.Info("*********************************Shutdown end*********************************")
}
