package algo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAllPairs(t *testing.T) {
	want := [][2]int{
		{1, 2}, {1, 3}, {2, 3},
	}
	var got [][2]int
	for a, b := range AllPairs([]int{1, 2, 3}) {
		got = append(got, [2]int{a, b})
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("AllPairs diff (-got, +want):\n%s", diff)
	}
}
