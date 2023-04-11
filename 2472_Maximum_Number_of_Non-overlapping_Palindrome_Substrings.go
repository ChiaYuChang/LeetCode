package leetcode

type Q2472 struct{}

func (q Q2472) MaxPalindromes(s string, k int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	if len(s) == 0 {
		return 0
	}

	// Dynamic Programming
	prev := make([]int, len(s)+1)
	cuts := make([]int, len(s)+1)
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
				l, r := last-d+1, last+d-1
				if r-l >= k {
					// only palindromes longer than k could be added
					cuts[l] = max(cuts[l], cuts[r]+1)
				}
			}
		}

		if curr[last-1] > 0 {
			for d := curr[last-1]; d > 0; d-- {
				l, r := last-d, last+d-1
				if r-l >= k {
					// only palindromes longer than k could be added
					cuts[l] = max(cuts[l], cuts[r]+1)
				}
			}
		}

		// cut[i] = max(cut[i], cut[i+1], ... cut[n])
		cuts[i] = max(cuts[i], cuts[i+1])
		prev = curr
		last--
	}

	return cuts[0]
}
