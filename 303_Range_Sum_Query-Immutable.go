package leetcode

type NumArray struct {
	cumSum []int
}

func NumArrayConstructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{cumSum: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.cumSum[right]
	}
	return this.cumSum[right] - this.cumSum[left-1]
}
