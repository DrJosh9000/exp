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

// Matrix implements matrix operations generically over a field.
type Matrix[T Field[T]] Grid[T]

// MakeMatrix returns a zero matrix of the given height and width.
func MakeMatrix[T Field[T]](h, w int) Matrix[T] {
	return Matrix[T](MakeGrid[T](h, w))
}

// DiagonalMatrix returns a square matrix with all zeros except for a given
// value x used for the main diagonal. This can be used to create the identity
// matrix.
func DiagonalMatrix[T Field[T]](n int, x T) Matrix[T] {
	m := MakeMatrix[T](n, n)
	for i := 0; i < n; i++ {
		m[i][i] = x
	}
	return m
}

// Size returns the dimensions of the matrix (number of rows and columns).
func (m Matrix[T]) Size() (h, w int) { return Grid[T](m).Size() }

// Transpose returns the matrix reflected across the main diagonal.
func (m Matrix[T]) Transpose() Matrix[T] { return Matrix[T](Grid[T](m).Transpose())}

// Add adds two matrices.
func (m Matrix[T]) Add(n Matrix[T]) Matrix[T] {
	mh, mw := m.Size()
	nh, nw := n.Size()
	if mh != nh || mw != nw {
		panic("adding matrices of incompatible sizes")
	}
	for j := range m {
		for i := range m[j] {
			m[j][i] = m[j][i].Add(n[j][i])
		}
	}
	return m
}

// Neg returns the matrix with all entries negated.
func (m Matrix[T]) Neg() Matrix[T] {
	for j := range m {
		for i := range m[j] {
			m[j][i] = m[j][i].Neg()
		}
	}
	return m
}

// ScalarMul returns the matrix with entries multiplied by k.
func (m Matrix[T]) ScalarMul(k T) Matrix[T] {
	for j := range m {
		for i := range m[j] {
			m[j][i] = k.Mul(m[j][i])
		}
	}
	return m
}

// Mul multiplies two matrices of compatible size (the width of m must equal
// the height of n)
func (m Matrix[T]) Mul(n Matrix[T]) Matrix[T] {
	mh, mw := m.Size()
	nh, nw := n.Size()
	if mw != nh {
		panic("multiplying matrices of incompatible sizes")
	}
	o := MakeMatrix[T](mh, nw)
	for i := range o {
		for j := range o[i] {
			for k := 0; k < mw; k++ {
				o[i][j] = o[i][j].Add(m[i][k].Mul(n[k][j]))
			}
		}
	}
	return o
}

// TODO: inverses