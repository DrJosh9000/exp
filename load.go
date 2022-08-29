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
package exp // import "github.com/DrJosh9000/exp"

import (
	"bufio"
	"log"
	"os"
)

// Must returns t if err is nil. If err is not nil, it calls log.Fatal.
// In other words, Must is a "do-or-die" wrapper for function calls that can
// return an error.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code (handle your errors yourself!)
// particularly because the logged message will be unhelpful.
func Must[T any](t T, err error) T {
	if err != nil {
		log.Fatalf("Must: %v", err)
	}
	return t
}

// MustFunc converts a func that can return errors into a func that
// calls log.Fatal on any error.
// This is a helper intended for very simple programs (e.g. Advent of Code)
// and is not recommended for production code (handle your errors yourself!)
// particularly because the logged message will be unhelpful.
func MustFunc[S, T any](f func(s S) (T, error)) func(S) T {
	return func(s S) T {
		t, err := f(s)
		if err != nil {
			log.Fatalf("MustFunc: %v", err)
		}
		return t
	}
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
		log.Fatalf("MustForEachLineIn opening file: %v", err)
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

