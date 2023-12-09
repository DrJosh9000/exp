package para

import (
	"testing"

	"github.com/DrJosh9000/exp/algo"
	"github.com/google/go-cmp/cmp"
)

func TestMap(t *testing.T) {
	for N := 0; N < 300; N++ {
		in := make([]int, N)
		f := func(x int) int { return x + 1 }
		got := Map(in, f)
		want := algo.Map(in, f)
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("Map([]int of len %d) diff (-got +want):\n%s", len(in), diff)
		}
	}
}
