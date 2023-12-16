package para

import (
	"runtime"
	"sync"
)

// Queue implements a goroutine-safe queue of items based on a channel.
type Queue[E any] struct {
	µ  sync.Mutex
	c  int
	ch chan E
}

// NewQueue returns a new queue with an initial list of items.
func NewQueue[E any](items ...E) *Queue[E] {
	q := &Queue[E]{
		ch: make(chan E, max(65536, len(items))),
	}
	q.c = len(items)
	for _, e := range items {
		q.ch <- e
	}
	return q
}

// Process calls f on each item in the queue in parallel (with GOMAXPROCS
// goroutines). f can enqueue more items by calling q.Push. It blocks until the
// queue is empty and all items have been processed.
func (q *Queue[E]) Process(f func(E)) {
	N := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for e := range q.ch {
				f(e)
				q.done()
			}
		}()
	}

	wg.Wait()
}

// Push appends to the end of the queue.
func (q *Queue[E]) Push(items ...E) {
	for _, e := range items {
		q.µ.Lock()
		q.c++
		q.ch <- e
		q.µ.Unlock()
	}
}

// done decrements the queue's internal counter. If there are no outstanding
// items and no queued items, the channel is closed.
func (q *Queue[E]) done() {
	q.µ.Lock()
	q.c--
	if q.c == 0 {
		close(q.ch)
	}
	q.µ.Unlock()
}
