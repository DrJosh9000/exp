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

import (
	"testing"

	"drjosh.dev/exp/grid"
	"github.com/google/go-cmp/cmp"
)

func TestMatrixMultiply(t *testing.T) {
	var M Matrix[int, Integer]
	I := M.IdentityMatrix(3)
	m := grid.Dense[int]{
		[]int{4, -2, 7},
		[]int{3, 3, 0},
		[]int{-2, 2, 1},
	}
	if diff := cmp.Diff(M.Mul(I, m), m); diff != "" {
		t.Errorf("I * m != m; diff:\n%s", diff)
	}
	if diff := cmp.Diff(M.Mul(m, I), m); diff != "" {
		t.Errorf("m * I != m; diff:\n%s", diff)
	}
}
