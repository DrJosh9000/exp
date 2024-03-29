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

func TestPowInt(t *testing.T) {
	want := 23
	for i := uint(1); i < 13; i++ {
		if got := Pow(23, i, func(x, y int) int { return x * y }); got != want {
			t.Errorf("Pow(23, %d, *) = %d, want %d", i, got, want)
		}
		want *= 23
	}
}
