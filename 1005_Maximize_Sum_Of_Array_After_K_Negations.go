package leetcode

import (
	"sort"
)

type Q1005 struct{}

func (q Q1005) LargestSumAfterKNegations(nums []int, k int) int {
	nums = append(nums, 0)
	sort.Sort(sort.IntSlice(nums))

	zeroLHS, zeroRHS := q.SearchLeftMost(nums, 0), q.SearchRightMost(nums, 0)

	// neg----------            pos-----------
	// [ -7, -3, -1] , 0, 0, 0, [ 1, 3, 6, 8 ]
	//                 ^     ^
	//                LHS   RHS
	neg, pos := nums[:zeroLHS], nums[zeroRHS+1:]
	maxSum := 0
	if len(neg) > k {
		for _, n := range neg {
			if k > 0 {
				maxSum -= n
				k--
			} else {
				maxSum += n
			}
		}
		for _, n := range pos {
			maxSum += n
		}
	} else {
		if zeroLHS != zeroRHS {
			// nums contains zeros
			k = 0
		} else {
			// nums does not contain zero
			k = (k - len(neg)) % 2
		}

		for _, n := range neg {
			maxSum -= n
		}
		for _, n := range pos {
			maxSum += n
		}

		if k == 1 {
			// abs of the max negative number and the min positive number
			var maxNeg, minPos int
			if len(neg) == 0 {
				maxNeg = -1 // no negative number
			} else {
				maxNeg = -neg[len(neg)-1]
			}

			if len(pos) == 0 {
				minPos = -1 // no positive number
			} else {
				minPos = pos[0]
			}

			// case 1 len(pos) > 0 && len(neg) > 0
			if minPos > 0 && maxNeg > 0 {
				if minPos > maxNeg {
					maxSum -= 2 * maxNeg
				} else {
					maxSum -= 2 * minPos
				}
			}
			// case 2 len(pos) > 0 && len(neg) = 0
			if minPos > 0 && maxNeg < 0 {
				maxSum -= 2 * minPos
			}
			// case 3 len(pos) = 0 && len(neg) > 0
			if minPos < 0 && maxNeg > 0 {
				maxSum -= 2 * maxNeg
			}
		}
	}
	return maxSum
}

func (q Q1005) SearchLeftMost(nums []int, target int) int {
	lhs, rhs, mid := 0, len(nums)-1, 0
	for lhs != rhs {
		mid = (lhs + rhs) / 2
		if nums[mid] < target {
			lhs = mid + 1
		} else {
			rhs = mid
		}
	}

	if nums[lhs] >= target {
		return lhs
	}

	return -1
}

func (q Q1005) SearchRightMost(nums []int, target int) int {
	lhs, rhs, mid := 0, len(nums)-1, 0
	for lhs != rhs {
		mid = (lhs + rhs + 1) / 2
		// fmt.Printf("LHS: %d, RHS: %d, MID: %d (%d)\n", lhs, rhs, mid, nums[mid])
		if nums[mid] > target {
			rhs = mid - 1
		} else {
			lhs = mid
		}
	}

	if nums[rhs] <= target {
		return rhs
	}

	return -1
}
