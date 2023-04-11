package leetcode

type Q1745 struct{}

func (q Q1745) CheckPartitioning(s string) bool {
	if len(s) == 0 {
		return false
	}

	// Dynamic Programming
	k := 3                        // partition number
	cuts := make([]int, len(s)+1) // one padding element
	cuts[len(s)] = 1 << k

	prev := make([]int, len(s)+1)
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

		// even length
		if curr[last] > 0 {
			for d := curr[last]; d > 1; d-- {
				cuts[last-d+1] |= cuts[last+d-1] >> 1
			}
		}

		// odd length
		if curr[last-1] > 0 {
			for d := curr[last-1]; d > 0; d-- {
				cuts[last-d] |= cuts[last+d-1] >> 1
			}
		}
		prev = curr
		last--
	}
	return (cuts[0] & 0b0001) > 0
}
