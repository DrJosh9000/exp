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

const sqrt2 = 1.414213562373095

// QR2 implements numbers in the algebraic number field ℚ(√2) (the rationals
// adjoined with √2). ℚ(√2) = {(a + b√2)/c : a,b,c ∈ ℤ}.
type QR2 struct {}

var (
	// ℚ(√2) is a field (over triples of ints).
	_ Field[[3]int] = QR2{}
	// Quaternions over ℚ(√2) are a division ring.
	_ DivisionRing[[4][3]int] = Quaternion[[3]int, QR2]{}
)

// Float returns a float64 representation of x.
func (QR2) Float(x [3]int) float64 {
	return (float64(x[0]) + sqrt2*float64(x[1])) / float64(x[2])
}

func (QR2) Format(x [3]int) string {
	return fmt.Sprintf("(%d+%d√2)/%d", x[0], x[1], x[2])
}

// Canon returns x in a canonical form (reduced fraction with
// positive denominator).
func (QR2) Canon(x [3]int) [3]int {
	if x[2] < 0 {
		x[0], x[1], x[2] = -x[0], -x[1], -x[2]
	}
	d := GCD(x[0], GCD(x[1], x[2]))
	x[0] /= d
	x[1] /= d
	x[2] /= d
	return x
}

// Add returns x+y.
func (q QR2) Add(x, y [3]int) [3]int {
	//   (a1 + b1√2)/c1 + (a2 + b2√2)/c2
	// = (a1/c1 + a2/c2) + (b1/c1 + b2/c2)√2
	// = ((a1c2 + a2c1) + (b1c2 + b2c1)√2)/c1c2
	return q.Canon([3]int{
		0: x[0]*y[2] + y[0]*x[2],
		1: x[1]*y[2] + y[1]*x[2],
		2: x[2] * y[2],
	})
}

// Neg returns -x.
func (QR2) Neg(x [3]int) [3]int {
	return [3]int{-x[0], -x[1], x[2]}
}

// Zero returns the triple representing 0.
func (QR2) Zero() [3]int {
	return [3]int{0, 0, 1}
}

// Mul returns x*y.
func (q QR2) Mul(x, y [3]int) [3]int {
	//   (a1 + b1√2)/c1 * (a2 + b2√2)/c2
	// = (a1 + b1√2)(a2 + b2√2) / c1c2
	// = (a1a2 + a2b1√2 + a1b2√2 + 2b1b2) / c1c2
	// = (a1a2+2b1b2 + (a2b1+a1b2)√2) / c1c2
	return q.Canon([3]int{
		0: x[0]*y[0] + 2*x[1]*y[1],
		1: y[0]*x[1] + x[0]*y[1],
		2: x[2] * y[2],
	})
}

// Inv returns 1/x.
func (q QR2) Inv(x [3]int) [3]int {
	//   c / (a + b√2)
	// = (c / (a + b√2)) * (a - b√2)/(a - b√2)
	// = c(a - b√2) / (a + b√2)(a - b√2)
	// = (ac - bc√2) / (a² - ab√2 + ab√2 - 2b²)
	// = (ac - bc√2) / (a² - 2b²)
	return q.Canon([3]int{
		0: x[0] * x[2],
		1: -x[1] * x[2],
		2: x[0]*x[0] - 2*x[1]*x[1],
	})
}

// Identity returns the triple representing 1.
func (QR2) Identity() [3]int {
	return [3]int{1, 0, 1}
}
