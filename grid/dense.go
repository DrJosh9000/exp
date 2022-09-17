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

package grid

import (
	"fmt"
	"image"
)

// Dense is a dense grid - a "2D array" generic type.
type Dense[T any] [][]T

// Make makes a Dense of width w and height h. 
//
// (If you want a sparse grid, use `make(Sparse[T])`.)
func Make[T any](h, w int) Dense[T] {
	g := make(Dense[T], h)
	for i := range g {
		g[i] = make([]T, w)
	}
	return g
}

// Map converts a Dense[S] into a Dense[T] by using a transformation function
// tf.
func Map[S, T any](g Dense[S], tf func(S) T) Dense[T] {
	ng := Make[T](g.Size())
	for j, row := range g {
		for i, x := range row {
			ng[j][i] = tf(x)
		}
	}
	return ng
}

// MapOrError converts a Dense[S] into a Dense[T] by using a transformation
// function tf. On the first error returned by tf, MapOrError returns an error.
func MapOrError[S, T any](g Dense[S], tf func(S) (T, error)) (Dense[T], error) {
	ng := Make[T](g.Size())
	for j, row := range g {
		for i, x := range row {
			y, err := tf(x)
			if err != nil {
				return nil, fmt.Errorf("transforming value at row %d, col %d: %w", j, i, err)
			}
			ng[j][i] = y
		}
	}
	return ng, nil
}

// ToSparse converts a dense grid into a sparse grid.
func (g Dense[T]) ToSparse() Sparse[T] {
	s := make(Sparse[T], g.Area())
	var p image.Point
	for p.Y = range g {
		for p.X = range g[p.Y] {
			s[p] = g[p.Y][p.X]
		}
	}
	return s
}

// Clone makes a copy of the grid.
func (g Dense[T]) Clone() Dense[T] {
	ng := Make[T](g.Size())
	for j := range g {
		for i := range g[j] {
			ng[j][i] = g[j][i]
		}
	}
	return ng
}

// Size returns the width and height of the Dense. If the height is zero the
// width will also be zero.
func (g Dense[T]) Size() (h, w int) {
	if len(g) == 0 {
		return 0, 0
	}
	return len(g), len(g[0])
}

// Area returns the area of the grid (width * height).
func (g Dense[T]) Area() int {
	h, w := g.Size()
	return h * w
}

// Fill fills the grid with the value v.
func (g Dense[T]) Fill(v T) {
	for _, row := range g {
		for i := range row {
			row[i] = v
		}
	}
}

// Transpose returns a new grid reflected about the diagonal.
func (g Dense[T]) Transpose() Dense[T] {
	h, w := g.Size()
	ng := Make[T](w, h) // note flipped dimensions
	for j, row := range g {
		for i := range row {
			ng[i][j] = row[i]
		}
	}
	return ng
}

// FlipH returns a new grid flipped horizontally (left becomes right).
func (g Dense[T]) FlipH() Dense[T] {
	ng := Make[T](g.Size())
	for j, row := range g {
		for i := range row {
			ng[j][i] = row[len(row)-i-1]
		}
	}
	return ng
}

// FlipV returns a new grid flipped vertically (top becomes bottom).
func (g Dense[T]) FlipV() Dense[T] {
	ng := Make[T](g.Size())
	for j, row := range g {
		for i := range row {
			ng[len(g)-j-1][i] = row[i]
		}
	}
	return ng
}

// RotateCW returns a new grid with entries rotated clockwise by 90 degrees.
func (g Dense[T]) RotateCW() Dense[T] {
	h, w := g.Size()
	ng := Make[T](w, h) // note flipped dimensions
	for j, row := range g {
		for i := range row {
			ng[i][h-j-1] = row[i]
		}
	}
	return ng
}

// RotateACW returns a new grid with entries rotated anticlockwise by 90
// degrees.
func (g Dense[T]) RotateACW() Dense[T] {
	h, w := g.Size()
	ng := Make[T](w, h) // note flipped dimensions
	for j, row := range g {
		for i := range row {
			ng[i][j] = row[w-i-1]
		}
	}
	return ng
}

// FromStringsFunc produces a grid from a slice of source strings s,
// containing data for one row per element, and a function for parsing each 
// string into a `[]T` row. Unequal-length rows are treated as an error.
func FromStringsFunc[T any](s []string, parse func(string) ([]T, error)) (Dense[T], error) {
	if len(s) == 0 {
		return nil, nil
	}
	g := make(Dense[T], len(s))
	for j, row := range s {
		t, err := parse(row)
		if err != nil {
			return nil, fmt.Errorf("parsing row %d: %w", j, err)
		}
		g[j] = t
		if got, want := len(g[j]), len(g[0]); got != want {
			return nil, fmt.Errorf("unequal row length (row %d length %d != row 0 length %d)", j, got, want)
		}
	}
	return g, nil
}
