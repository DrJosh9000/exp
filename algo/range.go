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
	"fmt"

	"golang.org/x/exp/constraints"
)

// Range represents an **inclusive** range of values.
type Range[T cmp.Ordered] struct {
	Min, Max T
}

// NewRange returns a range [from, to].
func NewRange[T cmp.Ordered](from, to T) Range[T] {
	return Range[T]{Min: from, Max: to}
}

func (r Range[T]) String() string {
	return fmt.Sprintf("[%v, %v]", r.Min, r.Max)
}

// Contains reports if r contains x.
func (r Range[T]) Contains(x T) bool {
	return r.Min <= x && x <= r.Max
}

// ContainsRange reports if r contains the entirety of s.
func (r Range[T]) ContainsRange(s Range[T]) bool {
	return r.Contains(s.Min) && r.Contains(s.Max)
}

// Clamp returns x if x is within the range, r.Min if x is less than the range,
// or r.Max if x is greather than the range.
func (r Range[T]) Clamp(x T) T {
	switch {
	case x < r.Min:
		return r.Min
	case x > r.Max:
		return r.Max
	}
	return x
}

// IsEmpty reports if the range is empty.
func (r Range[T]) IsEmpty() bool {
	return r.Max < r.Min
}

// Intersection returns the intersection of the two ranges. (It could be empty.)
func (r Range[T]) Intersection(s Range[T]) Range[T] {
	return Range[T]{
		Min: max(r.Min, s.Min),
		Max: min(r.Max, s.Max),
	}
}

// Union returns a range covering both ranges. (It could include subranges that
// don't overlap either r or s).
func (r Range[T]) Union(s Range[T]) Range[T] {
	return Range[T]{
		Min: min(r.Min, s.Min),
		Max: max(r.Max, s.Max),
	}
}

// RangeSubtract returns the set difference of two ranges (r - s).
// If r and s do not overlap, it returns only r.
// If s completely overlaps r, it returns an empty slice.
// If r overlaps s, and (r - s) is not empty, it returns one or two ranges
// depending on how the ranges overlap.
func RangeSubtract[T constraints.Integer](r, s Range[T]) []Range[T] {
	var rs []Range[T]
	if r0 := (Range[T]{Min: r.Min, Max: s.Min - 1}); !r0.IsEmpty() {
		rs = append(rs, r0)
	}
	if r0 := (Range[T]{Min: s.Max + 1, Max: r.Max}); !r0.IsEmpty() {
		rs = append(rs, r0)
	}
	return rs
}

// RangeAdd adds a number to a range.
func RangeAdd[T Real](r Range[T], d T) Range[T] {
	r.Min += d
	r.Max += d
	return r
}

// RangeMul multiplies a range by a number.
func RangeMul[T Real](r Range[T], m T) Range[T] {
	r.Min *= m
	r.Max *= m
	if m < 0 {
		r.Min, r.Max = r.Max, r.Min
	}
	return r
}
