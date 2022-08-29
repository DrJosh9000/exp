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

// MapMin finds the smallest value in the map m and returns the corresponding key and
// the value itself. If len(m) == 0, the zero values for K and V are returned. If
// there is a tie, the first key encountered is returned (which could be random).
func MapMin[K comparable, V Orderable](m map[K]V) (K, V) {
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

// MapMax finds the largest value in the map m and returns the corresponding key and
// the value itself. If len(m) == 0, the zero values for K and V are returned. If
// there is a tie, the first key encountered is returned (which could be random).
func MapMax[K comparable, V Orderable](m map[K]V) (K, V) {
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