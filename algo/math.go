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

// Abs returns the absolute value of x (with no regard for negative overflow).
//
// The math/cmplx package provides a version of Abs for complex types.
func Abs[T Real](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Max returns the argument closest to infinity. If no arguments are provided,
// Max returns 0.
func Max[T Real](x ...T) T {
	if len(x) == 0 {
		return 0
	}
	m := x[0]
	for _, t := range x[1:] {
		if t > m {
			m = t
		}
	}
	return m
}

// Min returns the argument closest to negative infinity. If no arguments are
// provided, Min returns 0.
func Min[T Real](x ...T) T {
	if len(x) == 0 {
		return 0
	}
	m := x[0]
	for _, t := range x[1:] {
		if t < m {
			m = t
		}
	}
	return m
}
