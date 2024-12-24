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

import "iter"

// FloodFill is an algorithm for finding single-source shortest paths in an
// unweighted directed graph. It follows the same conventions as the Dijkstra
// function. It assumes the weight of each edge is always 1, tallying distances
// as ints. This flood-fill is generic and makes minimal assumptions about each
// node. A more specific implementation than this one is more appropriate in
// some cases, e.g. flood-filling a 2D grid.
func FloodFill[T comparable](start T, visit func(T, int) (iter.Seq[T], error)) (map[T][]T, error) {
	prev := make(map[T][]T)
	dist := map[T]int{start: 0}
	q := []T{start}
	var node T
	for len(q) > 0 {
		node, q = q[0], q[1:]
		next, err := visit(node, dist[node])
		if err != nil {
			return prev, err
		}
		if next == nil {
			continue
		}
		for newnode := range next {
			newdist := dist[node] + 1
			olddist, seen := dist[newnode]
			if seen {
				switch {
				case olddist < newdist:
					continue
				case olddist == newdist:
					prev[newnode] = append(prev[newnode], node)
					continue
				}
			}
			dist[newnode] = newdist
			prev[newnode] = append(prev[newnode], node)
			q = append(q, newnode)
		}
	}
	return prev, nil
}
