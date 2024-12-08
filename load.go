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

// Package exp contains some code that should be considered experimental and
// comes with absolutely no guarantees whatsoever (particularly around
// compatibility, consistency, or functionality).
package exp // import "drjosh.dev/exp"

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"log"
	"os"
	"strconv"
	"strings"

	"drjosh.dev/exp/grid"
)

// Ptr returns a pointer to a variable having the value t.
// (Because you can't just say &true, &42, or &"foo"; you have to put it in a
// variable first.)
func Ptr[T any](t T) *T { return &t }

// Must0 panics if err is not nil.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code (handle errors properly!)
func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

// Must returns t if err is nil. If err is not nil, it panics.
// In other words, Must is a "do-or-die" wrapper for function calls that can
// return a value and an error.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code (handle errors properly!)
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// Must2 returns (t, u) if err is nil. If err is not nil, it panics.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code (handle errors properly!)
func Must2[T, U any](t T, u U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return t, u
}

// Must3 returns (t, u, v) if err is nil. If err is not nil, it panics.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code (handle errors properly!)
func Must3[T, U, V any](t T, u U, v V, err error) (T, U, V) {
	if err != nil {
		panic(err)
	}
	return t, u, v
}

// MustFunc converts a func (S -> (T, error)) into a func (S -> T) that instead
// panics on any error.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code (handle errors properly!)
func MustFunc[S, T any](f func(s S) (T, error)) func(S) T {
	return func(s S) T {
		t, err := f(s)
		if err != nil {
			panic(err)
		}
		return t
	}
}

// MustAtoi parses the string as an integer, or panics.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code (handle errors properly!)
func MustAtoi(s string) int { return Must(strconv.Atoi(s)) }

// MustSscanf parses strings, or panics.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code (handle errors properly!)
func MustSscanf(s, f string, a ...any) { Must(fmt.Sscanf(s, f, a...)) }

// MustCut is a version of strings.Cut that panics if sep is not found within s.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code (handle errors properly!)
func MustCut(s, sep string) (before, after string) {
	before, after, ok := strings.Cut(s, sep)
	if !ok {
		panic(fmt.Sprintf("%q not found in %q", sep, s))
	}
	return before, after
}

// MustForEachLineIn calls cb with each line in the file.
// It uses a bufio.Scanner internally, which can fail on longer lines.
// If an error is encountered, it calls log.Fatal.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code, particularly because the
// logged message may be somewhat unhelpful.
func MustForEachLineIn(path string, cb func(line string)) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("MustForEachLineIn: opening file: %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		cb(sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("MustForEachLineIn: scanner: %v", err)
	}
}

// LinesInFile yields each line in the file.
// It uses a bufio.Scanner internally, which can fail on longer lines.
// If an error is encountered, it calls log.Fatal.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code, particularly because the
// logged message may be somewhat unhelpful.
func LinesInFile(path string) iter.Seq[string] {
	return func(yield func(string) bool) {
		f, err := os.Open(path)
		if err != nil {
			log.Fatalf("MustForEachLineIn: opening file: %v", err)
		}
		defer f.Close()
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			if !yield(sc.Text()) {
				return
			}
		}
		if err := sc.Err(); err != nil {
			log.Fatalf("MustForEachLineIn: scanner: %v", err)
		}
	}
}

// MustReadLines reads the entire file into memory and returns a slice
// containing each line of text (essentially, strings.Split(contents, "\n"), but
// ignoring the final element if it is empty).
// If an error is encountered, it calls log.Fatal.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code, particularly because the
// logged message may be somewhat unhelpful.
func MustReadLines(path string) []string {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("MustReadLines: opening file: %v", err)
	}
	lines := strings.Split(string(b), "\n")
	if n1 := len(lines) - 1; lines[n1] == "" {
		return lines[:n1]
	}
	return lines
}

// MustReadDelimited reads the entire file into memory, splits the contents by
// a delimiter, trims leading and trailing spaces from each component, and
// returns the results as a slice.
// If an error is encountered, it calls log.Fatal.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code, particularly because the
// logged message may be somewhat unhelpful.
func MustReadDelimited(path, delim string) []string {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("MustReadLines: opening file: %v", err)
	}
	parts := strings.Split(string(b), delim)
	for i, l := range parts {
		parts[i] = strings.TrimSpace(l)
	}
	return parts
}

// MustReadInts reads the entire file into memory, splits the contents by the
// delimiter, parses each component as a decimal integer, and returns them as a
// slice.
// If an error is encountered, it calls log.Fatal.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code, particularly because the
// logged message may be somewhat unhelpful.
func MustReadInts(path, delim string) []int {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("MustReadInts: opening file: %v", err)
	}
	parts := strings.Split(string(b), delim)
	if n1 := len(parts) - 1; parts[n1] == "" {
		parts = parts[:n1]
	}
	out := make([]int, len(parts))
	for i, s := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			log.Fatalf("MustReadInts: parsing part %d %q: %v", i, s, err)
		}
		out[i] = n
	}
	return out
}

// MustReadByteGrid reads the entire file into memory and returns the contents
// in the form of a dense byte grid.
func MustReadByteGrid(path string) grid.Dense[byte] {
	return grid.BytesFromStrings(MustReadLines(path))
}

// Fmatchf wraps fmt.Fscanf, reporting whether input was scanned successfully.
func Fmatchf(input io.Reader, format string, into ...any) bool {
	_, err := fmt.Fscanf(input, format, into...)
	return err == nil
}

// Smatchf wraps fmt.Sscanf, reporting whether input was scanned successfully.
func Smatchf(input, format string, into ...any) bool {
	_, err := fmt.Sscanf(input, format, into...)
	return err == nil
}

// NonEmpty returns a copy of s with all non-empty strings.
func NonEmpty[S ~[]string](s S) S {
	t := make(S, 0, len(s))
	for _, x := range s {
		if x != "" {
			t = append(t, x)
		}
	}
	return t
}
