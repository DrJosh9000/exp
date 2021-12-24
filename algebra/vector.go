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

// Vector[float64, Real] is a vector space over the reals.
var _ VectorSpace[[]float64, float64] = Vector[float64, Real]{}

// Vector implements a vector space for vectors of type []T and operations from
// any ring R.
type Vector[T any, R Ring[T]] struct{}

// Add returns the vector sum v+w.
func (Vector[T, R]) Add(v, w []T) []T {
    if len(v) != len(w) {
        panic("mismatching vector sizes")
    }
    var r R
    o := make([]T, len(v))
	for i := range o {
        o[i] = r.Add(v[i], w[i])
    }
    return o
}

// Neg returns -v (the vector pointing in the opposite direction).
func (Vector[T, R]) Neg(v []T) []T {
    o := make([]T, len(v))
    var r R
	for i := range o {
        o[i] = r.Neg(v[i])
    }
    return o
}

// ScalarMul returns kv.
func (Vector[T, R]) ScalarMul(k T, v []T) []T {
    o := make([]T, len(v))
    var r R
	for i := range o {
        o[i] = r.Mul(k, v[i])
    }
    return o
}

// Dot returns the dot product (also known as inner product) v.w.
func (Vector[T, R]) Dot(v, w []T) T {
    if len(v) != len(w) {
        panic("mismatching vector sizes")
    }
    var s T
    var r R
	for i := range v {
        s = r.Add(s, r.Mul(v[i], w[i]))
    }
    return s
}

/*
// Quaternion returns v as a vector quaternion (a quaternion having zero scalar
// component).
func (v Vector[T]) Quaternion() Quaternion[T] {
	return Quaternion[T]{1: v.X, 2: v.Y, 3: v.Z}
}
*/