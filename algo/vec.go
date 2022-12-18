/*
   Copyright 2022 Josh Deprez

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

package algo

import (
	"image"
	"math"
)

// L1 returns the Manhattan norm of p.
func L1(p image.Point) int {
	return Abs(p.X) + Abs(p.Y)
}

// Linfty returns the L∞ norm of p.
func Linfty(p image.Point) int {
	return Max(Abs(p.X), Abs(p.Y))
}

// ExpandRect expands the Rectangle r to include p.
func ExpandRect(r *image.Rectangle, p image.Point) {
	if r.Min.X > p.X {
		r.Min.X = p.X
	}
	if r.Min.Y > p.Y {
		r.Min.Y = p.Y
	}
	if r.Max.X <= p.X {
		r.Max.X = p.X + 1
	}
	if r.Max.Y <= p.Y {
		r.Max.Y = p.Y + 1
	}
}

// Vec3 is a three-dimensional vector type over E.
type Vec3[E Real] [3]E

// Add returns x+y.
func (x Vec3[E]) Add(y Vec3[E]) Vec3[E] {
	x[0] += y[0]
	x[1] += y[1]
	x[2] += y[2]
	return x
}

// Sub returns x-y.
func (x Vec3[E]) Sub(y Vec3[E]) Vec3[E] {
	x[0] -= y[0]
	x[1] -= y[1]
	x[2] -= y[2]
	return x
}

// Mul returns the scalar product.
func (x Vec3[E]) Mul(k E) Vec3[E] {
	x[0] *= k
	x[1] *= k
	x[2] *= k
	return x
}

// Div returns the scalar product with (1/k).
func (x Vec3[E]) Div(k E) Vec3[E] {
	x[0] /= k
	x[1] /= k
	x[2] /= k
	return x
}

// Dot returns the dot product of x and y.
func (x Vec3[E]) Dot(y Vec3[E]) E {
	return x[0]*y[0] + x[1]*y[1] + x[2]*y[2]
}

// L1 returns the Manhattan norm.
func (x Vec3[E]) L1() E {
	return Abs(x[0]) + Abs(x[1]) + Abs(x[2])
}

// L2 returns the Euclidean norm.
func (x Vec3[E]) L2() float64 {
	x0, x1, x2 := float64(x[0]), float64(x[1]), float64(x[2])
	return math.Sqrt(x0*x0 + x1*x1 + x2*x2)
}

// Linfty returns the L∞ norm.
func (x Vec3[E]) Linfty() E {
	return Max(Abs(x[0]), Abs(x[1]), Abs(x[2]))
}

// AABB3 is a three-dimensional axis-aligned bounding box.
type AABB3[E Real] struct {
	Min, Max Vec3[E]
}

// TODO: various image.Point/image.Rectangle-like methods

// Vec4 is a four-dimensional vector type over E.
type Vec4[E Real] [4]E

// Add returns x+y.
func (x Vec4[E]) Add(y Vec4[E]) Vec4[E] {
	x[0] += y[0]
	x[1] += y[1]
	x[2] += y[2]
	x[3] += y[3]
	return x
}

// Sub returns x-y.
func (x Vec4[E]) Sub(y Vec4[E]) Vec4[E] {
	x[0] -= y[0]
	x[1] -= y[1]
	x[2] -= y[2]
	x[3] += y[3]
	return x
}

// Mul returns the scalar product.
func (x Vec4[E]) Mul(k E) Vec4[E] {
	x[0] *= k
	x[1] *= k
	x[2] *= k
	x[3] *= k
	return x
}

// Div returns the scalar product with (1/k).
func (x Vec4[E]) Div(k E) Vec4[E] {
	x[0] /= k
	x[1] /= k
	x[2] /= k
	x[3] /= k
	return x
}

// Dot returns the dot product of x and y.
func (x Vec4[E]) Dot(y Vec4[E]) E {
	return x[0]*y[0] + x[1]*y[1] + x[2]*y[2] + x[3]*y[3]
}

// L1 returns the Manhattan norm.
func (x Vec4[E]) L1() E {
	return Abs(x[0]) + Abs(x[1]) + Abs(x[2]) + Abs(x[3])
}

// L2 returns the Euclidean norm.
func (x Vec4[E]) L2() float64 {
	x0, x1, x2, x3 := float64(x[0]), float64(x[1]), float64(x[2]), float64(x[3])
	return math.Sqrt(x0*x0 + x1*x1 + x2*x2 + x3*x3)
}

// Linfty returns the L∞ norm.
func (x Vec4[E]) Linfty() E {
	return Max(Abs(x[0]), Abs(x[1]), Abs(x[2]), Abs(x[3]))
}
