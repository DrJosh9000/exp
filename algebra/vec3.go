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

// Vec3[Real] is a vector space over the reals.
var _ Vector[Vec3[Real], Real] = Vec3[Real]{}

// Vec3 implements a 3-dimensional vector generically, over any division ring.
type Vec3[T Ring[T]] struct {
	X, Y, Z T
}

func (v Vec3[T]) String() string {
	return fmt.Sprintf("(%v,%v,%v)", v.X, v.Y, v.Z)
}

// Add returns the vector sum v+w.
func (v Vec3[T]) Add(w Vec3[T]) Vec3[T] {
	return Vec3[T]{v.X.Add(w.X), v.Y.Add(w.Y), v.Z.Add(w.Z)}
}

// Neg returns -v (the vector pointing in the opposite direction).
func (v Vec3[T]) Neg() Vec3[T] {
	return Vec3[T]{v.X.Neg(), v.Y.Neg(), v.Z.Neg()}
}

// ScalarMul returns kv.
func (v Vec3[T]) ScalarMul(k T) Vec3[T] {
	return Vec3[T]{k.Mul(v.X), k.Mul(v.Y), k.Mul(v.Z)}
}

// Dot returns the dot product (also known as inner product) v.w.
func (v Vec3[T]) Dot(w Vec3[T]) T {
	return v.X.Mul(w.X).Add(v.Y.Mul(w.Y)).Add(v.Z.Mul(w.Z))
}

// Quaternion returns v as a vector quaternion (a quaternion having zero scalar
// component).
func (v Vec3[T]) Quaternion() Quaternion[T] {
	return Quaternion[T]{1: v.X, 2: v.Y, 3: v.Z}
}
