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

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{a: 7, b: 3, want: 1},
		{a: 3, b: 7, want: 1},
		{a: 32, b: 16, want: 16},
		{a: 32, b: 17, want: 1},
		{a: 100, b: 55, want: 5},
		{a: 210, b: 11, want: 1},
	}

	for _, test := range tests {
		if got, want := GCD(test.a, test.b), test.want; got != want {
			t.Errorf("GCD(%d, %d) = %d, want %d", test.a, test.b, got, want)
		}
	}
}

func TestXGCD(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{a: 7, b: 3, want: 1},
		{a: 3, b: 7, want: 1},
		{a: 32, b: 16, want: 16},
		{a: 32, b: 17, want: 1},
		{a: 100, b: 55, want: 5},
		{a: 210, b: 11, want: 1},
	}

	for _, test := range tests {
		got, x, y := XGCD(test.a, test.b)
		if got != test.want {
			t.Errorf("XGCD(%d, %d) = (gcd=%d, x=%d, y=%d), want GCD = %d", test.a, test.b, got, x, y, test.want)
		}
		if want := test.a*x + test.b*y; got != want {
			t.Errorf("XGCD(%d, %d) = (gcd=%d, x=%d, y=%d), but a*x + b*y = %d", test.a, test.b, got, x, y, want)
		}
	}
}
