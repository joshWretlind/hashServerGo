package utils

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

var ShutdownQueue = make(chan int, 1)

func StartUpShutdownWorkers(httpServer http.Server) {
	go shutdownWorker(ShutdownQueue, httpServer)
}

func shutdownWorker(shutdownQueue chan int, httpServer http.Server) {
	for range shutdownQueue {
		close(ShutdownQueue)
		JobWaitGroup.Wait()
		close(HashingJobQueue)
		var shutdownContext,_ = context.WithTimeout(context.Background(), 5*time.Minute)
		httpServer.Shutdown(shutdownContext)
		fmt.Println("Server shutdown")
		os.Exit(0)
	}
}
