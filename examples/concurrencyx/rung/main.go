package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/oklog/run"
)

func main() {
	ln, _ := net.Listen("tcp", ":0")
	var g run.Group // watch this video: https://www.youtube.com/watch?v=LHe1Cb_Ud_M&t=15m45s
	{
		cancel := make(chan struct{})
		// Add actors into Group
		// Actors are defined as a pair of functions: an execute function, and an interrupt function
		g.Add(func() error {
			select {
			case <-time.After(time.Second):
				fmt.Printf("The first actor had its time elapsed\n")
				return nil
			case <-cancel:
				fmt.Printf("The first actor was canceled\n")
				return nil
			}
		}, func(err error) {
			fmt.Printf("The first actor was interrupted with: %v\n", err)
			close(cancel)
		})
	}
	{
		g.Add(func() error {
			fmt.Printf("The second actor is returning immediately\n")
			return errors.New("immediate teardown")
		}, func(err error) {
			// Note that this interrupt function is called, even though
			// the corresponding execute function has already returned.
			fmt.Printf("The second actor was interrupted with: %v\n", err)
		})
	}
	{
		g.Add(func() error {
			defer fmt.Printf("http.Serve returned\n")
			return http.Serve(ln, http.NewServeMux())
		}, func(err error) {
			ln.Close()
			fmt.Printf("HTTP Listener actor was interrupted with: %v\n", err)
		})
	}
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				ln.Close()
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(err error) {
			close(cancelInterrupt)
			fmt.Printf("Interruption Listener actor was interrupted with: %v\n", err)

		})
	}
	/*
		Run, which concurrently runs all of the actors, waits until the first actor exits,
		invokes the interrupt functions, and finally returns control to the caller
		only once all actors have returned.
		This general-purpose API allows callers to model pretty much any runnable task,
		and achieve well-defined lifecycle semantics for the group.
	*/
	fmt.Printf("The group was terminated with: %v\n", g.Run())
}
