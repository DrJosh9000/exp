package algo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRangeSubtract(t *testing.T) {
	tests := []struct {
		r, s Range[int]
		want []Range[int]
	}{
		{NewRange(0, 4), NewRange(5, 7), []Range[int]{{0, 4}}},
		{NewRange(0, 4), NewRange(-5, -1), []Range[int]{{0, 4}}},
		{NewRange(0, 4), NewRange(0, 4), nil},
		{NewRange(0, 4), NewRange(-6, 7), nil},
		{NewRange(0, 4), NewRange(-3, 2), []Range[int]{{3, 4}}},
		{NewRange(0, 4), NewRange(3, 7), []Range[int]{{0, 2}}},
		{NewRange(0, 4), NewRange(2, 2), []Range[int]{{0, 1}, {3, 4}}},
	}

	for _, test := range tests {
		got := RangeSubtract(test.r, test.s)
		if diff := cmp.Diff(got, test.want); diff != "" {
			t.Errorf("RangeSubtract(%v, %v) diff (-got +want):\n%s", test.r, test.s, diff)
		}
	}
}
