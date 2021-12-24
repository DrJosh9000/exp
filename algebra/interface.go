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
// inverse, and multiplication). The ring also has a (multiplicative) identity.
type Ring[T any] interface {
	Add(T, T) T
	Neg(T) T
    Zero() T
	Mul(T, T) T
    Identity() T
}

// DivisionRing is a Ring that also has multiplicative inverses (which can be
// expected to panic on zero). This is the same interface required for Field,
// but has a separate name to remind us that multiplication is not commutative.
type DivisionRing[T any] interface {
	Ring[T]
	Inv(T) T
}

// Field is a commutative division ring. The interface is the same.
type Field[T any] DivisionRing[T]

// VectorSpace describes a vector space (supports addition, scalar
// multiplication, and dot product).
type VectorSpace[V, K any] interface {
	Add(V, V) V
	Neg(V) V
	ScalarMul(K, V) V
	Dot(V, V) K
}
