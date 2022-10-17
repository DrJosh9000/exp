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
	"sort"

	"golang.org/x/exp/constraints"
)

// SortAsc sorts the slice in ascending order.
func SortAsc[S ~[]E, E constraints.Ordered](s S) {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
}

// SortDesc sort the slice in descending order.
func SortDesc[S ~[]E, E constraints.Ordered](s S) {
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
}

// SortByMapAsc stably sorts the slice using the map to provide values to
// compare.
func SortByMapAsc[S ~[]K, M ~map[K]V, K comparable, V constraints.Ordered](s S, m M) {
	sort.SliceStable(s, func(i, j int) bool { return m[s[i]] < m[s[j]] })
}

// SortByMapDesc stably sorts the slice using the map to provide values to
// compare.
func SortByMapDesc[S ~[]K, M ~map[K]V, K comparable, V constraints.Ordered](s S, m M) {
	sort.SliceStable(s, func(i, j int) bool { return m[s[i]] > m[s[j]] })
}

// SortByFuncAsc stably sorts the slice using the function to provide values to
// compare.
func SortByFuncAsc[S ~[]E, E any, V constraints.Ordered](s S, f func(E) V) {
	sort.SliceStable(s, func(i, j int) bool { return f(s[i]) < f(s[j]) })
}

// SortByFuncDesc stably sorts the slice using the function to provide values to
// compare.
func SortByFuncDesc[S ~[]E, E any, V constraints.Ordered](s S, f func(E) V) {
	sort.SliceStable(s, func(i, j int) bool { return f(s[i]) > f(s[j]) })
}
