package leetcode

import (
	"math"
)

func MyAtoi(s string) int {
	if s == "" {
		return 0
	}

	byte2int := func(i byte) (int, bool) {
		if i < '0' || i > '9' {
			return 0, false
		}
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}[i-'0'], true
	}

	// trim leading space
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			s = s[i:]
			break
		}
	}

	if len(s) < 1 {
		return 0
	}

	sign := true
	if _, ok := byte2int(s[0]); !ok {
		switch s[0] {
		case '-':
			sign = false
		case '+':
			sign = true
		default:
			return 0
		}
		s = s[1:]
	}

	n := 0
	for i := 0; i < len(s); i++ {
		d, ok := byte2int(s[i])
		if ok {
			if sign {
				if math.MaxInt32-10*n-d < 0 {
					return math.MaxInt32
				}
				n = 10*n + d
			} else {
				if math.MinInt32-10*n+d > 0 {
					return math.MinInt32
				}
				n = 10*n - d
			}
		} else {
			return n
		}
	}
	return n
}
