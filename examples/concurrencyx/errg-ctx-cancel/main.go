package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	ctx, done := context.WithTimeout(context.Background(), 10*time.Second)
	/*
		WithContext returns a new Group and an associated Context derived from ctx.
		The derived Context is canceled the first time a function passed
		to Go returns a non-nil error or the first time Wait returns,
		whichever occurs first.
	*/
	g, gctx := errgroup.WithContext(ctx)
	defer done()
	// goroutine to check for signals to gracefully finish all functions
	g.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

		select {
		case sig := <-signalChannel:
			fmt.Printf("Received signal: %s\n", sig)
			//done()
		case <-gctx.Done():
			fmt.Printf("closing signal goroutine\n")
			return gctx.Err()
		}

		return nil
	})

	// just a ticker every 2s
	g.Go(func() error {
		ticker := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Printf("ticker 2s ticked\n")
				// test error: Uncomment the below code to test returning non-nil error value
				//return fmt.Errorf("test error ticker 2s")
			case <-gctx.Done():
				fmt.Printf("closing ticker 2s goroutine\n")
				return gctx.Err()
			}
		}
	})

	// just a ticker every 1s
	g.Go(func() error {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Printf("ticker 1s ticked\n")
			case <-gctx.Done():
				fmt.Printf("closing ticker 1s goroutine\n")
				return gctx.Err()
			}
		}
	})

	// force a stop after 60s
	//time.AfterFunc(60*time.Second, func() {
	//	fmt.Println("force finished after 60s")
	//	done()
	//})

	// wait for all errg-ctx-cancel goroutines
	err := g.Wait()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			fmt.Println("ctx-cancel was canceled")
		} else {
			fmt.Printf("received error: %v\n", err)
		}
	} else {
		fmt.Println("finished clean")
	}

	//if err := g.Wait(); err != nil {
	//	fmt.Println("from last: ", err)
	//}

}
