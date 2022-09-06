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

// Set is a generic set type based on map. 
// There's a million of these now; what harm is another?
//
// Set aims to work in a slice-like fashion with nil-valued sets, i.e.:
//   var s Set[int]
//   s = s.Insert(420, 69)   // s now contains 420 and 69
// However, Insert (and Add) do not make new sets if the set is non-nil, e.g.:
//   s := make(Set[int])
//   s.Insert(420, 69)   // as above
type Set[T comparable] map[T]struct{}

// ToSlice returns a new slice with all the elements of the set in random order.
func (s Set[T]) ToSlice() []T {
	sl := make([]T, 0, len(s))
	for x := range s {
		sl = append(sl, x)
	}
	return sl
}

// Insert inserts x into the set. If s == nil, Insert returns a new set
// containing x. This allows patterns like `m[k] = m[k].Insert(x)` (for a map m
// with set values).
func (s Set[T]) Insert(x ...T) Set[T] {
	if s == nil {
		s = make(Set[T])
	}
	for _, x := range x {
		s[x] = struct{}{}
	}
	return s
}

// make(Set[T]), delete(s, x), and len(s) are so simple that I'm not making
// methods.

// Contains reports whether s contains x.
func (s Set[T]) Contains(x T) bool {
	_, c := s[x]
	return c
}

// Add adds the elements from t into s, and returns s. If s == nil, Add 
// returns a new set which is a copy of t.
func (s Set[T]) Add(t Set[T]) Set[T] {
	if s == nil {
		s = make(Set[T], len(t))
	}
	for x := range t {
		s = s.Insert(x)
	}
	return s
}

// Subtract removes all elements in t from s, and returns s.
func (s Set[T]) Subtract(t Set[T]) Set[T] {
	if s == nil {
		return s
	}
	for x := range t {
		delete(s, x)
	}
	return s
}

// Copy returns a copy of the set. It is equivalent to Set[T].Add(nil, s).
func (s Set[T]) Copy() Set[T] {
	return Set[T].Add(nil, s)
}

// Union returns a new set containing elements from both sets.
func (s Set[T]) Union(t Set[T]) Set[T] {
	u := make(Set[T], len(s) + len(t))
	u.Add(s)
	u.Add(t)
	return u
}

// Intersection returns a new set with elements common to both sets. 
func (s Set[T]) Intersection(t Set[T]) Set[T] {
	u := make(Set[T])
	for x := range s {
		if t.Contains(x) {
			u.Insert(x)
		}
	}
	return u
}

// Difference returns a new set with elements from s that are not in t.
func (s Set[T]) Difference(t Set[T]) Set[T] {
	return s.Copy().Subtract(t)
}

// SymmetricDifference returns a new set with elements that are either in s, or
// in t, but not in both.
func (s Set[T]) SymmetricDifference(t Set[T]) Set[T] {
	u := make(Set[T])
	for x := range s {
		if !t.Contains(x) {
			u.Insert(x)
		}
	}
	for x := range t {
		if !s.Contains(x) {
			u.Insert(x)
		}
	}
	return u
}

// SubsetOf reports whether s is a subset or equal to t.
func (s Set[T]) SubsetOf(t Set[T]) bool {
	for x := range s {
		if !t.Contains(x) {
			return false
		}
	}
	return true
}

// Equal reports whether two sets are equal.
func (s Set[T]) Equal(t Set[T]) bool {
	return s.SubsetOf(t) && t.SubsetOf(s)
}