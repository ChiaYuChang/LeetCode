package leetcode

type Q0583 struct{}

func (q Q0583) MinDistance(word1 string, word2 string) int {
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
	}

	if len(word1) == 0 {
		return len(word2)
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	prev := make([]int, len(word1)+1)
	for i := range prev {
		prev[i] = i
	}

	for i := range word2 {
		curr := make([]int, len(word1)+1)
		curr[0] = prev[0] + 1
		for j := range word1 {
			if word2[i] == word1[j] {
				curr[j+1] = min(min(prev[j], prev[j+1]+1), curr[j]+1)
			} else {
				curr[j+1] = min(prev[j+1]+1, curr[j]+1)
			}
		}
		prev = curr
	}
	return prev[len(word1)]
}
