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

// Gen applies "Sign" to b.P.
// Examples:
// - (a⁷).Gen() = a
// - (b⁻⁵).Gen() = b⁻¹
// - (c⁰).Gen() = c⁰
func (b Basic) Gen() Basic {
	return Basic{b.G, Sign(b.P)}
}

// Inv returns the inverse of b.
func (b Basic) Inv() Basic {
	return Basic{G: b.G, P: -b.P}
}

// Pow returns b^n.
func (b Basic) Pow(n int) Basic {
	return Basic{G: b.G, P: b.P * n}
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

// Pow returns w^n. It doesn't go out of its way to reduce the result.
func (w Word) Pow(n int) Word {
	if n == 0 || len(w) == 0 {
		return nil
	}
	if len(w) == 1 {
		return Word{w[0].Pow(n)}
	}
	if n < 0 {
		w = w.Inv()
		n = -n
	}
	switch n {
	case 0:
		return nil
	case 1:
		return w
	case 2:
		return Mul(w, w)
	default:
		y := w.Pow(n / 2)
		y = Mul(y, y)
		if n%2 == 1 {
			y = Mul(y, w)
		}
		return y
	}
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

// Inv returns a new word containing the inverse of w.
// It does not require the input to be reduced, nor does it reduce
// its output.
// Example: (aba⁻¹b⁻¹).Inv() = bab⁻¹a⁻¹.
func (w Word) Inv() Word {
	x := slices.Clone(w)
	slices.Reverse(x)
	for i := range x {
		x[i].P = -x[i].P
	}
	return x
}

// LRot returns the word left-rotated by a single generator.
// For an input word W = ab...yz, where a...z are generators,
// W.LRot() = a⁻¹Wa = b...yza.
// If the word is a relator for a presentation, this returns an
// equivalent relator.
// It does not require the receiver to be reduced, but reduces
// its output if the receiver is nontrivial.
// Examples:
// - (aba⁻¹b⁻¹).LRot() = bab⁻¹a
// - (a³ba³).LRot() = a²ba⁴
func (w Word) LRot() Word {
	if len(w) <= 1 {
		return w
	}
	g := w[0].Gen()
	return Mul(Word{g.Inv()}, w, Word{g})
}

// RRot returns the word right-rotated by a single generator.
// For an input word W = ab...yz, where a...z are generators,
// W.RRot() = zWz⁻¹ = zab...y.
// If the word is a relator for a presentation, this returns an
// equivalent relator.
// It does not require the receiver to be reduced, but reduces
// its output if the receiver is nontrivial.
// Examples:
// - (aba⁻¹b⁻¹).RRot() = b⁻¹aba⁻¹
// - (a³ba³).RRot() = a⁴ba²
func (w Word) RRot() Word {
	if len(w) <= 1 {
		return w
	}
	g := w[len(w)-1].Gen()
	return Mul(Word{g}, w, Word{g.Inv()})
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
