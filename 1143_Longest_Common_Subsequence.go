package leetcode

type Q1143 struct{}

func (q Q1143) LongestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) > len(text2) {
		text1, text2 = text2, text1
	}

	if len(text1) == 0 {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	prev := make([]int, len(text1)+1)

	for i := range text2 {
		curr := make([]int, len(text1)+1)
		for j := range text1 {
			if text2[i] == text1[j] {
				curr[j+1] = max(max(prev[j]+1, prev[j+1]), curr[j])
			} else {
				curr[j+1] = max(prev[j+1], curr[j])
			}
		}
		prev = curr
	}
	return prev[len(text1)]
}
