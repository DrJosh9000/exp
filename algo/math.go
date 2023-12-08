/*
   Copyright 2022 Josh Deprez

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

	   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package algo

import (
	"cmp"

	"golang.org/x/exp/constraints"
)

// Abs returns the absolute value of x (with no regard for negative overflow).
//
// The math/cmplx package provides a version of Abs for complex types.
func Abs[T Real](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// GCD returns the greatest common divisor of a and b.
func GCD[T constraints.Integer](a, b T) T {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// XGCD returns the greatest common divisor of a and b, as well as the Bézout
// coefficients (x and y such that ax + by = GCD(a, b)).
// For arithmetic on large integers, use math/big.
func XGCD[T constraints.Integer](a, b T) (d, x, y T) {
	// Based on https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm.
	or, r := a, b
	os, s := T(1), T(0)
	ot, t := T(0), T(1)
	for r != 0 {
		q := or / r
		or, r = r, or-q*r
		os, s = s, os-q*s
		ot, t = t, ot-q*t
	}
	return or, os, ot
}

// Max returns the greatest argument (using `>`). If no arguments are provided,
// Max returns the zero value for T.
func Max[T cmp.Ordered](x ...T) T {
	var m T
	if len(x) == 0 {
		return m
	}
	m = x[0]
	for _, t := range x[1:] {
		if t > m {
			m = t
		}
	}
	return m
}

// Min returns the least argument (using `<`). If no arguments are provided, Min
// returns the zero value for T.
func Min[T cmp.Ordered](x ...T) T {
	var m T
	if len(x) == 0 {
		return m
	}
	m = x[0]
	for _, t := range x[1:] {
		if t < m {
			m = t
		}
	}
	return m
}

// Sum sums any slice where the elements support the + operator.
// If len(in) == 0, the zero value for E is returned.
func Sum[S ~[]E, E Addable](in S) E {
	var accum E
	for _, x := range in {
		accum += x
	}
	return accum
}

// Prod computes the product of elements in any slice where the element type
// is numeric. If len(in) == 0, 1 is returned.
func Prod[S ~[]E, E Numeric](in S) E {
	var accum E = 1
	for _, x := range in {
		accum *= x
	}
	return accum
}

// MapMin finds the least value in the map m and returns the corresponding
// key and the value itself. If len(m) == 0, the zero values for K and V are
// returned. If there is a tie, the first key encountered is returned (which
// could be random).
func MapMin[M ~map[K]V, K comparable, V cmp.Ordered](m M) (K, V) {
	b := false
	var bestk K
	var minv V
	for k, v := range m {
		if !b || v < minv {
			b, bestk, minv = true, k, v
		}
	}
	return bestk, minv
}

// MapMax finds the greatest value in the map m and returns the corresponding
// key and the value itself. If len(m) == 0, the zero values for K and V are
// returned. If there is a tie, the first key encountered is returned (which
// could be random).
func MapMax[M ~map[K]V, K comparable, V cmp.Ordered](m M) (K, V) {
	b := false
	var bestk K
	var maxv V
	for k, v := range m {
		if !b || v > maxv {
			b, bestk, maxv = true, k, v
		}
	}
	return bestk, maxv
}

// MapRange reports the minimum and maximum values in the map m, and their
// corresponding keys. It does the work of MapMin and MapMax in one loop.
// If m is empty, the zero values for K and V are returned.
func MapRange[M ~map[K]V, K comparable, V cmp.Ordered](m M) (mink, maxk K, minv, maxv V) {
	minb, maxb := false, false
	for k, v := range m {
		if !minb || v < minv {
			minb, mink, minv = true, k, v
		}
		if !maxb || v > maxv {
			maxb, maxk, maxv = true, k, v
		}
	}
	return mink, maxk, minv, maxv
}

// MapKeyRange reports the minimum and maximum keys in the map m.
// If m is empty, the zero value for K is returned.
func MapKeyRange[M ~map[K]V, K cmp.Ordered, V any](m M) (min, max K) {
	minb, maxb := false, false
	for k := range m {
		if !minb || k < min {
			minb, min = true, k
		}
		if !maxb || k > max {
			maxb, max = true, k
		}
	}
	return min, max
}

// NextPermutation reorders s into the next permutation (in the lexicographic
// order), reporting if it was able to do so. Based on Knuth.
func NextPermutation[S ~[]E, E cmp.Ordered](s S) bool {
	if len(s) < 2 {
		return false
	}
	n1 := len(s) - 1
	i := n1 - 1
	for ; s[i] >= s[i+1]; i-- {
		if i == 0 {
			return false
		}
	}
	k := n1
	for s[i] >= s[k] {
		k--
	}
	s[i], s[k] = s[k], s[i]
	for j, k := i+1, n1; j < k; j, k = j+1, k-1 {
		s[j], s[k] = s[k], s[j]
	}
	return true
}
