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

import "sync"

// Topic is a pub-sub topic.
type Topic[T any] struct {
	mu     sync.Mutex
	subs   []chan T
	closed bool
}

// Pub publishes the value v onto the topic.
func (q *Topic[T]) Pub(v T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		panic("topic is closed")
	}
	for _, c := range q.subs {
		c <- v
	}
}

// Sub subscribes to subsequent messages on the topic.
func (q *Topic[T]) Sub() <-chan T {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.closed {
		panic("topic is closed")
	}
	ch := make(chan T, 1)
	q.subs = append(q.subs, ch)
	return ch
}

// Close closes the topic.
func (q *Topic[T]) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()
	for _, c := range q.subs {
		close(c)
	}
	q.subs = nil
	q.closed = true
}
