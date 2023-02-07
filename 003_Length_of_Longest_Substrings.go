package leetcode

func LengthOfLongestSubstring(s string) int {
	maxSubStrLen := 0
	charset := make(map[rune]int, 0)

	leftEnd := 0
	for rightEnd, r := range s {
		index, ok := charset[r]
		if ok {
			subStrLen := rightEnd - leftEnd
			if subStrLen > maxSubStrLen {
				maxSubStrLen = subStrLen
			}
			if index+1 > leftEnd {
				leftEnd = index + 1
			}
		}
		charset[r] = rightEnd
	}

	// the end point to the last repeated char index
	if len(s)-leftEnd > maxSubStrLen {
		return len(s) - leftEnd
	}
	return maxSubStrLen
}
