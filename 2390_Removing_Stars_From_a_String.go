package leetcode

import (
	"strings"
)

type Q2390 struct{}

func (q Q2390) RemoveStars(s string) string {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	s = strings.TrimLeft(s, "*")
	bs := make([]byte, len(s)-strings.Count(s, "*"))

	end := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			bs[end], end = s[i], end+1
		} else {
			end = max(0, end-1)
		}
	}
	return string(bs[:end])
}
