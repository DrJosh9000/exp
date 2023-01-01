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
	"image"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTranspose(t *testing.T) {
	g := Dense[int]{
		[]int{0, 1, 2, 3},
		[]int{3, 4, 5, 6},
		[]int{6, 7, 8, 9},
	}
	want := Dense[int]{
		[]int{0, 3, 6},
		[]int{1, 4, 7},
		[]int{2, 5, 8},
		[]int{3, 6, 9},
	}
	got := g.Transpose()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("g.Transpose diff:\n%s", diff)
	}
}

func TestFlipH(t *testing.T) {
	g := Dense[int]{
		[]int{0, 1, 2, 3},
		[]int{3, 4, 5, 6},
		[]int{6, 7, 8, 9},
	}
	want := Dense[int]{
		[]int{3, 2, 1, 0},
		[]int{6, 5, 4, 3},
		[]int{9, 8, 7, 6},
	}
	got := g.FlipH()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("g.FlipH diff:\n%s", diff)
	}
}

func TestFlipV(t *testing.T) {
	g := Dense[int]{
		[]int{0, 1, 2, 3},
		[]int{3, 4, 5, 6},
		[]int{6, 7, 8, 9},
	}
	want := Dense[int]{
		[]int{6, 7, 8, 9},
		[]int{3, 4, 5, 6},
		[]int{0, 1, 2, 3},
	}
	got := g.FlipV()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("g.FlipV diff:\n%s", diff)
	}
}

func TestRotateCW(t *testing.T) {
	g := Dense[int]{
		[]int{0, 1, 2, 3},
		[]int{3, 4, 5, 6},
		[]int{6, 7, 8, 9},
	}
	want := Dense[int]{
		[]int{6, 3, 0},
		[]int{7, 4, 1},
		[]int{8, 5, 2},
		[]int{9, 6, 3},
	}
	got := g.RotateCW()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("g.RotateCW diff:\n%s", diff)
	}
}

func TestRotateACW(t *testing.T) {
	g := Dense[int]{
		[]int{0, 1, 2, 3},
		[]int{3, 4, 5, 6},
		[]int{6, 7, 8, 9},
	}
	want := Dense[int]{
		[]int{3, 6, 9},
		[]int{2, 5, 8},
		[]int{1, 4, 7},
		[]int{0, 3, 6},
	}
	got := g.RotateACW()
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("g.RotateACW diff:\n%s", diff)
	}
}

func TestResize(t *testing.T) {
	g := Dense[int]{
		[]int{0, 1, 2, 3},
		[]int{3, 4, 5, 6},
		[]int{6, 7, 8, 9},
	}

	tests := []struct {
		r    image.Rectangle
		want Dense[int]
	}{
		{
			r:    image.Rect(1, 1, 1, 1),
			want: nil,
		},
		{
			r: image.Rect(1, 1, 3, 2),
			want: Dense[int]{
				[]int{4, 5},
			},
		},
		{
			r:    image.Rect(0, 0, 4, 3),
			want: g,
		},
		{
			r: image.Rect(-1, -1, 5, 4),
			want: Dense[int]{
				[]int{0, 0, 0, 0, 0, 0},
				[]int{0, 0, 1, 2, 3, 0},
				[]int{0, 3, 4, 5, 6, 0},
				[]int{0, 6, 7, 8, 9, 0},
				[]int{0, 0, 0, 0, 0, 0},
			},
		},
		{
			r: image.Rect(-1, -1, 3, 2),
			want: Dense[int]{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 1, 2},
				[]int{0, 3, 4, 5},
			},
		},
		{
			r: image.Rect(1, 1, 5, 4),
			want: Dense[int]{
				[]int{4, 5, 6, 0},
				[]int{7, 8, 9, 0},
				[]int{0, 0, 0, 0},
			},
		},
	}

	for _, test := range tests {
		if diff := cmp.Diff(g.Resize(test.r), test.want); diff != "" {
			t.Errorf("g.Resize(%v) diff:\n%s", test.r, diff)
		}
	}
}
