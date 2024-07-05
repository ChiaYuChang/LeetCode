package leetcode

import "math"

type Q2256 struct{}

const MAXDIFF int = math.MaxInt

func (q Q2256) abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (q Q2256) MinimumAverageDifference(nums []int) int {
	total := 0
	for _, n := range nums {
		if total+n < total {
			panic("overflow")
		}
		total += n
	}

	minAvg := MAXDIFF
	minPos := 0

	lhs, rhs := 0, total
	for i := 0; i < len(nums)-1; i++ {
		lhs += nums[i]
		rhs -= nums[i]
		if avg := q.abs(lhs/(i+1) - rhs/(len(nums)-i-1)); avg < minAvg {
			minAvg = avg
			minPos = i
		}
	}
	if total/len(nums) < minAvg {
		return len(nums) - 1
	}
	return minPos
}
