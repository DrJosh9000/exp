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
	"image/color"
	"strings"
)

// Dense is a dense grid - a "2D array" generic type.
type Dense[T any] [][]T

// Make makes a dense grid of width w and height h.
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
	if len(g) == 0 {
		return nil
	}
	ng := Make[T](g.Size())
	for j, row := range g {
		for i, x := range row {
			ng[j][i] = tf(x)
		}
	}
	return ng
}

// MapToSparse converts a Dense[S] into a Sparse[T] by using a transformation
// function tf, which should also report if the element should be kept.
func MapToSparse[S, T any](g Dense[S], tf func(S) (T, bool)) Sparse[T] {
	if len(g) == 0 {
		return nil
	}
	ng := make(Sparse[T])
	for j, row := range g {
		for i, x := range row {
			if v, ok := tf(x); ok {
				ng[image.Point{i, j}] = v
			}
		}
	}
	return ng
}

// MapOrError converts a Dense[S] into a Dense[T] by using a transformation
// function tf. On the first error returned by tf, MapOrError returns an error.
func MapOrError[S, T any](g Dense[S], tf func(S) (T, error)) (Dense[T], error) {
	if len(g) == 0 {
		return nil, nil
	}
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

// ToRGBA converts a dense grid into an RGBA image using a colouring function.
func (g Dense[T]) ToRGBA(cf func(T) color.Color) *image.RGBA {
	img := image.NewRGBA(g.Bounds())
	for y, row := range g {
		for x, t := range row {
			img.Set(x, y, cf(t))
		}
	}
	return img
}

func (g Dense[T]) String() string {
	if g == nil {
		return "nil"
	}
	if len(g) == 0 {
		return "[]"
	}

	leftAlign := false
	padding := true

	var t T
	switch any(t).(type) {
	case bool, byte, rune:
		padding = false
	case string:
		leftAlign = true
	}

	// Format the grid into strings.
	h := Map(g, func(x T) string {
		switch x := any(x).(type) {
		case string:
			// Render as themselves
			return x
		case byte:
			// Render as themselves with no padding
			return string(x)
		case rune:
			// Render as themselves with no padding
			return string(x)
		case bool:
			// Render bools as space(false) or █(true) with no padding
			if x {
				return "█"
			}
			return " "
		}
		return fmt.Sprint(x)
	})

	// Find column widths large enough for all items.
	cw := make([]int, len(g[0]))
	for _, row := range h {
		for i, el := range row {
			if len(el) > cw[i] {
				cw[i] = len(el)
			}
		}
	}

	// Build the output string.
	sb := new(strings.Builder)
	sb.WriteString("[\n")
	for _, row := range h {
		sb.WriteString(" [ ")
		for i, el := range row {
			if !padding {
				sb.WriteString(el)
				continue
			}
			pad := strings.Repeat(" ", cw[i]-len(el))
			if i != 0 {
				sb.WriteRune(' ')
			}
			if !leftAlign {
				sb.WriteString(pad)
			}
			sb.WriteString(el)
			if leftAlign {
				sb.WriteString(pad)
			}
		}
		sb.WriteString(" ]\n")
	}
	sb.WriteString("]")
	return sb.String()
}

// Clone makes a copy of the grid.
func (g Dense[T]) Clone() Dense[T] {
	ng := Make[T](g.Size())
	for j := range g {
		copy(ng[j], g[j])
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

// Bounds returns a rectangle the size of the grid.
func (g Dense[T]) Bounds() image.Rectangle {
	h, w := g.Size()
	return image.Rect(0, 0, w, h)
}

// Fill fills the grid with the value v.
func (g Dense[T]) Fill(v T) {
	for _, row := range g {
		for i := range row {
			row[i] = v
		}
	}
}

// FillRect fills a sub-rectangle of the grid with the value v.
func (g Dense[T]) FillRect(r image.Rectangle, v T) {
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			g[y][x] = v
		}
	}
}

// Map applies a transformation function to each element in the grid in-place.
func (g Dense[T]) Map(f func(T) T) {
	for _, row := range g {
		for i, x := range row {
			row[i] = f(x)
		}
	}
}

// MapRect applies a transformation function to a sub-rectangle of the grid.
func (g Dense[T]) MapRect(r image.Rectangle, f func(T) T) {
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			g[y][x] = f(g[y][x])
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

// Resize returns a new grid that is r.Dx * r.Dy in size, containing values from
// g. Resize can be used for producing subgrids and supergrids.
func (g Dense[T]) Resize(r image.Rectangle) Dense[T] {
	if r.Empty() {
		return nil
	}
	ng := Make[T](r.Dy(), r.Dx())
	cr := g.Bounds().Intersect(r) // region in g being copied
	for j := cr.Min.Y; j < cr.Max.Y; j++ {
		copy(ng[j-r.Min.Y][cr.Min.X-r.Min.X:], g[j][cr.Min.X:cr.Max.X])
	}
	return ng
}

// BytesFromStrings produces a byte grid from a slice of strings.
func BytesFromStrings(s []string) Dense[byte] {
	g, _ := FromStringsFunc(s, func(s string) ([]byte, error) {
		return []byte(s), nil
	})
	return g
}

// RunesFromStrings produces a byte grid from a slice of strings.
func RunesFromStrings(s []string) Dense[rune] {
	g, _ := FromStringsFunc(s, func(s string) ([]rune, error) {
		return []rune(s), nil
	})
	return g
}

// FromStringsFunc produces a grid from a slice of source strings s,
// containing data for one row per element, and a function for parsing each
// string into a `[]T` row. Uneven length rows are backfilled on the right
// with the zero value for T, to make all rows equal ength.
func FromStringsFunc[T any](s []string, parse func(string) ([]T, error)) (Dense[T], error) {
	if len(s) == 0 {
		return nil, nil
	}
	g := make(Dense[T], len(s))
	w := 0
	for j, row := range s {
		t, err := parse(row)
		if err != nil {
			return nil, fmt.Errorf("parsing row %d: %w", j, err)
		}
		g[j] = t
		if len(t) > w {
			w = len(t)
		}
	}

	// Backfill gaps on right
	for k, t := range g {
		if len(t) == w {
			continue
		}
		g[k] = append(g[k], make([]T, w-len(t))...)
	}
	return g, nil
}

// Freq produces a summary of the number of different values in a grid.
func Freq[T comparable](g Dense[T]) map[T]int {
	h := make(map[T]int)
	for _, row := range g {
		for _, v := range row {
			h[v]++
		}
	}
	return h
}
