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

import (
	"cmp"
	"iter"
)

// AStar implements A* search, a variant of Dijkstra's algorithm which takes
// an additional heuristic h into account. h(x) should return an estimate of
// d(x, end). If h(x) <= d(x, end) (that is, h underestimates) then AStar
// will find the shortest path. It returns a map of each node to the previous
// node(s) in the shortest path(s) to that node. This predecessor map is only
// complete for visited nodes.
//
// The algorithm starts with a given start node and assumes the zero value for D
// is the starting distance for that node. It then repeatedly calls visit,
// passing each node and the length of the shortest path to that node. visit
// should either return a new iterator over nodes and weights (the neighbours of
// the node it was given, and the weight of the edge connecting it) or an error.
// If visit returns a non-nil error, the algorithm halts and passes both the error
// and the partial map of predecessors, back to the caller. The algorithm takes
// care of tracking nodes that have already been visited - since visit does not
// need to track already-visited nodes, it can safely return all known neighbours
// of a node.
func AStar[T comparable, D cmp.Ordered](start T, h func(T) D, visit func(T, D) (iter.Seq2[T, D], error)) (map[T][]T, error) {
	prev := make(map[T][]T)
	done := make(map[T]bool)
	var zero D
	dist := map[T]D{start: zero}
	pq := new(PriQueue[T, D])
	pq.Push(start, zero)
	for pq.Len() > 0 {
		node, _ := pq.Pop()
		if done[node] {
			continue
		}
		done[node] = true
		it, err := visit(node, dist[node])
		if err != nil {
			return prev, err
		}
		if it == nil {
			continue
		}
		for newnode, weight := range it {
			newdist := dist[node] + weight
			if olddist, seen := dist[newnode]; seen {
				switch {
				case olddist < newdist:
					continue
				case olddist == newdist:
					prev[newnode] = append(prev[newnode], node)
					continue
				}
			}
			// !seen || (seen && olddist > newdist)
			prev[newnode] = []T{node}
			dist[newnode] = newdist
			pq.Push(newnode, newdist+h(newnode))
		}
	}
	return prev, nil
}

// Dijkstra is an implementation of Dijkstra's algorithm for single-source
// shortest paths on a directed, non-negatively weighted graph. It is equivalent
// to AStar with a heuristic that always returns zero. It returns a map of each
// node to the previous node in the shortest path to that node. This predecessor
// map is only complete for visited nodes.
//
// The algorithm starts with a given start node and assumes the zero value for D
// is the starting distance for that node. It then repeatedly calls visit,
// passing each node and the length of the shortest path to that node. visit
// should either return a new iterator over items and weights (the neighbours of
// the node it was given, and the weight of the edge connecting it) or an error.
// If visit returns a non-nil error, the algorithm halts and passes both the error
// and the partial map of predecessors, back to the caller. The algorithm takes
// care of tracking nodes that have already been visited - since visit does not
// need to track already-visited nodes, it can safely return all known neighbours
// of a node.
func Dijkstra[T comparable, D cmp.Ordered](start T, visit func(T, D) (iter.Seq2[T, D], error)) (map[T][]T, error) {
	var zero D
	return AStar(start, func(T) D { return zero }, visit)
}
