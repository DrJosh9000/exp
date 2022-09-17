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
