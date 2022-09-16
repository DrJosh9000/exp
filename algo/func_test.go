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

package algo

import "testing"

func TestFoldl(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}

	tests := []struct {
		op     func(int, int) int
		opName string
		want   int
	}{
		{
			op:     func(x, y int) int { return x + y },
			opName: "+",
			want:   21,
		},
		{
			op:     func(x, y int) int { return x * y },
			opName: "*",
			want:   720,
		},
		{
			op:     func(x, y int) int { return x - y },
			opName: "-",
			want:   -19,
		},
		{
			op:     func(x, y int) int { return x ^ y },
			opName: "^",
			want:   7,
		},
	}

	for _, test := range tests {
		if got := Foldl(nums, test.op); got != test.want {
			t.Errorf("Foldl(nums, %s) = %d, want %d", test.opName, got, test.want)
		}
	}
}

func TestFoldr(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}

	tests := []struct {
		op     func(int, int) int
		opName string
		want   int
	}{
		{
			op:     func(x, y int) int { return x + y },
			opName: "+",
			want:   21,
		},
		{
			op:     func(x, y int) int { return x * y },
			opName: "*",
			want:   720,
		},
		{
			op:     func(x, y int) int { return x - y },
			opName: "-",
			want:   -9,
		},
		{
			op:     func(x, y int) int { return x ^ y },
			opName: "^",
			want:   7,
		},
	}

	for _, test := range tests {
		if got := Foldr(nums, test.op); got != test.want {
			t.Errorf("Foldr(nums, %s) = %d, want %d", test.opName, got, test.want)
		}
	}
}
