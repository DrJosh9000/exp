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

// Pow raises base to the power pow, where multiplication is given by op.
// The only requirements are that:
// * op must be 3-power associative for base (i.e. 
//   `op(op(base,base),base) == op(base,op(base,base))`), and
// * pow must be strictly positive (pow >= 1).
// op is called O(log(pow) + bits.OnesCount(pow)) times.
//
// If you are working with big numbers, use math/big.
// 
// Note that, for this implementation, op need *not* define a monoid (i.e. 
// have an identity element in T) or even be generally associative. 3-power
// associativity implies commutativity among powers of a value.
// For Pow to be able to handle a pow == 0 in full generality would require
// an additional parameter or more convention-setting. But you can easily
// test for pow == 0 yourself before calling.
func Pow[T any](base T, pow uint, op func(T, T) T) T {
	if pow < 1 {
		panic("pow must be at least 1")
	}
	if pow == 1 {
		return base
	}
	var accum T
	ini := false
	for {
		if pow % 2 == 1 {
			if ini {
				accum = op(accum, base)
			} else {
				accum, ini = base, true
			}
		}
		pow /= 2
		if pow < 1 {
			return accum
		}
		base = op(base, base)
	}
}