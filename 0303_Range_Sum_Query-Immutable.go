package leetcode

type Q0303 struct{}

type Q0303NumArray struct {
	cumSum []int
}

func (q Q0303) NumArrayConstructor(nums []int) Q0303NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return Q0303NumArray{cumSum: nums}
}

func (this *Q0303NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.cumSum[right]
	}
	return this.cumSum[right] - this.cumSum[left-1]
}
