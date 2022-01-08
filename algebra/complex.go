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

import "fmt"

var (
	// ℂ (as Complex[float64, Real]) is a field.
	_ Field[[2]float64] = Complex[float64, Real]{}
	// ℂ is also a 2-dimensional vector space over ℝ.
	_ VectorSpace[[2]float64, float64] = Complex[float64, Real]{}

	// ℤ[𝕚] is a ring.
	_ Ring[[2]int] = GaussianInteger{}
	// ℤ[𝕚] is also a 2-dimensional vector space over ℤ.
	_ VectorSpace[[2]int, int] = GaussianInteger{}

	// Nothing stopping me making triply-nested complexes.
	// I'm mad with power!
	_ Field[[2][2]complex128] = Complex[[2]complex128, Complex[complex128, Cmplx]]{}
)

// GaussianInteger implements the integral domain ℤ[𝕚] using int.
type GaussianInteger = Complex[int, Integer]

// Complex implements complex numbers generically as a 2-dimensional algebra
// over any ring.
type Complex[T any, R Ring[T]] struct{}

// Format formats a complex number.
func (Complex[T, R]) Format(z [2]T) string {
	return fmt.Sprintf("%v + %v𝕚", z[0], z[1])
}

// Add returns z+w.
func (Complex[T, R]) Add(z, w [2]T) [2]T {
	var r R
	return [2]T{
		r.Add(z[0], w[0]), 
		r.Add(z[1], w[1]),
	}
}

// Neg returns -z.
func (Complex[T, R]) Neg(z [2]T) [2]T {
	var r R
	return [2]T{r.Neg(z[0]), r.Neg(z[1])}
}

// Zero returns 0 + 0𝕚.
func (Complex[T, R]) Zero() [2]T {
	var r R
	return [2]T{r.Zero(), r.Zero()}
}

// Mul returns z*w.
func (Complex[T, R]) Mul(z, w [2]T) [2]T {
	var r R
	return [2]T{
		r.Add(r.Mul(z[0], w[0]), r.Neg(r.Mul(z[1], w[1]))),
		r.Add(r.Mul(z[0], w[1]), r.Mul(z[1], w[0])),
	}
}

// Identity returns 1 + 0𝕚.
func (Complex[T, R]) Identity() [2]T {
	var r R
	return [2]T{r.Identity(), r.Zero()}
}

// Inv returns 1/z, or panics if R is not a division ring or is zero.
func (c Complex[T, R]) Inv(z [2]T) [2]T {
	var r R
	dr := any(r).(DivisionRing[T])
	d := dr.Inv(c.Dot(z, z))
	return c.ScalarMul(d, c.Conjugate(z))
}

// Conjugate returns the conjugate of z.
func (Complex[T, R]) Conjugate(z [2]T) [2]T {
	var r R
	return [2]T{z[0], r.Neg(z[1])}
}

// Dot returns Re(z)*Re(w) + Im(z)*Im(w).
func (Complex[T, R]) Dot(z, w [2]T) T {
	var r R
	return r.Add(r.Mul(z[0], w[0]), r.Mul(z[1], w[1]))
}

// ScalarMul returns k*z.
func (Complex[T, R]) ScalarMul(k T, z [2]T) [2]T {
	var r R
	return [2]T{r.Mul(k, z[0]), r.Mul(k, z[1])}
}
