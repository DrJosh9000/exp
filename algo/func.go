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
	"fmt"
	"iter"

	"golang.org/x/exp/constraints"
)

// This file implements some functions that are unnecessary in readable Go.
// (Just write the loop!)

// Map calls f with each element of in, to build the output slice.
func Map[S ~[]X, X, Y any](in S, f func(X) Y) []Y {
	out := make([]Y, len(in))
	for i, x := range in {
		out[i] = f(x)
	}
	return out
}

// MapOrErr calls f with each element of in, to build the output slice.
// It stops and returns the first error returned by f.
func MapOrErr[S ~[]X, X, Y any](in S, f func(X) (Y, error)) ([]Y, error) {
	out := make([]Y, len(in))
	for i, x := range in {
		y, err := f(x)
		if err != nil {
			return nil, err
		}
		out[i] = y
	}
	return out, nil
}

// MustMap calls f with each element of in, to build the output slice.
// It panics with the first error returned by f.
func MustMap[S ~[]X, X, Y any](in S, f func(X) (Y, error)) []Y {
	out := make([]Y, len(in))
	for i, x := range in {
		y, err := f(x)
		if err != nil {
			panic(fmt.Sprintf("at index %d: %v", i, err))
		}
		out[i] = y
	}
	return out
}

// MapMap is like Map, but for maps. This seems thoroughly pointless.
func MapMap[M ~map[K1]V1, K1, K2 comparable, V1, V2 any](m M, f func(K1, V1) (K2, V2)) map[K2]V2 {
	n := make(map[K2]V2, len(m))
	for k1, v1 := range m {
		k2, v2 := f(k1, v1)
		n[k2] = v2
	}
	return n
}

// MustMapMap is like MustMap, but for maps. This seems extra pointless.
func MustMapMap[M ~map[K1]V1, K1, K2 comparable, V1, V2 any](m M, f func(K1, V1) (K2, V2, error)) map[K2]V2 {
	n := make(map[K2]V2, len(m))
	for k1, v1 := range m {
		k2, v2, err := f(k1, v1)
		if err != nil {
			panic(fmt.Sprintf("at key %v: %v", k1, err))
		}
		n[k2] = v2
	}
	return n
}

// Foldl implements a functional "reduce" operation over slices.
// Loosely: Foldl(in, f) = f(f(f(...f(in[0], in[1]), in[2]),...), in[len(in)-1]).
// For example, if in is []int, Foldl(in, func(x, y int) int { return x + y })
// computes the sum. (The Sum function achieves the same thing in less code.)
// If len(in) == 0, the zero value for E is returned.
func Foldl[S ~[]E, E any](in S, f func(E, E) E) E {
	var accum E
	if len(in) == 0 {
		return accum
	}
	accum = in[0]
	for _, x := range in[1:] {
		accum = f(accum, x)
	}
	return accum
}

// Foldr is the same as Foldl, but considers elements in the reverse.
func Foldr[S ~[]E, E any](in S, f func(E, E) E) E {
	var accum E
	if len(in) == 0 {
		return accum
	}
	accum = in[len(in)-1]
	n2 := len(in) - 2
	for i := range in[1:] {
		accum = f(accum, in[n2-i])
	}
	return accum
}

// Reverse reverses a slice.
func Reverse[S ~[]E, E any](s S) {
	n1 := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[n1-i] = s[n1-i], s[i]
	}
}

// Count counts the number of occurrences of e in the slice s, in O(len(s))
// time. This is faster than Freq when counting only one or a few values.
func Count[S ~[]E, E comparable](s S, e E) int {
	count := 0
	for _, x := range s {
		if x == e {
			count++
		}
	}
	return count
}

// MapCount counts the number of occurrences of v in the map m, in O(len(m))
// time. This is faster than MapFreq when counting only one or a few values.
func MapCount[M ~map[K]V, K, V comparable](m M, v V) int {
	count := 0
	for _, x := range m {
		if x == v {
			count++
		}
	}
	return count
}

// Freq counts the frequency of each item in a slice.
func Freq[S ~[]E, E comparable](s S) map[E]int {
	h := make(map[E]int)
	for _, x := range s {
		h[x]++
	}
	return h
}

// FreqIter counts the frequency of each item in an iterator.
func FreqIter[E comparable](i iter.Seq[E]) map[E]int {
	h := make(map[E]int)
	for x := range i {
		h[x]++
	}
	return h
}

// FreqIter2 counts the frequency of each second item in an iterator.
func FreqIter2[X any, E comparable](i iter.Seq2[X, E]) map[E]int {
	h := make(map[E]int)
	for _, x := range i {
		h[x]++
	}
	return h
}

// MapFreq counts the frequency of each value in a map.
func MapFreq[M ~map[K]V, K, V comparable](m M) map[V]int {
	h := make(map[V]int)
	for _, x := range m {
		h[x]++
	}
	return h
}

// MapFromSlice creates a map[int]E from a slice of E. There's usually no
// reason for this.
func MapFromSlice[S ~[]E, E any](s S) map[int]E {
	m := make(map[int]E, len(s))
	for i, e := range s {
		m[i] = e
	}
	return m
}

// SliceFromMap creates a slice of V from a map[K]V, plus an offset value
// such that s[k] == m[k+offset] for all k in m.
// The slice will be exactly large enough to contain all keys (offsetted).
// The result will be extremely memory wasteful if the keys in m are few and far
// between.
// Entries in the slice with no corresponding entry in the map will have the
// zero value.
func SliceFromMap[M ~map[K]V, K constraints.Integer, V any](m M) ([]V, K) {
	min, max := MapKeyRange(m)
	s := make([]V, max-min+1)
	for k, v := range m {
		s[k-min] = v
	}
	return s, min
}
