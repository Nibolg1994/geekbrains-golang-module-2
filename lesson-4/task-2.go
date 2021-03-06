package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	doneCh := make(chan error)
	var ctx, cancel = context.WithCancel(context.Background())

	go func(ctx context.Context, cancel context.CancelFunc) {
		<-sigs
		var secondCtx, secondCancel = context.WithTimeout(ctx, 1*time.Second)
		select {
		case <-secondCtx.Done():
			cancel()
			secondCancel()
		}
	}(ctx, cancel)

	go background(doneCh)

	var err error
	select {
	case <-ctx.Done():
		err = ctx.Err()
		fmt.Println(err)
	case err = <-doneCh:
		fmt.Println("func done", err)
	}
}

func background(ch chan error) {
	for i := 1; i < 10; i++ {
		fmt.Println("... processing ...")
		time.Sleep(1 * time.Second)
	}
	ch <- nil
}
