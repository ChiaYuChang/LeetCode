package leetcode

type Q1047 struct{}

func (q Q1047) RemoveDuplicates(s string) string {
	alwayNotMatchChar := byte('*') // a char that does not match any char in s

	bs := make([]byte, len(s))
	end := 0
	for prev, i := alwayNotMatchChar, 0; i < len(s); i++ {
		if s[i] != prev {
			bs[end] = s[i]
			prev = bs[end]
			end++
		} else {
			if end > 0 {
				end--
				if end > 0 {
					prev = bs[end-1]
				} else {
					prev = alwayNotMatchChar
				}
			}
		}
	}
	return string(bs[:end])
}
