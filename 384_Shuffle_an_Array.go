package leetcode

import (
	"math/rand"
)

type Solution struct {
	ori    []int
	length int
}

func SolutionConstructor(nums []int) Solution {
	return Solution{ori: nums, length: len(nums)}
}

func (this *Solution) Reset() []int {
	return this.ori
}

func (this *Solution) Shuffle() []int {
	shfflnms := make([]int, this.length)
	copy(shfflnms, this.ori)
	for i := range shfflnms {
		j := rand.Int() % this.length
		shfflnms[i], shfflnms[j] = shfflnms[j], shfflnms[i]
	}
	return shfflnms
}
