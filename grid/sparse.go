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

package grid

import "image"

// Sparse is a sparse grid - a map from points in 2D to values.
type Sparse[T any] map[image.Point]T

// Clone makes a copy of the grid.
func (g Sparse[T]) Clone() Sparse[T] {
	ng := make(Sparse[T], len(g))
	for p, v := range g {
		ng[p] = v
	}
	return ng
}

// Bounds returns a rectangle containing all points in the sparse grid. It
// operates in O(len(g)) time.
func (g Sparse[T]) Bounds() image.Rectangle {
	if len(g) == 0 {
		return image.Rectangle{}
	}
	var bounds image.Rectangle
	more := false
	for p := range g {
		if !more {
			bounds.Min = p
			bounds.Max = p.Add(image.Pt(1, 1))
			more = true
			continue
		}
		if p.X < bounds.Min.X {
			bounds.Min.X = p.X
		}
		if p.X >= bounds.Max.X {
			bounds.Max.X = p.X + 1
		}
		if p.Y < bounds.Min.Y {
			bounds.Min.Y = p.Y
		}
		if p.Y >= bounds.Max.Y {
			bounds.Max.Y = p.Y + 1
		}
	}
	return bounds
}

// ToDense converts a sparse grid into a dense grid and an offset (of the dense
// grid relative to the original sparse grid).
func (g Sparse[T]) ToDense() (Dense[T], image.Point) {
	if len(g) == 0 {
		return nil, image.Pt(0, 0)
	}
	bounds := g.Bounds()
	d := Make[T](bounds.Dy(), bounds.Dx())
	for p, v := range g {
		p = p.Sub(bounds.Min)
		d[p.Y][p.X] = v
	}
	return d, bounds.Min
}
