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

import "golang.org/x/exp/constraints"

// Numeric types have any of the built-in numeric types as the underlying type.
type Numeric interface {
	Real | constraints.Complex
}

// Addable types have any of the built-in types that support the + operator
// as the underlying type.
type Addable interface {
	Numeric | ~string
}

// Real types have any of the built-in integer or float types (but not complex)
// as the underlying type.
type Real interface {
	constraints.Integer | constraints.Float
}
