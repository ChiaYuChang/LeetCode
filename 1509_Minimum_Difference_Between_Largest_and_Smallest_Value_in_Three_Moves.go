package leetcode

import (
	"math"
	"sort"
)

type Q1509 struct{}

func (q Q1509) min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (q Q1509) MinDifference(nums []int) int {
	if len(nums) <= 3 {
		return 0
	}
	sort.Ints(nums)

	minDiff := math.MaxInt
	for i, j := 0, len(nums)-4; j < len(nums); i, j = i+1, j+1 {
		minDiff = q.min(minDiff, nums[j]-nums[i])
	}
	return minDiff
}
