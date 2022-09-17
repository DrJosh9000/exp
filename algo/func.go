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

// This file implements some functions that are unnecessary in readable Go.
// (Just write the loop!)

// Map calls f with each element of in, to build the output slice.
func Map[S, T any](in []S, f func(S) T) []T {
	out := make([]T, len(in))
	for i, x := range in {
		out[i] = f(x)
	}
	return out
}

// MapOrErr calls f with each element of in, to build the output slice.
// It stops and returns the first error returned by f.
func MapOrErr[S, T any](in []S, f func(S) (T, error)) ([]T, error) {
	out := make([]T, len(in))
	for i, x := range in {
		y, err := f(x)
		if err != nil {
			return nil, err
		}
		out[i] = y
	}
	return out, nil
}

// Foldl implements a functional "reduce" operation over slices.
// Loosely: Foldl(in, f) = f(f(f(...f(in[0], in[1]), in[2]),...), in[len(in)-1]).
// For example, if in is []int, Foldl(in, func(x, y int) int { return x + y })
// computes the sum. (The Sum function achieves the same thing in less code.)
// If len(in) == 0, the zero value for T is returned.
func Foldl[T any](in []T, f func(T, T) T) T {
	var accum T
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
func Foldr[T any](in []T, f func(T, T) T) T {
	var accum T
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
func Reverse[T any](s []T) {
	n1 := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[n1-i] = s[n1-i], s[i]
	}
}

// Freq counts the frequency of each item in the slice.
func Freq[T comparable](s []T) map[T]int {
	m := make(map[T]int)
	for _, x := range s {
		m[x]++
	}
	return m
}
