package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/oklog/run"
)

func main() {
	var g run.Group
	{
		cancel := make(chan struct{})
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
		ln, _ := net.Listen("tcp", ":0")
		g.Add(func() error {
			defer fmt.Printf("http.Serve returned\n")
			return http.Serve(ln, http.NewServeMux())
		}, func(err error) {
			ln.Close()
			fmt.Printf("The third actor was interrupted with: %v\n", err)
		})
	}
	fmt.Printf("The group was terminated with: %v\n", g.Run())
}
