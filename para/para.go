// Package para implements parallel versions of some things in algo.
package para

import (
	"runtime"
	"sync"

	"drjosh.dev/exp/algo"
)

// Do calls f with each element of in.
// It does this in parallel, using up to GOMAXPROCS goroutines.
func Do[S ~[]E, E any](in S, f func(E)) {
	if len(in) == 0 {
		return
	}

	var wg sync.WaitGroup

	for _, chunk := range Divvy(in, runtime.GOMAXPROCS(0)) {
		wg.Add(1)
		chunk := chunk

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
	if len(in) == 0 {
		return []Y{}
	}

	N := runtime.GOMAXPROCS(0)
	inchunks := Divvy(in, N)
	out := make([]Y, len(in))
	outchunks := Divvy(out, N)

	var wg sync.WaitGroup

	for i := range inchunks {
		wg.Add(1)
		inchunk, outchunk := inchunks[i], outchunks[i]

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
	if len(in) == 0 {
		var zero E
		return zero
	}

	chunks := Divvy(in, runtime.GOMAXPROCS(0))
	out := make([]E, len(chunks))

	var wg sync.WaitGroup

	for i, chunk := range chunks {
		wg.Add(1)
		i, chunk := i, chunk

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

// Divvy divides a slice into up to n subslices of approximately equal size.
func Divvy[S ~[]E, E any](in S, n int) []S {
	if n == 1 {
		return []S{in}
	}

	n = min(n, len(in))
	cn, rem := len(in)/n, len(in)%n

	out := make([]S, 0, n)
	for i := 0; i < n; i++ {
		offset := i * cn
		count := cn
		if i < rem {
			offset += i
			count++
		} else {
			offset += rem
		}
		out = append(out, in[offset:][:count])
	}
	return out
}
