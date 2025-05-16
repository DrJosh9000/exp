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
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWordString(t *testing.T) {
	tests := []struct {
		input Word
		want  string
	}{
		{Word{}, "1"},
		{Word{{'a', 1}}, "a"},
		{Word{{'a', 1}, {'a', -1}}, "aa⁻¹"},
		{Word{{'b', -1}, {'b', 1}}, "b⁻¹b"},
		{Word{{'a', 2}, {'a', -1}}, "a²a⁻¹"},
		{Word{{'a', 1}, {'b', 1}}, "ab"},
		{Word{{'a', 1}, {'b', 1}, {'b', -1}, {'a', -1}}, "abb⁻¹a⁻¹"},
		{Word{{'a', 2}, {'b', -3}, {'b', 3}, {'a', -2}}, "a²b⁻³b³a⁻²"},
	}

	for _, test := range tests {
		got := test.input.String()
		if diff := cmp.Diff(got, test.want); diff != "" {
			t.Errorf("(%#v).String() = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestWordReduce(t *testing.T) {
	tests := []struct {
		input, want Word
	}{
		{Word{}, Word{}},
		{Word{{'a', 1}}, Word{{'a', 1}}},
		{Word{{'a', 1}, {'a', 1}}, Word{{'a', 2}}},
		{Word{{'a', 1}, {'a', -1}}, Word{}},
		{Word{{'b', -1}, {'b', 1}}, Word{}},
		{Word{{'a', 2}, {'a', -1}}, Word{{'a', 1}}},
		{Word{{'a', 1}, {'b', 1}}, Word{{'a', 1}, {'b', 1}}},
		{Word{{'a', 1}, {'b', 1}, {'b', -1}, {'a', -1}}, Word{}},
		{Word{{'a', 2}, {'b', -3}, {'b', 3}, {'a', -2}}, Word{}},
		{
			input: Word{{'a', 1}, {'b', 1}, {'a', -1}, {'b', -1}},
			want:  Word{{'a', 1}, {'b', 1}, {'a', -1}, {'b', -1}},
		},
		{
			input: Word{{'a', 2}, {'b', -3}, {'a', 3}, {'b', -2}},
			want:  Word{{'a', 2}, {'b', -3}, {'a', 3}, {'b', -2}},
		},
	}

	for _, test := range tests {
		got := slices.Clone(test.input)
		got = got.Reduce()
		if diff := cmp.Diff(got, test.want); diff != "" {
			t.Errorf("(%v).Reduce() = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestWordInverse(t *testing.T) {
	tests := []struct {
		input, want Word
	}{
		{Word{}, Word{}},
		{Word{{'a', 1}}, Word{{'a', -1}}},
		{Word{{'a', 1}, {'a', 1}}, Word{{'a', -1}, {'a', -1}}},
		{Word{{'a', 1}, {'a', -1}}, Word{{'a', 1}, {'a', -1}}},
		{Word{{'a', 1}, {'b', -1}}, Word{{'b', 1}, {'a', -1}}},
		{Word{{'a', -1}, {'b', 1}}, Word{{'b', -1}, {'a', 1}}},
		{Word{{'a', 2}, {'b', -1}}, Word{{'b', 1}, {'a', -2}}},
		{Word{{'a', 1}, {'b', 1}}, Word{{'b', -1}, {'a', -1}}},
		{
			input: Word{{'a', 1}, {'b', 1}, {'a', -1}, {'b', -1}},
			want:  Word{{'b', 1}, {'a', 1}, {'b', -1}, {'a', -1}},
		},
		{
			input: Word{{'a', 2}, {'b', -3}, {'a', 3}, {'b', -2}},
			want:  Word{{'b', 2}, {'a', -3}, {'b', 3}, {'a', -2}},
		},
	}

	for _, test := range tests {
		got := test.input.Inv()
		if diff := cmp.Diff(got, test.want); diff != "" {
			t.Errorf("(%v).Inverse() = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestWordRotate(t *testing.T) {
	tests := []struct {
		input, want Word
	}{
		{Word{}, Word{}},
		{Word{{'a', 1}}, Word{{'a', 1}}},
		{Word{{'a', 1}, {'a', 1}}, Word{{'a', 2}}},
		{Word{{'a', 1}, {'a', -1}}, Word{}},
		{Word{{'a', 1}, {'b', -1}}, Word{{'b', -1}, {'a', 1}}},
		{Word{{'a', -1}, {'b', 1}}, Word{{'b', 1}, {'a', -1}}},
		{Word{{'a', 2}, {'b', -1}}, Word{{'a', 1}, {'b', -1}, {'a', 1}}},
		{Word{{'a', 1}, {'b', 1}}, Word{{'b', 1}, {'a', 1}}},
		{
			input: Word{{'a', 1}, {'b', 1}, {'a', -1}, {'b', -1}},
			want:  Word{{'b', 1}, {'a', -1}, {'b', -1}, {'a', 1}},
		},
		{
			input: Word{{'a', 2}, {'b', -3}, {'a', 3}, {'b', -2}},
			want:  Word{{'a', 1}, {'b', -3}, {'a', 3}, {'b', -2}, {'a', 1}},
		},
	}

	for _, test := range tests {
		got := test.input.LRot()
		if diff := cmp.Diff(got, test.want); diff != "" {
			t.Errorf("(%v).Rotate() = %v, want %v", test.input, got, test.want)
		}
	}
}
