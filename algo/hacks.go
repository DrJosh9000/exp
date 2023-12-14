package algo

// CyclicPredict predicts what would appear at slice index n if the input slice
// were that long. It does this by linearly searching for a cycle at the end.
func CyclicPredict[S ~[]E, E comparable](s S, n int) E {
	if n < 0 {
		panic("negative input provided to CyclicPredict")
	}
	if n < len(s) {
		return s[n]
	}
	seen := make(map[E]int)
	for i := range s {
		j := len(s) - i - 1
		x := s[j]
		if k, yes := seen[x]; yes {
			cyc := k - j
			return s[len(s)-cyc:][(n-len(s))%cyc]
		}
		seen[x] = j
	}
	panic("no cycle detected")
}
