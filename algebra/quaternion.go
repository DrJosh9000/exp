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
	// ℍ is a division ring.
	_ DivisionRing[[4]float64] = Quaternion[float64, Real]{}
	// ℍ is also a vector space over the reals.
	_ VectorSpace[[4]float64, float64] = Quaternion[float64, Real]{}
)

// Quaternion implements quaternions generically as an algebra over some
// other ring. Traditional quaternions (ℍ) use ℝ.
type Quaternion[T any, R Ring[T]] struct{}

// Format formats x into a string.
func (Quaternion[T, R]) Format(x [4]T) string {
	return fmt.Sprintf("%v + %v𝕚 + %v𝕛 + %v𝕜", x[0], x[1], x[2], x[3])
}

// Add returns x+y.
func (Quaternion[T, R]) Add(x, y [4]T) [4]T {
	var r R
	return [4]T{
		r.Add(x[0], y[0]),
		r.Add(x[1], y[1]),
		r.Add(x[2], y[2]),
		r.Add(x[3], y[3]),
	}
}

// Neg returns -x.
func (Quaternion[T, R]) Neg(x [4]T) [4]T {
	var r R
	return [4]T{r.Neg(x[0]), r.Neg(x[1]), r.Neg(x[2]), r.Neg(x[3])}
}

// Zero returns 0 + 0𝕚 + 0𝕛 + 0𝕜
func (Quaternion[T, R]) Zero() [4]T {
	var r R
	return [4]T{r.Zero(), r.Zero(), r.Zero(), r.Zero()}
}

// Identity returns 1 + 0𝕚 + 0𝕛 + 0𝕜
func (Quaternion[T, R]) Identity() [4]T {
	var r R
	return [4]T{r.Identity(), r.Zero(), r.Zero(), r.Zero()}
}

// Conjugate returns the quaternion conjugate. This is equal to the inverse
// for rotation quaternions (those with norm 1).
func (Quaternion[T, R]) Conjugate(x [4]T) [4]T {
	var r R
	return [4]T{x[0], r.Neg(x[1]), r.Neg(x[2]), r.Neg(x[3])}
}

// Dot returns the dot product of q with r (treating them as 4D vectors).
func (Quaternion[T, R]) Dot(x, y [4]T) T {
	var r R
	s := r.Mul(x[0], y[0])
	s = r.Add(s, r.Mul(x[1], y[1]))
	s = r.Add(s, r.Mul(x[2], y[2]))
	s = r.Add(s, r.Mul(x[3], y[3]))
	return s
}

// Inv returns x⁻¹, or panics if R is not a division ring or x.x has no inverse.
func (q Quaternion[T, R]) Inv(x [4]T) [4]T {
	var r R
	dr := any(r).(DivisionRing[T])
	d := dr.Inv(q.Dot(x, x))
	return q.ScalarMul(d, q.Conjugate(x))
}

// ScalarMul returns k*x.
func (Quaternion[T, R]) ScalarMul(k T, x [4]T) [4]T {
	var r R
	return [4]T{r.Mul(k, x[0]), r.Mul(k, x[1]), r.Mul(k, x[2]), r.Mul(k, x[3])}
}

// Mul returns the quaternion product qr.
func (Quaternion[T, R]) Mul(x, y [4]T) [4]T {
	var r R
	var z [4]T
	z[0] = r.Mul(x[0], y[0])
	z[0] = r.Add(z[0], r.Neg(r.Mul(x[1], y[1])))
	z[0] = r.Add(z[0], r.Neg(r.Mul(x[2], y[2])))
	z[0] = r.Add(z[0], r.Neg(r.Mul(x[3], y[3])))
	z[1] = r.Mul(x[0], y[1])
	z[1] = r.Add(z[1], r.Mul(x[1], y[0]))
	z[1] = r.Add(z[1], r.Mul(x[2], y[3]))
	z[1] = r.Add(z[1], r.Neg(r.Mul(x[3], y[2])))
	z[2] = r.Mul(x[0], y[2])
	z[2] = r.Add(z[2], r.Neg(r.Mul(x[1], y[3])))
	z[2] = r.Add(z[2], r.Mul(x[2], y[0]))
	z[2] = r.Add(z[2], r.Mul(x[3], y[1]))
	z[3] = r.Mul(x[0], y[3])
	z[3] = r.Add(z[3], r.Mul(x[1], y[2]))
	z[3] = r.Add(z[3], r.Neg(r.Mul(x[2], y[1])))
	z[3] = r.Add(z[3], r.Mul(x[3], y[0]))
	return z
}

/*
// Vec3 returns the vector component of q.
func (q Quaternion[T]) Vec3() Vec3[T] {
	return Vec3[T]{q[1], q[2], q[3]}
}

// Rotate returns the conjugate product qvq⁻¹. If q is a rotation quaternion
// q = cos(θ/2) + (uᵢ𝕚 + uⱼ𝕛 + uₖ𝕜)*sin(θ/2)
// then Rotate rotates v by the angle θ about the axis (uᵢ,uⱼ,uₖ).
func (q Quaternion[T]) Rotate(v Vec3[T]) Vec3[T] {
	return q.Mul(v.Quaternion()).Mul(q.Conjugate()).Vec3()
}
*/
