/*
   Copyright 2021 Josh Deprez

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

import "container/heap"

// PriQueue implements a priority queue of items of type T each having a
// priority of type D. It uses container/heap under the hood.
type PriQueue[T any, D Orderable] minHeap[T, D]

// Push adds an item to the queue with a priority.
func (pq *PriQueue[T, D]) Push(item T, priority D) {
	heap.Push((*minHeap[T, D])(pq), WeightedItem[T, D]{
		Item: item,
		Weight: priority,
	})
}

// Pop removes the item with the least priority value, and returns both it and
// its priority.
func (pq *PriQueue[T, D]) Pop() (T, D) {
	hi := heap.Pop((*minHeap[T, D])(pq)).(WeightedItem[T, D])
	return hi.Item, hi.Weight
}

// Len returns the size of the queue.
func (pq *PriQueue[T, D]) Len() int { return len(*pq) }

// minHeap provides the underlying implementation of heap.Interface.
type minHeap[T any, D Orderable] []WeightedItem[T, D]

func (h minHeap[T, D]) Len() int            { return len(h) }
func (h minHeap[T, D]) Less(i, j int) bool  { return h[i].Weight < h[j].Weight }
func (h minHeap[T, D]) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap[T, D]) Push(x interface{}) { *h = append(*h, x.(WeightedItem[T, D])) }
func (h *minHeap[T, D]) Pop() interface{} {
	n1 := len(*h) - 1
	i := (*h)[n1]
	*h = (*h)[0:n1]
	return i
}

// WeightedItem is an item together with a weight value.
type WeightedItem[T any, D Orderable] struct{
	Item T
	Weight D
}
