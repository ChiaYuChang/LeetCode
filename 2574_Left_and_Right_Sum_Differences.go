package leetcode

type Q2574 struct{}

func (q Q2574) LeftRigthDifference(nums []int) []int {
	if len(nums) == 1 {
		return []int{0}
	}

	var i, total, tmp int
	for _, n := range nums {
		total += n
	}

	i, tmp, nums[0] = 1, nums[0], total-nums[0]
	for ; i < len(nums); i++ {
		nums[i], tmp = nums[i-1]-tmp-nums[i], nums[i]
		if nums[i-1] < 0 {
			nums[i-1] = -nums[i-1]
		}
	}

	if nums[i-1] < 0 {
		nums[i-1] = -nums[i-1]
	}
	return nums
}
