package leetcode

type Q2259 struct{}

func (q Q2259) RemoveDigit(number string, digit byte) string {
	maxSubStr := ""
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			substr := number[:i] + number[i+1:]
			if maxSubStr == "" || substr > maxSubStr {
				maxSubStr = substr
			}
		}
	}
	return maxSubStr
}
