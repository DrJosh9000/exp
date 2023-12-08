// Package para implements parallel versions of some things in algo.
package para

import (
	"runtime"
	"sync"
)

// Map calls f with each element of in, to build the output slice.
// It does the mapping in parallel with GOMAXPROCS goroutines.
func Map[S ~[]X, X, Y any](in S, f func(X) Y) []Y {
	out := make([]Y, len(in))

	N := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		i := i
		go func() {
			defer wg.Done()

			for j := i; j < len(in); j += N {
				out[j] = f(in[j])
			}
		}()
	}

	wg.Wait()
	return out
}
