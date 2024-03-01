package signal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func HandleShutdown(cancel context.CancelFunc) {
	exitCh := make(chan os.Signal, 1)

	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	<-exitCh

	fmt.Println("got termination signal")
	cancel()
}
