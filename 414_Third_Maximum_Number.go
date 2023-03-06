package leetcode

import (
	"fmt"
	"math"
)

type Q414 struct{}

type Q414Triplet struct {
	n    int
	nums []int
}

func NewQ414Triplet() *Q414Triplet {
	return &Q414Triplet{n: 0, nums: []int{math.MinInt, math.MinInt, math.MinInt}}
}

func (t Q414Triplet) String() string {
	return fmt.Sprintf("%d: %v", t.n, t.nums)
}

func (t Q414Triplet) IsIn(n int) bool {
	return n == t.nums[0] || n == t.nums[1] || n == t.nums[2]
}

func (t *Q414Triplet) Append(n int) {
	if t.IsIn(n) || (t.n == 3 && n < t.nums[0]) {
		return
	}

	t.nums[0] = n
	for i := 1; i < 3; i++ {
		if t.nums[i-1] > t.nums[i] {
			t.nums[i-1], t.nums[i] = t.nums[i], t.nums[i-1]
		}
	}

	if t.n < 3 {
		t.n++
	}

	return
}

func (t *Q414Triplet) Top3() int {
	if t.n == 3 {
		return t.nums[0]
	}
	return t.nums[2]
}

func (q Q414) ThirdMax(nums []int) int {
	triplet := NewQ414Triplet()
	for _, n := range nums {
		triplet.Append(n)
	}
	return triplet.Top3()
}
