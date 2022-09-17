package algo

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// SortAsc sorts the slice in ascending order.
func SortAsc[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
}

// SortDesc sort the slice in descending order.
func SortDesc[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
}

// SortByMapAsc stably sorts the slice using the map to provide values to
// compare.
func SortByMapAsc[K comparable, V constraints.Ordered](s []K, m map[K]V) {
	sort.SliceStable(s, func(i, j int) bool { return m[s[i]] < m[s[j]] })
}

// SortByMapDesc stably sorts the slice using the map to provide values to
// compare.
func SortByMapDesc[K comparable, V constraints.Ordered](s []K, m map[K]V) {
	sort.SliceStable(s, func(i, j int) bool { return m[s[i]] > m[s[j]] })
}

// SortByFuncAsc stably sorts the slice using the function to provide values to
// compare.
func SortByFuncAsc[T any, V constraints.Ordered](s []T, f func(T) V) {
	sort.SliceStable(s, func(i, j int) bool { return f(s[i]) < f(s[j]) })
}

// SortByFuncDesc stably sorts the slice using the function to provide values to
// compare.
func SortByFuncDesc[T any, V constraints.Ordered](s []T, f func(T) V) {
	sort.SliceStable(s, func(i, j int) bool { return f(s[i]) > f(s[j]) })
}
