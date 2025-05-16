/*
   Copyright 2025 Josh Deprez

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

package group

import (
	"slices"
	"strconv"
	"strings"
)

var superscripts = []rune{
	'-': '⁻',
	'0': '⁰', '1': '¹', '2': '²', '3': '³', '4': '⁴',
	'5': '⁵', '6': '⁶', '7': '⁷', '8': '⁸', '9': '⁹',
}

// Sign returns 1 if x > 0, -1 if x < 0, and 0 if x == 0.
func Sign(x int) int {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

// Basic is a generator-power pair.
type Basic struct {
	G rune
	P int
}

// Generator applies "Sign" to b.P.
// Examples:
// - (a⁷).Generator() = a
// - (b⁻⁵).Generator() = b⁻¹
// - (c⁰).Generator() = c⁰
func (b Basic) Generator() Basic {
	return Basic{b.G, Sign(b.P)}
}

// Inverse returns the inverse of b.
func (b Basic) Inverse() Basic {
	return Basic{G: b.G, P: -b.P}
}

func (b Basic) String() string {
	if b.P == 1 {
		return string(b.G)
	}
	p := strconv.Itoa(b.P)
	return string(b.G) + strings.Map(func(r rune) rune { return superscripts[r] }, p)
}

// Word is a string of Basics.
type Word []Basic

// Mul concatenates and then freely reduces.
func Mul(ws ...Word) Word {
	return slices.Concat(ws...).Reduce()
}

// Reduce freely reduces the word. Note that it modifies the original word.
func (w Word) Reduce() Word {
	w = slices.DeleteFunc(w, func(b Basic) bool {
		return b.P == 0
	})
	x := w[:0]
	for _, r := range w {
		if len(x) == 0 {
			x = append(x, r)
			continue
		}
		l := &x[len(x)-1]
		if l.G != r.G {
			x = append(x, r)
			continue
		}
		if l.P+r.P == 0 {
			x = x[:len(x)-1]
			continue
		}
		l.P += r.P
	}
	// Zero out remaining elements.
	for i := len(x); i < len(w); i++ {
		w[i] = Basic{}
	}
	return x
}

// Inverse returns a new word containing the inverse of w.
// It does not require the input to be reduced, nor does it reduce
// its output.
// Example: (aba⁻¹b⁻¹).Inverse() = bab⁻¹a⁻¹.
func (w Word) Inverse() Word {
	x := slices.Clone(w)
	slices.Reverse(x)
	for i := range x {
		x[i].P = -x[i].P
	}
	return x
}

// Rotate returns the word left-rotated by a single generator.
// If the word is a relator for a presentation, this returns an
// equivalent relator.
// It does not require the receiver to be reduced, but reduces
// its output if it is nontrivial.
// Examples:
// - (aba⁻¹b⁻¹).Rotate() = ba⁻¹b⁻¹a
// - (a³ba³).Rotate() = a²ba⁴
func (w Word) Rotate() Word {
	if len(w) <= 1 {
		return w
	}
	// It's easier to reason about this by just using the algebra
	g := w[0].Generator()
	return Mul(Word{g.Inverse()}, w, Word{g})
}

func (w Word) String() string {
	if len(w) == 0 {
		return "1"
	}
	var sb strings.Builder
	for _, gp := range w {
		sb.WriteString(gp.String())
	}
	return sb.String()
}
