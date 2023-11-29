package exp

import (
	"context"
	"math/rand"
	"time"
)

// Retry creates a channel on which the current time is sent up to count times.
// The first time is sent immediately, with the nth subsequent time sent after
// waiting a random duration D ~ Uniform[0, initial*grow^n), i.e the upper bound
// of wait duration between sends is the geometric sequence:
// initial, initial*grow, initial*grow^2, ...
// Because the channel is unbuffered, the overall duration between consecutive
// sends can be greater than the time waited (by the time taken by the
// receive side), but the overall duration is not necessarily the sum of the
// Retry wait and the work. For example,
//
//	  for range Retry(ctx, 5, 1*time.Second, 1.0) {
//		     // nothing
//	  }
//
// will take around 2 seconds (sum of 4 random durations from U[0, 1)), but
//
//	  for range Retry(ctx, 5, 1*time.Second, 1.0) {
//		     time.Sleep(2*time.Second)
//	  }
//
// will take 10 seconds (Retry tries to send during the sleep in the loop).
//
// To ensure resources (internal goroutines) are cleaned up, cancel the context.
// The channel is closed when the internal goroutine returns.
func Retry(ctx context.Context, count int, initial time.Duration, grow float64) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		defer close(ch)

		select {
		case ch <- time.Now():
		case <-ctx.Done():
			return
		}

		df := float64(initial)
		for n := 1; n < count; n++ {
			d := time.Duration(rand.Int63n(int64(df)))
			timer := time.NewTimer(d)
			select {
			case t := <-timer.C:
				select {
				case ch <- t:

				case <-ctx.Done():
					timer.Stop()
					return
				}

			case <-ctx.Done():
				timer.Stop()
				return
			}

			df *= grow
		}
	}()
	return ch
}
