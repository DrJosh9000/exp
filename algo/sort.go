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
	"sort"
)

// SortSlice is like sort.Slice, but the less func receives the items
// themselves.
func SortSlice[S ~[]E, E any](s S, less func(a, b E) bool) {
	sort.Sort(sliceByLess[S, E]{
		slice: s,
		less:  less,
	})
}

// SortSlice is like sort.SliceStable, but the less func receives the items
// themselves.
func SortSliceStable[S ~[]E, E any](s S, less func(a, b E) bool) {
	sort.Stable(sliceByLess[S, E]{
		slice: s,
		less:  less,
	})
}

type sliceByLess[S ~[]E, E any] struct {
	slice S
	less  func(a, b E) bool
}

func (b sliceByLess[S, E]) Len() int      { return len(b.slice) }
func (b sliceByLess[S, E]) Swap(i, j int) { b.slice[i], b.slice[j] = b.slice[j], b.slice[i] }
func (b sliceByLess[S, E]) Less(i, j int) bool {
	return b.less(b.slice[i], b.slice[j])
}

// SortAsc sorts the slice in ascending order.
func SortAsc[S ~[]E, E cmp.Ordered](s S) {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
}

// SortDesc sort the slice in descending order.
func SortDesc[S ~[]E, E cmp.Ordered](s S) {
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
}

// SortByMapAsc stably sorts the slice using the map to provide values to
// compare.
func SortByMapAsc[S ~[]K, M ~map[K]V, K comparable, V cmp.Ordered](s S, m M) {
	sort.SliceStable(s, func(i, j int) bool { return m[s[i]] < m[s[j]] })
}

// SortByMapDesc stably sorts the slice using the map to provide values to
// compare.
func SortByMapDesc[S ~[]K, M ~map[K]V, K comparable, V cmp.Ordered](s S, m M) {
	sort.SliceStable(s, func(i, j int) bool { return m[s[i]] > m[s[j]] })
}

// SortByFuncAsc stably sorts the slice using the function to provide values to
// compare.
func SortByFuncAsc[S ~[]E, E any, V cmp.Ordered](s S, f func(E) V) {
	sort.SliceStable(s, func(i, j int) bool { return f(s[i]) < f(s[j]) })
}

// SortByFuncDesc stably sorts the slice using the function to provide values to
// compare.
func SortByFuncDesc[S ~[]E, E any, V cmp.Ordered](s S, f func(E) V) {
	sort.SliceStable(s, func(i, j int) bool { return f(s[i]) > f(s[j]) })
}
