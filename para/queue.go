package para

import "sync"

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

// Push appends to the end of the queue. Be sure to Push before calling Done.
func (q *Queue[E]) Push(e E) {
	q.µ.Lock()
	q.ch <- e
	q.c++
	q.µ.Unlock()
}

// Pop is equivalent to a channel read.
func (q *Queue[E]) Pop() (e E, ok bool) {
	e, ok = <-q.ch
	return e, ok
}

// Done decrements the queue's internal counter. If there are no outstanding
// items and no queued items, the channel is closed.
func (q *Queue[E]) Done() {
	q.µ.Lock()
	q.c--
	if q.c == 0 {
		close(q.ch)
	}
	q.µ.Unlock()
}
