package leetcode

type Q1539 struct{}

func (q Q1539) FindKthMissingPositive(arr []int, k int) int {
	lhs, rhs := 0, len(arr)-1 // Left Hand Side, Right Hand Side

	// greather than the last element
	if k > arr[rhs]-len(arr) {
		return k + len(arr)
	}

	// less than the first element
	if k <= arr[lhs]-1 {
		return k
	}

	// binary search
	for rhs > lhs {
		mid := (lhs + rhs) / 2
		if arr[mid]-mid-1 < k {
			lhs = mid + 1
		} else {
			rhs = mid - 1
		}
	}

	// rhs should be equal to lhs
	// arr[lhs]-1: the k  th missing interger
	// arr[lhs]+1: the k+1th missing interger
	shft := k - (arr[lhs] - lhs - 1)
	if shft <= 0 {
		shft--
	}
	return arr[lhs] + shft
}
