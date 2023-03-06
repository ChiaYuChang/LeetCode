package leetcode

import (
	"math"
	"sort"
)

func MinimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}

	sort.Sort(sort.IntSlice(nums))
	diff, mindiff := 0, math.MaxInt

	for i := len(nums) - 1; i <= k; i-- {
		diff = nums[i-k] - nums[i]
		if diff < mindiff {
			mindiff = diff
		}
	}
	return mindiff
}
