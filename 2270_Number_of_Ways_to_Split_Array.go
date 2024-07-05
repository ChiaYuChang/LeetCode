package leetcode

type Q2270 struct{}

func (q Q2270) WaysToSplitArray(nums []int) int {
	LHS, RHS := 0, 0
	for _, num := range nums {
		RHS += num
	}

	count := 0
	for i := 0; i < len(nums)-1; i++ {
		LHS += nums[i]
		RHS -= nums[i]
		if LHS >= RHS {
			count++
		}
	}
	return count
}
