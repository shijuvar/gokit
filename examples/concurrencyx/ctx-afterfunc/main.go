package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	waitOnCond := func(ctx context.Context, cond *sync.Cond, conditionMet func() bool) error {
		stopf := context.AfterFunc(ctx, func() {
			// We need to acquire cond.L here to be sure that the Broadcast
			// below won't occur before the call to Wait, which would result
			// in a missed signal (and deadlock).
			cond.L.Lock()
			defer cond.L.Unlock()

			// If multiple goroutines are waiting on cond simultaneously,
			// we need to make sure we wake up exactly this one.
			// That means that we need to Broadcast to all of the goroutines,
			// which will wake them all up.
			//
			// If there are N concurrent calls to waitOnCond, each of the goroutines
			// will spuriously wake up O(N) other goroutines that aren't ready yet,
			// so this will cause the overall CPU cost to be O(N²).
			cond.Broadcast()
		})
		defer stopf()

		// Since the wakeups are using Broadcast instead of Signal, this call to
		// Wait may unblock due to some other goroutine's context becoming done,
		// so to be sure that ctx is actually done we need to check it in a loop.
		for !conditionMet() {
			cond.Wait()
			if ctx.Err() != nil {
				return ctx.Err()
			}
		}

		return nil
	}

	cond := sync.NewCond(new(sync.Mutex))

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
			defer cancel()

			cond.L.Lock()
			defer cond.L.Unlock()

			err := waitOnCond(ctx, cond, func() bool { return false })
			fmt.Println(err)
		}()
	}
	wg.Wait()

}
