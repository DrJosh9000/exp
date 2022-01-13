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

// Package algo implements a few generic algorithms.
package algo

// Dijkstra is an implementation of Dijkstra's algorithm for single-source 
// shortest paths on a directed, non-negatively weighted graph. The algorithm
// starts with a given node and assumes the zero value for D is the starting
// distance for that node. It then calls visit, passing each node and the length
// of the shortest path to that node. visit should either return a new
// collection of items and weights (the neighbours of the node it
// was given, and the weight of the edge connecting it) or an error. If visit
// returns a non-nil error, the algorithm halts and passes the error back to the
// caller. visit does not need to track already-visited nodes - it can return
// all known neighbours of a node.
// TODO: a means of returning a shortest path?
func Dijkstra[T comparable, D Orderable](start T, visit func(T, D) ([]WeightedItem[T, D], error)) error {
	var zero D
	done := make(map[T]bool)
	dist := map[T]D{start: zero}
	pq := new(PriQueue[T, D])
	pq.Push(start, zero)
	for pq.Len() > 0 {
		node, _ := pq.Pop()
		done[node] = true
		next, err := visit(node, dist[node])
		if err != nil {
			return err
		}
		for _, wi := range next {
			if done[wi.Item] {
				continue
			}
			newdist := dist[node] + wi.Weight
			if olddist, seen := dist[wi.Item]; seen && olddist >= newdist {
				continue
			}
			dist[wi.Item] = newdist
			pq.Push(wi.Item, newdist)
		}
	}
	return nil
}