package leetcode

type Q0516 struct{}

func (q Q0516) LongestPalindromeSubseq(s string) int {
	max := func(nums ...int) int {
		m := nums[0]
		for i := 1; i < len(nums); i++ {
			if m < nums[i] {
				m = nums[i]
			}
		}
		return m
	}

	prev := make([]int, len(s)+1)
	for row := len(s) - 1; row >= 0; row-- {
		curr := make([]int, len(s)+1)
		for col := 0; col < len(s); col++ {
			if s[col] == s[row] {
				curr[col+1] = max(prev[col]+1, curr[col], prev[col+1])
			} else {
				curr[col+1] = max(curr[col], prev[col+1])
			}
		}
		prev = curr
	}
	return prev[len(s)]
}
