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

// Matrix implements matrix algebra on Grid[T] given a ring R.
type Matrix[T any, R Ring[T]] struct{}

// ZeroMatrix returns a zero matrix of size h x w.
func ZeroMatrix[T any, R Ring[T]](h, w int) Grid[T] {
	var r R
	m := MakeGrid[T](h, w)
	for i := range m {
		for j := range m[i] {
			m[i][j] = r.Zero()
		}
	}
	return m
}

// IdentityMatrix returns an identity matrix of size n.
func IdentityMatrix[T any, R Ring[T]](n int) Grid[T] {
	var r R
	m := MakeGrid[T](n, n)
	for i := range m {
		for j := range m[i] {
			if i == j {
				m[i][j] = r.Identity()
			} else {
				m[i][j] = r.Zero()
			}
		}
	}
	return m
}

// Add adds two matrices.
func (Matrix[T, R]) Add(m, n Grid[T]) Grid[T] {
	h, w := m.Size()
	nh, nw := n.Size()
	if h != nh || w != nw {
		panic("adding matrices of incompatible sizes")
	}
	o := MakeGrid[T](h, w)
	var r R
	for i := range o {
		for j := range o[i] {
			o[i][j] = r.Add(m[i][j], n[i][j])
		}
	}
	return o
}

// Neg returns the matrix with all entries negated.
func (Matrix[T, R]) Neg(m Grid[T]) Grid[T] {
	var r R
	o := MakeGrid[T](m.Size())
	for i := range o {
		for j := range o[i] {
			o[i][j] = r.Neg(m[i][j])
		}
	}
	return o
}

// ScalarMul returns the matrix with entries multiplied by k.
func (Matrix[T, R]) ScalarMul(k T, m Grid[T]) Grid[T] {
	var r R
	o := MakeGrid[T](m.Size())
	for i := range o {
		for j := range o[i] {
			o[i][j] = r.Mul(k, m[i][j])
		}
	}
	return o
}

// Mul multiplies two matrices of compatible size (the width of m must equal
// the height of n)
func (Matrix[T, R]) Mul(m, n Grid[T]) Grid[T] {
	mh, mw := m.Size()
	nh, nw := n.Size()
	if mw != nh {
		panic("multiplying matrices of incompatible sizes")
	}
	var r R
	o := ZeroMatrix[T, R](mh, nw)
	for i := range o {
		for j := range o[i] {
			for k := 0; k < mw; k++ {
				o[i][j] = r.Add(o[i][j], r.Mul(m[i][k], n[k][j]))
			}
		}
	}
	return o
}


/*
// Inv returns the matrix inverse, or panics if T is not a division ring or
// the matrix is singular.
func (m Matrix[T]) Inv() Matrix[T] {
	h, w := m.Size()
	if h == 0 || w == 0 {
		return m
	}
	if h != w {
		panic("inverting non-square matrix")
	}

}
*/