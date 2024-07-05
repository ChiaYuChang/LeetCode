package leetcode

type Q1312 struct{}

func (q Q1312) MinInsertions(s string) int {
	l := len(s)
	if l <= 1 {
		return 0
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	prev := make([]int, l+1)
	for i := range prev {
		prev[i] = i
	}

	// Global alignment
	for i := l - 1; i >= 0; i-- {
		curr := make([]int, l+1)
		curr[0] = prev[0] + 1
		for j := 0; j < l; j++ {
			curr[j+1] = min(curr[j]+1, prev[j+1]+1)
			if s[i] == s[j] {
				curr[j+1] = min(curr[j+1], prev[j])
			}
		}
		prev = curr
	}
	return prev[l] / 2
}
