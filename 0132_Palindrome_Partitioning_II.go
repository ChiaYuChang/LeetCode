package leetcode

type Q0132 struct{}

func (q Q0132) MinCut(s string) int {
	if len(s) == 0 {
		return 0
	}

	// Dynamic Programming
	prev := make([]int, len(s)+1)
	cuts := make([]int, len(s)+1)
	for i, j := len(s), -1; i >= 0; i-- {
		cuts[i], j = j, j+1
	}

	last := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		curr := make([]int, last+1)
		// margin
		if s[i] == s[0] {
			curr[0] = 1
		}
		for j := 1; j <= last && j < len(s); j++ {
			if s[i] == s[j] && j > 0 {
				curr[j] = prev[j-1] + 1
			}
		}

		if curr[last] > 0 {
			for d := curr[last]; d > 1; d-- {
				cuts[last-d+1] = q.Min(cuts[last-d+1], cuts[last+d-1]+1)
			}
		}

		if curr[last-1] > 0 {
			for d := curr[last-1]; d > 0; d-- {
				cuts[last-d] = q.Min(cuts[last-d], cuts[last+d-1]+1)
			}
		}
		prev = curr
		last--
	}

	return cuts[0]
}

func (q Q0132) Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
