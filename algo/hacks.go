package algo

import (
	"golang.org/x/exp/constraints"
)

// IntegerPredict predicts what would appear at slice index n if the input slice
// were that long. It does this by searching for a cycle (modulo a fixed diff)
// and then extrapolating.
func IntegerPredict[S ~[]E, E constraints.Integer](s S, n int) E {
	if n < 0 {
		panic("negative input provided to IntegerPredict")
	}
	if n < len(s) {
		return s[n]
	}
	ls1 := len(s) - 1
	var period int
	var slope E
periodLoop:
	for period = len(s) / 2; period > 0; period-- {
		slope = s[ls1] - s[ls1-period]
		for i := 1; i < period; i++ {
			j := ls1 - i
			if dy := s[j] - s[j-period]; dy != slope {
				continue periodLoop
			}
		}
		break periodLoop
	}
	if period <= 0 {
		panic("no cycle detected")
	}
	//log.Printf("found period = %d, slope = %d", period, slope)
	rem := n - len(s)
	return slope*E(1+rem/period) + s[len(s)-period:][rem%period]
}
