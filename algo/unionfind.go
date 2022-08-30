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

import "math/rand"

// DisjointSets implements union-find algorithms for disjoint sets.
type DisjointSets[K comparable] map[K]K

// Find returns the representative element of the set containing x.
// It freely modifies d.
// If x is not contained in d, Find inserts x as a new disjoint singleton set
// within d and returns x.
func (d DisjointSets[K]) Find(x K) K {
   if _, found := d[x]; !found {
      d[x] = x
      return x
   }
   for x != d[x] {
      d[x] = d[d[x]] // path compression
      x = d[x]
   }
   return x
}

// Union merges the set containing x with the set containing y.
// If both of the elements are not contained in d, a new set is created.
func (d DisjointSets[K]) Union(x, y K) {
   p, q := d.Find(x), d.Find(y)
   if p == q {
      return
   }
   if rand.Intn(2) == 0 {
      d[p] = q
   } else {
      d[q] = p
   }
}