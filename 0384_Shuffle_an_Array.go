package leetcode

import (
	"math/rand"
)

type Q0384 struct{}

type Q0384Solution struct {
	ori    []int
	length int
}

func (q Q0384) NewSolution(nums []int) Q0384Solution {
	return Q0384Solution{ori: nums, length: len(nums)}
}

func (this *Q0384Solution) Reset() []int {
	return this.ori
}

func (this *Q0384Solution) Shuffle() []int {
	shfflnms := make([]int, this.length)
	copy(shfflnms, this.ori)
	for i := range shfflnms {
		j := rand.Int() % this.length
		shfflnms[i], shfflnms[j] = shfflnms[j], shfflnms[i]
	}
	return shfflnms
}
