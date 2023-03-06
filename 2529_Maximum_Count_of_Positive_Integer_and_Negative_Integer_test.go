package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindZeroBoundary(t *testing.T) {
	type testCase struct {
		name string
		nums []int
		ans  int
	}

	tcs := []testCase{
		{
			name: "w/ zeros",
			nums: []int{-3, -2, -1, -1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 3},
			ans:  6,
		},
		{
			name: "w/o zeros more positive numbers",
			nums: []int{-3, -2, -1, -1, 1, 1, 1, 2, 3},
			ans:  5,
		},
		{
			name: "w/o zeros more negative numbers",
			nums: []int{-3, -2, -1, -1, 1, 1, 2},
			ans:  4,
		},
		{
			name: "start with zeros",
			nums: []int{0, 0, 5, 20, 66, 1314},
			ans:  4,
		},
		{
			name: "end with zeros",
			nums: []int{-1314, -66, -20, -5, 0, 0, 0, 0},
			ans:  4,
		},
		{
			name: "Case 1",
			nums: []int{-2, -1, -1, 1, 2, 3},
			ans:  3,
		},
		{
			name: "Case 2",
			nums: []int{-3, -2, -1, 0, 0, 1, 2},
			ans:  3,
		},
		{
			name: "Case 3",
			nums: []int{5, 20, 66, 1314},
			ans:  4,
		},
		{
			name: "All zero",
			nums: []int{0, 0, 0, 0, 0},
			ans:  0,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(t,
					MaxiMumCount(tc.nums),
					tc.ans,
				)
			},
		)
	}
}

func MaxiMumCount(nums []int) int {
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return len(nums)
	}

	if nums[0] == 0 && nums[len(nums)-1] == 0 {
		return 0
	}

	lhs, rhs, mid := 0, len(nums), 0
	for rhs-lhs > 1 {
		mid = (lhs + rhs) / 2
		if nums[mid] < 0 {
			lhs = mid
		} else if nums[mid] > 0 {
			rhs = mid
		} else if nums[mid] == 0 {
			rhs = FindZerosBoundary(mid, rhs, nums, true)
			lhs = FindZerosBoundary(lhs, mid, nums, false)
			break
		}
	}
	nneg := lhs + 1
	npos := len(nums) - rhs

	if nneg > npos {
		return nneg
	}
	return npos
}

func FindZerosBoundary(lb, ub int, nums []int, direction bool) int {
	lhs, rhs, mid := lb, ub, 0

	for rhs-lhs > 1 {
		mid = (lhs + rhs) / 2
		if nums[mid] == 0 {
			if direction {
				lhs = mid
			} else {
				rhs = mid
			}
		} else {
			if direction {
				rhs = mid
			} else {
				lhs = mid
			}
		}
	}

	if direction {
		return rhs
	}
	return lhs

}
