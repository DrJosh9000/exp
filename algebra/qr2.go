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
// adjoined with √2). It can be shown that ℚ(√2) = {(a + b√2)/c : a,b,c ∈ ℤ}.
// All arithmetic here is done with ints.
type QR2 struct {
	A, B, C int
}

var (
	// Common QR2 values.
	QR2MinusOne        = QR2{-1, 0, 1}
	QR2MinusRoot2Over2 = QR2{0, -1, 2} // equal to -1/√2
	QR2Zero            = QR2{0, 0, 1}
	QR2Root2Over2      = QR2{0, 1, 2} // equal to 1/√2
	QR2One             = QR2{1, 0, 1}

	// ℚ(√2) is a field.
	_ Field[QR2] = QR2{}
	// Quaternions over ℚ(√2) are a division ring.
	_ DivisionRing[Quaternion[QR2]] = Quaternion[QR2]{}
)

// Float returns a float64 representation of x.
func (x QR2) Float() float64 {
	return (float64(x.A) + sqrt2*float64(x.B)) / float64(x.C)
}

func (x QR2) String() string {
	return fmt.Sprintf("(%d+%d√2)/%d", x.A, x.B, x.C)
}

// Canon returns x in a canonical form (reduced fraction with
// positive denominator).
func (x QR2) Canon() QR2 {
	if x.C < 0 {
		x.A, x.B, x.C = -x.A, -x.B, -x.C
	}
	d := GCD(x.A, GCD(x.B, x.C))
	x.A /= d
	x.B /= d
	x.C /= d
	return x
}

// Neg returns -x.
func (x QR2) Neg() QR2 {
	return QR2{-x.A, -x.B, x.C}
}

// Add returns x+y.
func (x QR2) Add(y QR2) QR2 {
	//   (a1 + b1√2)/c1 + (a2 + b2√2)/c2
	// = (a1/c1 + a2/c2) + (b1/c1 + b2/c2)√2
	// = ((a1c2 + a2c1) + (b1c2 + b2c1)√2)/c1c2
	return QR2{
		A: x.A*y.C + y.A*x.C,
		B: x.B*y.C + y.B*x.C,
		C: x.C * y.C,
	}.Canon()
}

// Inv returns 1/x.
func (x QR2) Inv() QR2 {
	//   c / (a + b√2)
	// = (c / (a + b√2)) * (a - b√2)/(a - b√2)
	// = c(a - b√2) / (a + b√2)(a - b√2)
	// = (ac - bc√2) / (a² - ab√2 + ab√2 - 2b²)
	// = (ac - bc√2) / (a² - 2b²)
	return QR2{
		A: x.A * x.C,
		B: -x.B * x.C,
		C: x.A*x.A - 2*x.B*x.B,
	}.Canon()
}

// Mul returns x*y.
func (x QR2) Mul(y QR2) QR2 {
	//   (a1 + b1√2)/c1 * (a2 + b2√2)/c2
	// = (a1 + b1√2)(a2 + b2√2) / c1c2
	// = (a1a2 + a2b1√2 + a1b2√2 + 2b1b2) / c1c2
	// = (a1a2+2b1b2 + (a2b1+a1b2)√2) / c1c2
	return QR2{
		A: x.A*y.A + 2*x.B*y.B,
		B: y.A*x.B + x.A*y.B,
		C: x.C * y.C,
	}.Canon()
}


