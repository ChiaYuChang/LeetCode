package leetcode

type Q0724 struct{}

func (q Q0724) PivotIndex(nums []int) int {
	LHS, RHS := 0, 0
	for _, num := range nums {
		RHS += num
	}

	for i, num := range nums {
		RHS -= num
		if RHS == LHS {
			return i
		}
		LHS += num
	}
	return -1
}
