package exp

import (
	"context"
	"errors"
)

// Break is used to exit the RangeCh loop early without error.
var Break = errors.New("break")

// RangeCh ranges over a channel in a context-aware way.
func RangeCh[T any](ctx context.Context, ch <-chan T, f func(T) error) error {
	for {
		select {
		case t, open := <-ch:
			if !open {
				return nil
			}	
			if err := f(t); err != nil {
				if errors.Is(err, Break) {
					return nil
				}
				return err
			}
		
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func ExampleRangeCh() {
	ctx := context.Background()
	ch := make(chan struct{})
	
	_ = RangeCh(ctx, ch, func(struct{}) error {
		return nil
	})
}