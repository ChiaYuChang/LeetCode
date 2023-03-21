package leetcode

import "strconv"

type Q0443 struct{}

func (q Q0443) StringCompression(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	curC := 0
	count := 1
	n := 0
	for j := 1; j < len(chars); j++ {
		if chars[j] == chars[curC] {
			count++
		} else {
			chars[n] = chars[curC]
			n++
			if count > 1 {
				countStr := strconv.Itoa(count)
				for k := range countStr {
					chars[n] = countStr[k]
					n++
				}
			}
			curC = j
			count = 1
		}
	}

	// append the last character
	chars[n] = chars[curC]
	n++
	if count > 1 {
		countStr := strconv.Itoa(count)
		for k := range countStr {
			chars[n] = countStr[k]
			n++
		}
	}
	return n
}
