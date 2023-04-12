package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

type Q0179 struct{}

func LargestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	s := make([]string, len(nums))
	for i := range s {
		s[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(s, func(i, j int) bool {
		if s[i] == s[j] {
			return false
		}

		var l1, l2 int
		var b1, b2 byte

		l1, l2 = len(s[i]), len(s[j])
		k, k1, k2 := 0, 0, 0
		// 3 vs 333 should return false
		for b1 == b2 && k < l2*l1 {
			if k1 >= l1 {
				k1 -= l1
			}

			if k2 >= l2 {
				k2 -= l2
			}

			b1 = s[i][k1]
			b2 = s[j][k2]
			k, k1, k2 = k+1, k1+1, k2+1
		}
		return b1 > b2
	})

	if s[0] == "0" {
		return "0"
	}
	return strings.Join(s, "")
}
