package leetcode

type Q0704 struct{}

func (q Q0704) Search(nums []int, target int) int {
	lhs, rhs, mid := 0, len(nums)-1, 0
	for rhs >= lhs {
		mid = (rhs + lhs) >> 1
		if nums[mid] > target {
			rhs = mid - 1
		} else if target > nums[mid] {
			lhs = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
