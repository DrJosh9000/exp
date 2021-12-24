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

package algebra

// Grid is a "2D array" generic type.
type Grid[T any] [][]T

// MakeGrid makes a Grid of width w and height h.
func MakeGrid[T any](h, w int) Grid[T] {
	g := make(Grid[T], h)
	for j := range g {
		g[j] = make([]T, w)
	}
	return g
}

// Size returns the width and height of the Grid. If the height is zero the
// width will also be zero.
func (g Grid[T]) Size() (h, w int) {
	h = len(g)
	if h == 0 {
		return 0, 0
	} 
	return len(g[0]), h
}

// Fill fills the Grid with the value v.
func (g Grid[T]) Fill(v T) {
	for _, row := range g {
		for i := range row {
			row[i] = v
		}
	}
}

// Transpose returns a new Grid reflected about the diagonal.
func (g Grid[T]) Transpose() Grid[T] {
	w, h := g.Size()
	ng := MakeGrid[T](h, w)
	for j, row := range g {
		for i := range row {
			ng[i][j] = row[i]
		}
	}
	return ng
}


// FlipHorizontal returns a new Grid flipped horizontally (left becomes right).
func (g Grid[T]) FlipHorizontal() Grid[T] {
	ng := MakeGrid[T](g.Size())
	for j, row := range g {
		for i := range row {
			ng[j][i] = row[len(row)-i-1]
		}
	}
	return ng
}

// FlipVertical returns a new Grid flipped vertically (top becomes bottom).
func (g Grid[T]) FlipVertical() Grid[T] {
	ng := MakeGrid[T](g.Size())
	for j, row := range g {
		for i := range row {
			ng[len(g)-j-1][i] = row[i]
		}
	}
	return ng
}

// Rotate returns a new Grid rotated clockwise by 90 degrees.
func (g Grid[T]) Rotate() Grid[T] {
	w, h := g.Size()
	ng := MakeGrid[T](h, w)
	for j, row := range g {
		for i := range row {
			ng[i][w-j-1] = row[i]
		}
	}
	return ng
}

// DigitGrid converts a grid of digits into a Grid[int].
func DigitGrid(digits []string) Grid[int] {
	h := len(digits)
	if h == 0 {
		return nil
	}
	w := len(digits[0])
	g := MakeGrid[int](h, w)
	for j, row := range digits {
		for i, d := range row {
			g[j][i] = int(d - '0')
		}
	}
	return g
}

// RuneGrid converts a grid of characters into a Grid[rune].
func RuneGrid(runes []string) Grid[rune] {
	h := len(runes)
	if h == 0 {
		return nil
	}
	w := len(runes[0])
	g := MakeGrid[rune](h, w)
	for j, row := range runes {
		for i, r := range row {
			g[j][i] = r
		}
	}
	return g
}
