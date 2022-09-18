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
	"math/rand"
	"testing"
)

// Fewer lookups in a larger map, or more lookups in a smaller map?
// Seems fewer lookups wins:
//    BenchmarkDisjoint-36              541929              2124 ns/op
//    BenchmarkFlipDisjoint-36            5152            223984 ns/op

func flipDisjoint[T comparable](s, t Set[T]) bool {
	// Opposite of Set[T].Disjoint.
	if len(s) < len(t) {
		s, t = t, s
	}
	for x := range s {
		if t.Contains(x) {
			return false
		}
	}
	return true
}

func randomSet(n int) Set[int] {
	s := make(Set[int], n)
	for i := 0; i < n; i++ {
		var r int
		for {
			r = rand.Int()
			if !s.Contains(r) {
				break
			}
		}
		s.Insert(r)	
	}
	return s
}

func BenchmarkDisjoint(b *testing.B) {
	s, t := randomSet(100), randomSet(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s, t = t, s
		_ = s.Disjoint(t)
	}
}

func BenchmarkFlipDisjoint(b *testing.B) {
	s, t := randomSet(100), randomSet(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s, t = t, s
		_ = flipDisjoint(s, t)
	}
}