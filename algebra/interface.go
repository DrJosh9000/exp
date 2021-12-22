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

// Ring describes a ring (algebraic structure that supports addition, additive
// inverse, and multiplication).
type Ring[T any] interface {
	Add(T) T
	Neg() T
	Mul(T) T
}

// DivisionRing is a Ring that also has multiplicative inverses (which can be
// expected to panic on the additive identity). This is the same interface
// required for Field, but has a separate name to remind us that in general
// b⁻¹a != ab⁻¹.
type DivisionRing[T any] interface {
	Ring[T]
	Inv() T
}

// Field is a commutative division ring
type Field[T any] DivisionRing[T]

// Vector describes a vector space (supports addition, scalar multiplication,
// and dot product).
type Vector[V, K any] interface {
	Add(V) V
	Neg() V
	ScalarMul(K) V
	Dot(V) K
}