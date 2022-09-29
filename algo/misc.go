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

import "image"

var (
	// Neigh4 is a slice containing a step up, right, down, and left
	// (N, E, S, W).
	Neigh4 = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	// Neigh8 is a slice containing a step to all 8 neighbouring cells in order
	// from top-left to bottom-right.
	Neigh8 = []image.Point{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0} /*0, 0*/, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
)
