package leetcode

import (
	"math/rand"
	"time"
)

type Q0398 struct{}

type Q0398MapSolution struct {
	IndexMap map[int]*[]int
	Rng      *rand.Rand
}

func (q Q0398) MapSolution(nums []int) Q0398MapSolution {
	s := Q0398MapSolution{
		IndexMap: make(map[int]*[]int),
		Rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	for i, n := range nums {
		if _, ok := s.IndexMap[n]; !ok {
			ary := make([]int, 0)
			s.IndexMap[n] = &ary
		}
		*(s.IndexMap[n]) = append(*(s.IndexMap[n]), i)
	}
	return s
}

func (this *Q0398MapSolution) Pick(target int) int {
	a := this.IndexMap[target]
	return (*a)[this.Rng.Intn(len(*a))]
}

type Q0398RSSolution struct {
	nums *[]int
	Rng  *rand.Rand
}

func (q Q0398) ReservoirSamplingSolution(nums []int) Q0398RSSolution {
	return Q0398RSSolution{
		nums: &nums,
		Rng:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Q0398RSSolution) Pick(target int) int {
	n := 1
	rnIdx := -1

	for i := 0; i < len(*this.nums); i++ {
		if (*this.nums)[i] == target {
			if this.Rng.Intn(n) == 0 {
				rnIdx = i
			}
			n++
		}
	}
	return rnIdx
}
