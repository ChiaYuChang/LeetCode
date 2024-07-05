package leetcode

type Q3110 struct{}

func (q Q3110) ScoreOfString(s string) int {
	score := 0
	for i := 0; i < len(s)-1; i++ {
		score += q.abs(int(s[i]) - int(s[i+1]))
	}
	return score
}

func (q Q3110) abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
