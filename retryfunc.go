package exp

import (
	"context"
	"iter"
	"math/rand"
	"time"
)

// Retry is like RetryChan, but returns a Go 1.23
// iterator and does not use an extra goroutine under the hood.
// With the rangefunc experiment enabled, you can write:
//
//	for t := range exp.Retry(ctx, 5, 1*time.Second, 2.0) {
//		// try to do thing here
//	}
//
// Without Go 1.23, you can exercise it manually:
//
//	exp.Retry(ctx, 5, 1*time.Second, 2.0)(func(t time.Time) bool {
//		// try to do thing here
//		// return true to keep retrying
//		return true
//	})
//
// Because there is no underlying goroutine, to match the behaviour of original-
// flavour Retry, Retry measures how long `yield` takes and subtracts it from
// the next wait. Because `yield` can take arbitrarily long, it can therefore be
// called again immediately.
func Retry(ctx context.Context, count int, initial time.Duration, grow float64) iter.Seq[time.Time] {
	return func(yield func(time.Time) bool) {
		if err := ctx.Err(); err != nil {
			return
		}

		t := time.Now()
		if !yield(t) {
			return
		}
		yieldDur := time.Since(t)

		df := float64(initial)
		for n := 1; n < count; n++ {
			d := time.Duration(rand.Int63n(int64(df))) - yieldDur
			timer := time.NewTimer(d)
			select {
			case t := <-timer.C:
				if !yield(t) {
					return
				}
				yieldDur = time.Since(t)

			case <-ctx.Done():
				timer.Stop()
				return
			}

			df *= grow
		}
	}
}
