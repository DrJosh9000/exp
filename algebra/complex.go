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

// ℂ is a field.
var _ Field[Complex[Real]] = Complex[Real]{}

// Complex implements complex numbers generically over any division ring.
type Complex[T DivisionRing[T]] struct {
	X, Y T
}

func (z Complex[T]) String() string { 
	return fmt.Sprint("%v + %v𝕚", z.X, z.Y) 
}

// Add returns z+w.
func (z Complex[T]) Add(w Complex[T]) Complex[T] { 
	return Complex[T]{z.X.Add(w.X), z.Y.Add(w.Y)}
}

// Neg returns -z.
func (z Complex[T]) Neg() Complex[T] { 
	return Complex[T] {z.X.Neg(), z.Y.Neg()}
}

// Mul returns z*w.
func (z Complex[T]) Mul(w Complex[T]) Complex[T] { 
	return Complex[T]{
		X: z.X.Mul(w.X).Add(z.Y.Mul(w.Y).Neg()), 
		Y: z.X.Mul(w.Y).Add(z.Y.Mul(w.X)),
	}
}

// Inv returns 1/z.
func (z Complex[T]) Inv() Complex[T] { 
	d := z.X.Mul(z.X).Add(z.Y.Mul(z.Y)).Inv()
	return Complex[T]{z.X.Mul(d), z.Y.Neg().Mul(d)}
}

// Conjugate returns the conjugate of z.
func (z Complex[T]) Conjugate() Complex[T] {
	return Complex[T]{z.X, z.Y.Neg()}
}