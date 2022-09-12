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

// Package stream provides a variety of context-aware generic functions that
// operate on channels.
package stream

import (
	"context"
	"fmt"	
)

// NopSource closes the channel. The implementation is trivial, but is 
// provided out of a sense of dedication to completeness.
func NopSource[T any](_ context.Context, out chan<- T) {
	close(out)
}

// NopSink reads and discards all the values from the channel.
func NopSink[T any](ctx context.Context, in <-chan T) error {
	for {
		select {
		case _, open := <-in:
			if !open {
				return nil
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// SliceSource writes the values in the slice to the channel, one at a time,
// and closes the channel when done.
func SliceSource[T any](ctx context.Context, in []T, out chan<- T) error {
	for _, x := range in {
		select {
		case out <- x:
			// success
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	close(out)
	return nil
}

// SliceSink reads the values from the channel into a slice, until the channel
// is closed, then returns the slice.
func SliceSink[T any](ctx context.Context, in <-chan T) ([]T, error) {
	var out []T
	for {
		select {
		case x, open := <-in:
			if !open {
				return out, nil
			}
			out = append(out, x)
		case <-ctx.Done():
			return out, ctx.Err()
		}
	}
}

// TextFileSource reads lines from the file one at a time, and sends them into
// the channel. When reading is complete the channel is closed.
func TextFileSource(ctx context.Context, path string, out chan<- string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		select {
		case out <- sc.Text():
			// success
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	if err := sc.Err(); err != nil {
		return fmt.Errorf("scanner: %w", err)
	}
	close(out)
	return nil
}

// Transform receives values from in, applies the transform tf, and sends the
// results on out. It does not receive a new value before sending the previous
// result.
func Transform[S, T any](ctx context.Context, in <-chan S, out chan<- T, tf func(context.Context, S) (T, error)) error {
	for {
		select {
		case s, open := <-in:
			if !open {
				return nil
			}
			t, err := tf(ctx, s)
			if err != nil {
				return err
			}
			select {
			case out <- t:
				// success
			case <-ctx.Done():
				return ctx.Err()
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
