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

import "testing"

func TestQR2Canon(t *testing.T) {
	tests := []struct {
		in, want [3]int
	}{
		{in: [3]int{0, 0, 0}, want: [3]int{0, 0, 1}},
		{in: [3]int{0, 0, 1}, want: [3]int{0, 0, 1}},
		{in: [3]int{1, 0, 1}, want: [3]int{1, 0, 1}},
		{in: [3]int{0, 1, 1}, want: [3]int{0, 1, 1}},
		{in: [3]int{1, 1, 1}, want: [3]int{1, 1, 1}},
		{in: [3]int{0, 0, 2}, want: [3]int{0, 0, 1}},
		{in: [3]int{2, 0, 2}, want: [3]int{1, 0, 1}},
		{in: [3]int{2, 2, 2}, want: [3]int{1, 1, 1}},
		{in: [3]int{0, 2, 2}, want: [3]int{0, 1, 1}},
		{in: [3]int{1, 1, -1}, want: [3]int{-1, -1, 1}},
		{in: [3]int{2, 2, -2}, want: [3]int{-1, -1, 1}},
	}

	var r QR2
	for _, test := range tests {
		if got, want := r.Canon(test.in), test.want; got != want {
			t.Errorf("QR2{}.Canon(%v) = %v, want %v", r.Format(test.in), r.Format(got), r.Format(want))
		}
	}
}

func TestQR2Add(t *testing.T) {
	tests := []struct {
		x, y, want [3]int
	}{
		{x: [3]int{0, 0, 1}, y: [3]int{0, 0, 1}, want: [3]int{0, 0, 1}},
		{x: [3]int{0, 0, 1}, y: [3]int{1, 0, 1}, want: [3]int{1, 0, 1}},
		{x: [3]int{1, 0, 1}, y: [3]int{0, 1, 1}, want: [3]int{1, 1, 1}},
		{x: [3]int{0, 1, 1}, y: [3]int{0, -1, 1}, want: [3]int{0, 0, 1}},
		{x: [3]int{1, 1, 1}, y: [3]int{1, 1, 1}, want: [3]int{2, 2, 1}},
		{x: [3]int{1, 0, 2}, y: [3]int{1, 0, 2}, want: [3]int{1, 0, 1}},
		{x: [3]int{1, 0, 2}, y: [3]int{0, 1, 2}, want: [3]int{1, 1, 2}},
		{x: [3]int{3, 0, 1}, y: [3]int{0, 7, 1}, want: [3]int{3, 7, 1}},
		{x: [3]int{3, 7, 2}, y: [3]int{5, 2, 3}, want: [3]int{19, 25, 6}},
	}

	var r QR2
	for _, test := range tests {
		if got, want := r.Add(test.x, test.y), test.want; got != want {
			t.Errorf("QR2{}.Add(%v, %v) = %v, want %v", r.Format(test.x), r.Format(test.y), r.Format(got), r.Format(want))
		}
	}
}
func TestQR2Mul(t *testing.T) {
	tests := []struct {
		x, y, want [3]int
	}{
		{x: [3]int{0, 0, 1}, y: [3]int{0, 0, 1}, want: [3]int{0, 0, 1}},
		{x: [3]int{0, 0, 1}, y: [3]int{1, 0, 1}, want: [3]int{0, 0, 1}},
		{x: [3]int{1, 0, 1}, y: [3]int{0, 1, 1}, want: [3]int{0, 1, 1}},
		{x: [3]int{0, 1, 1}, y: [3]int{0, -1, 1}, want: [3]int{-2, 0, 1}},
		{x: [3]int{1, 1, 1}, y: [3]int{1, 1, 1}, want: [3]int{3, 2, 1}},
		{x: [3]int{1, 0, 2}, y: [3]int{1, 0, 2}, want: [3]int{1, 0, 4}},
		{x: [3]int{1, 0, 2}, y: [3]int{0, 1, 2}, want: [3]int{0, 1, 4}},
		{x: [3]int{0, 1, 2}, y: [3]int{0, 1, 2}, want: [3]int{1, 0, 2}},
		{x: [3]int{3, 0, 1}, y: [3]int{0, 7, 1}, want: [3]int{0, 21, 1}},
		{x: [3]int{3, 7, 2}, y: [3]int{5, 2, 3}, want: [3]int{43, 41, 6}},
	}

	var r QR2
	for _, test := range tests {
		if got, want := r.Mul(test.x, test.y), test.want; got != want {
			t.Errorf("QR2{}.Mul(%v, %v) = %v, want %v", r.Format(test.x), r.Format(test.y), r.Format(got), r.Format(want))
		}
	}
}
