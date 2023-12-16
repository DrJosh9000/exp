package para

import (
	"sync/atomic"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue(1)

	var sum atomic.Int64
	q.Process(func(x int) {
		if x > 100 {
			return
		}
		sum.Add(int64(x))
		q.Push(2*x, 2*x+1)
	})

	if got, want := int(sum.Load()), 5050; got != want {
		t.Errorf("sum = %d, want %d", got, want)
	}
}
