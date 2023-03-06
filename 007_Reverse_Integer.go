package leetcode

import "math"

func ReverseInt(n int) int {
	if n == 0 {
		return 0
	}

	sign := n > 0
	revn := 0

	if !sign {
		n = -n
	}

	for n > 0 {
		if sign {
			// positive integer
			if math.MaxInt32-revn*10-n%10 < 0 {
				// 2147483647
				// check overflow
				return 0
			}
			revn = revn*10 + n%10
		} else {
			// negative  integer
			if math.MinInt32-revn*10+n%10 > 0 {
				// -2147483648
				// check overflow
				return 0
			}
			revn = revn*10 - n%10
		}
		n = n / 10
	}
	return revn
}
