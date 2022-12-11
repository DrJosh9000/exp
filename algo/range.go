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

// Range represents an inclusive range of values.
type Range[T constraints.Ordered] struct {
	Min, Max T
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
		Min: Max(r.Min, s.Min),
		Max: Min(r.Max, s.Max),
	}
}

// Union returns a range covering both ranges. (It could include subranges that
// don't overlap either r or s).
func (r Range[T]) Union(s Range[T]) Range[T] {
	return Range[T]{
		Min: Min(r.Min, s.Min),
		Max: Max(r.Max, s.Max),
	}
}
