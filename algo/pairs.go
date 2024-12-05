/*
   Copyright 2024 Josh Deprez

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

import "iter"

// AllPairs iterates over all pairs of items in the slice.
// e.g. AllPairs([]int{1, 2, 3}) iterates the pairs (1, 2), (1, 3), (2, 3).
func AllPairs[S ~[]E, E any](s S) iter.Seq2[E, E] {
	return func(yield func(a, b E) bool) {
		for i := range len(s) - 1 {
			for j := i + 1; j < len(s); j++ {
				if !yield(s[i], s[j]) {
					return
				}
			}
		}
	}
}
