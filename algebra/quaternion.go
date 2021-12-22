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
	_ DivisionRing[Quaternion[Real]] = Quaternion[Real]{}
	// ℍ is also a vector space over the reals.
	_ Vector[Quaternion[Real], Real] = Quaternion[Real]{}
)

// Quaternion implements quaternions generically as a division algebra over some 
// other division ring. Traditional quaternions (ℍ) use ℝ.
type Quaternion[T DivisionRing[T]] [4]T

func (q Quaternion[T]) String() string { 
	return fmt.Sprint("%v + %v𝕚 + %v𝕛 + %v𝕜", q[0], q[1], q[2], q[3]) 
}

func (q Quaternion[T]) Neg() Quaternion[T] {
	return Quaternion[T]{q[0].Neg(), q[1].Neg(), q[2].Neg(), q[3].Neg()}
}

func (q Quaternion[T]) Add(r Quaternion[T]) Quaternion[T] {
	return Quaternion[T]{q[0].Add(r[0]), q[1].Add(r[1]), q[2].Add(r[2]), q[3].Add(r[3])}
}

// Conjugate returns the quaternion conjugate. This is equal to the inverse
// for rotation quaternions (those with norm 1).
func (q Quaternion[T]) Conjugate() Quaternion[T] {
	return Quaternion[T]{q[0], q[1].Neg(), q[2].Neg(), q[3].Neg()}
}

func (q Quaternion[T]) Dot(r Quaternion[T]) T {
	return q[0].Mul(r[0]).Add(q[1].Mul(r[1])).Add(q[2].Mul(r[2])).Add(q[3].Mul(r[3]))
}

// Inv returns the inverse quaternion.
func (q Quaternion[T]) Inv() Quaternion[T] {
	return q.Conjugate().ScalarMul(q.Dot(q).Inv())
}

// ScalarMul multiplies q by a scalar.
func (q Quaternion[T]) ScalarMul(x T) Quaternion[T] {
	return Quaternion[T]{q[0].Mul(x), q[1].Mul(x), q[2].Mul(x), q[3].Mul(x)}
}

// Mul returns the quaternion product qr.
func (q Quaternion[T]) Mul(r Quaternion[T]) Quaternion[T] {
	return Quaternion[T]{
		q[0].Mul(r[0]).Add(q[1].Mul(r[1]).Neg()).Add(q[2].Mul(r[2]).Neg()).Add(q[3].Mul(r[3]).Neg()),
		q[0].Mul(r[1]).Add(q[1].Mul(r[0])).Add(q[2].Mul(r[3])).Add(q[3].Mul(r[2]).Neg()),
		q[0].Mul(r[2]).Add(q[1].Mul(r[3]).Neg()).Add(q[2].Mul(r[0])).Add(q[3].Mul(r[1])),
		q[0].Mul(r[3]).Add(q[1].Mul(r[2])).Add(q[2].Mul(r[1]).Neg()).Add(q[3].Mul(r[0])),
	}
}

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
