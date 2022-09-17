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

// Package parse contains string parsing helper functions.
package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// Digits converts a string of decimal digits (0-9) into `[]int`, where
// each element is the value of a digit.
func Digits(s string) ([]int, error) {
	r := make([]int, len(s))
	for i, d := range s {
		if d < '0' || d > '9' {
			return nil, fmt.Errorf("rune %c at pos %d is not a digit [0-9]", d, i)
		}
		r[i] = int(d - '0')
	}
	return r, nil
}

// Ints converts whitespace-separated ints into `[]int`.
func Ints(s string) ([]int, error) {
	fs := strings.Fields(s)
	r := make([]int, len(fs))
	for i, f := range fs {
		y, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		r[i] = y
	}
	return r, nil
}

// Fields returns `strings.Fields(s), nil`.
func Fields(s string) ([]string, error) {
	return strings.Fields(s), nil
}

// Bytes returns `[]byte(s), nil`.
func Bytes(s string) ([]byte, error) {
	return []byte(s), nil
}

// Runes returns `[]rune(s), nil`.
func Runes(s string) ([]rune, error) {
	return []rune(s), nil
}
