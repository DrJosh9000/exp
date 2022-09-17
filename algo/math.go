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

import "golang.org/x/exp/constraints"

// Abs returns the absolute value of x (with no regard for negative overflow).
//
// The math/cmplx package provides a version of Abs for complex types.
func Abs[T Real](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Max returns the greatest argument (using `>`). If no arguments are provided,
// Max returns the zero value for T.
func Max[T constraints.Ordered](x ...T) T {
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
func Min[T constraints.Ordered](x ...T) T {
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
func MapMin[M ~map[K]V, K comparable, V constraints.Ordered](m M) (K, V) {
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
func MapMax[M ~map[K]V, K comparable, V constraints.Ordered](m M) (K, V) {
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
