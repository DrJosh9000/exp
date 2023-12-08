// Package para implements parallel versions of some things in algo.
package para

import (
	"runtime"
	"sync"

	"github.com/DrJosh9000/exp/algo"
)

// Do calls f with each element of in.
// It does this in parallel, using up to GOMAXPROCS goroutines.
func Do[S ~[]E, E any](in S, f func(E)) {
	N := min(runtime.GOMAXPROCS(0), len(in))
	if N == 0 {
		return
	}
	cn := len(in) / N

	var wg sync.WaitGroup
	wg.Add(N)

	for i := 0; i < N; i++ {
		chunk := in[i*cn:]
		if len(chunk) > cn {
			chunk = chunk[:cn]
		}

		go func() {
			defer wg.Done()

			for _, e := range chunk {
				f(e)
			}
		}()
	}

	wg.Wait()
}

// Map calls f with each element of in, to build the output slice.
// It does the mapping in parallel, using up to GOMAXPROCS goroutines.
func Map[S ~[]X, X, Y any](in S, f func(X) Y) []Y {
	N := min(runtime.GOMAXPROCS(0), len(in))
	if N == 0 {
		return nil
	}
	cn := len(in) / N

	out := make([]Y, len(in))

	var wg sync.WaitGroup
	wg.Add(N)

	for i := 0; i < N; i++ {
		inchunk := in[i*cn:]
		outchunk := out[i*cn:]
		if len(inchunk) > cn {
			inchunk = inchunk[:cn]
		}

		go func() {
			defer wg.Done()

			for j, e := range inchunk {
				outchunk[j] = f(e)
			}
		}()
	}

	wg.Wait()
	return out
}

// Reduce reduces a slice to a single value. For consistent results, f should
// be associative and the zero value for E should be an identity element.
// It does the reduction in parallel, using up to GOMAXPROCS goroutines.
func Reduce[S ~[]E, E any](in S, f func(E, E) E) E {
	N := min(runtime.GOMAXPROCS(0), len(in))
	if N == 0 {
		var zero E
		return zero
	}
	cn := len(in) / N
	out := make([]E, N)
	var wg sync.WaitGroup
	wg.Add(N)

	for i := 0; i < N; i++ {
		i := i
		chunk := in[i*cn:]
		if len(chunk) > cn {
			chunk = chunk[:cn]
		}
		go func() {
			defer wg.Done()

			for _, e := range chunk {
				out[i] = f(out[i], e)
			}
		}()
	}

	wg.Wait()
	return algo.Foldl(out, f)
}
