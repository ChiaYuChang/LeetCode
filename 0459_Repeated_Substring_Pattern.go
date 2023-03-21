package leetcode

type Q0459 struct{}

func (q Q0459) RepeatedSubstringPattern(s string) bool {
	pre := make([]int, len(s))
	cur := make([]int, len(s))

	max := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				cur[j] = pre[j-1] + 1
				if max < cur[j] {
					max = cur[j]
				}
			}
		}
		pre = cur
		cur = make([]int, len(s))
		if max > 1 {
			return true
		}
	}
	return false
}
