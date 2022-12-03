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

// Package emu provides a generalised virtual machine for executing toy
// programs that operate solely on integers. It is implemented using a kind of
// degenerate JIT compilation. Instructions are translated into Go, which is
// compiled into a plugin, which is then loaded using the plugin package.
//
// This code is not security-hardened in any way. Using this code is a bad idea.
package emu

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"plugin"
	"strings"

	"github.com/DrJosh9000/exp/algo"
)

const preamble = `package main

func Run(r, m []int, send func(int) error, recv func() (int, error)) error {
`

const postamble = "\treturn nil\n}\n"

// RunFunc is the type of the function that is compiled.
type RunFunc = func(r, m []int, send func(int) error, recv func() (int, error)) error

// TranslatorFunc translates an input line number and line of code into an
// implementation (one or more lines of Go) and a set of absolute jump targets
// (line numbers). A TranslatorFunc should not produce labels; Transpile takes
// care of inserting necessary labels. However, a TranslatorFunc must implement
// its own jumps (e.g. `goto l%d`).
type TranslatorFunc func(line int, args []string) (impl string, jumpTargets []int, err error)

// Translate translates a program into a Go implementation and writes it to w.
func Translate(w io.Writer, program []string, translators map[string]TranslatorFunc) error {
	if _, err := fmt.Fprint(w, preamble); err != nil {
		return fmt.Errorf("writing preamble: %w", err)
	}

	targets := make(algo.Set[int])
	lines := make([]string, len(program))
	for lno, line := range program {
		ls := strings.Fields(line)
		tl, ok := translators[ls[0]]
		if !ok {
			return fmt.Errorf("unknown opcode %q on line %d", ls[0], lno)
		}
		impl, jts, err := tl(lno, ls[1:])
		if err != nil {
			return fmt.Errorf("translating line %d: %w", lno, err)
		}
		for _, jt := range jts {
			targets[jt] = struct{}{}
		}
		lines[lno] = impl
	}

	for lno, line := range lines {
		if targets.Contains(lno) {
			if _, err := fmt.Fprintf(w, "l%d:\n", lno); err != nil {
				return fmt.Errorf("writing label %d: %w", lno, err)
			}
		}
		if _, err := fmt.Fprintf(w, "\t%s\n", line); err != nil {
			return fmt.Errorf("writing line %d: %w", lno, err)
		}
	}
	// Jump past the end?
	if end := len(lines); targets.Contains(end) {
		if _, err := fmt.Fprintf(w, "l%d:\n", end); err != nil {
			return fmt.Errorf("writing label %d: %w", end, err)
		}
	}

	if _, err := fmt.Fprint(w, postamble); err != nil {
		return fmt.Errorf("writing postamble: %w", err)
	}
	return nil
}

// Transpile transpiles a program using a set of opcode translators, and returns
// a func natively implementing the program.
func Transpile(program []string, translators map[string]TranslatorFunc) (RunFunc, error) {
	// --- Write the temporary Go file. --- \\
	f, err := os.CreateTemp("", "emu*.go")
	if err != nil {
		return nil, fmt.Errorf("creating temp file: %w", err)
	}
	fname := f.Name()
	defer os.Remove(fname)
	bf := bufio.NewWriter(f)
	if err := Translate(bf, program, translators); err != nil {
		return nil, fmt.Errorf("translating program: %w", err)
	}
	if err := bf.Flush(); err != nil {
		return nil, fmt.Errorf("flushing buffer: %w", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("closing file: %w", err)
	}

	// --- Compile to a plugin --- \\
	soname := fname + ".so"
	cmd := exec.Command("go", "build", "-o", soname, "-buildmode=plugin", fname)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("compiling temporary file: %w", err)
	}
	defer os.Remove(soname)

	// --- Open the plugin and find the entry point --- \\
	p, err := plugin.Open(soname)
	if err != nil {
		return nil, fmt.Errorf("opening plugin: %w", err)
	}
	rf, err := p.Lookup("Run")
	if err != nil {
		return nil, fmt.Errorf("looking up Run in plugin: %w", err)
	}
	r, ok := rf.(RunFunc)
	if !ok {
		return nil, fmt.Errorf("symbol Run has bad type %T", rf)
	}
	return r, nil
}
