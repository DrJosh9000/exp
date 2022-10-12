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
// implementation (one or more lines of Go) and a set of jump targets (line
// numbers). A TranslatorFunc should not produce labels; Transpile takes care
// of inserting necessary labels.
type TranslatorFunc func(line int, args []string) (impl string, jumpTargets []int, err error)

// Translate translates a program into a Go implementation and a set of jump
// targets.
func Translate(program []string, translators map[string]TranslatorFunc) ([]string, algo.Set[int], error) {
	targets := make(algo.Set[int])
	out := make([]string, len(program))
	for lno, line := range program {
		ls := strings.Fields(line)
		tl, ok := translators[ls[0]]
		if !ok {
			return nil, nil, fmt.Errorf("unknown opcode %q on line %d", ls[0], lno)
		}
		impl, jts, err := tl(lno, ls[1:])
		if err != nil {
			return nil, nil, fmt.Errorf("translating line %d: %w", lno, err)
		}
		for _, jt := range jts {
			targets[jt] = struct{}{}
		}
		out[lno] = impl
	}
	return out, targets, nil
}

// Transpile transpiles a program using a set of opcode translators, and returns
// a func natively implementing the program.
func Transpile(program []string, translators map[string]TranslatorFunc) (RunFunc, error) {

	// --- Translate first before writing files or invoking the compiler. --- \\
	impl, jts, err := Translate(program, translators)
	if err != nil {
		return nil, err
	}

	// --- Write the temporary Go file. --- \\
	f, err := os.CreateTemp("", "emu*.go")
	if err != nil {
		return nil, fmt.Errorf("creating temp file: %w", err)
	}
	fname := f.Name()
	defer os.Remove(fname)
	bf := bufio.NewWriter(f)
	if _, err := bf.WriteString(preamble); err != nil {
		return nil, fmt.Errorf("writing preamble: %w", err)
	}
	for lno, line := range impl {
		if jts.Contains(lno) {
			if _, err := fmt.Fprintf(bf, "l%d:\n", lno); err != nil {
				return nil, fmt.Errorf("writing label %d: %w", lno, err)
			}
		}
		if _, err := fmt.Fprintf(bf, "\t%s\n", line); err != nil {
			return nil, fmt.Errorf("writing line %d: %w", lno, err)
		}
	}
	if _, err := bf.WriteString(postamble); err != nil {
		return nil, fmt.Errorf("writing postamble: %w", err)
	}
	if err := bf.Flush(); err != nil {
		return nil, fmt.Errorf("flushing buffer: %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("closing file: %w", err)
	}

	// --- Compile to a plugin --- \\
	soname := fname + ".so"
	cmd := exec.Command("go", "build", "-o", soname, "-buildmode=plugin", fname)
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
