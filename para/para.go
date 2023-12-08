// Package para implements parallel versions of some things in algo.
package para

import (
	"runtime"
	"sync"

	"github.com/DrJosh9000/exp/algo"
)

// Map calls f with each element of in, to build the output slice.
// It does the mapping in parallel, using GOMAXPROCS goroutines.
func Map[S ~[]X, X, Y any](in S, f func(X) Y) []Y {
	N := runtime.GOMAXPROCS(0)

	out := make([]Y, len(in))
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

// Reduce reduces a slice to a single value.
// It does the mapping in parallel, using GOMAXPROCS goroutines.
func Reduce[S ~[]E, E any](in S, f func(E, E) E) E {
	N := runtime.GOMAXPROCS(0)
	out := make([]E, N)
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		i := i
		go func() {
			defer wg.Done()

			for j := i; j < len(in); j += N {
				out[i] = f(out[i], in[j])
			}
		}()
	}

	wg.Wait()
	return algo.Foldl(out, f)
}
