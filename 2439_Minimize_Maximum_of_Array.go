package leetcode

type Q2439 struct{}

func (q Q2439) MinimizeArrayValue(nums []int) int {
	for i := 1; i < len(nums); i++ {
		// cumulated sum
		nums[i] = nums[i] + nums[i-1]

		// average to nums[i]
		avgToI := nums[i] / (i + 1)
		if nums[i]%(i+1) != 0 {
			avgToI++
		}

		if avgToI > nums[0] {
			// store the minmax value at nums[0]
			nums[0] = avgToI
		}
	}
	return nums[0]
}
