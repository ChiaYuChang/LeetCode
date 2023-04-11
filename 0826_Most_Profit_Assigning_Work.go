package leetcode

import (
	"fmt"
	"sort"
)

type Q0826 struct{}

type Work [2]int

func (j Work) String() string {
	return fmt.Sprintf("{D: %3d, P: %3d}", j[0], j[1])
}

func MaxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	works := make([]Work, len(difficulty))
	for i, d := range difficulty {
		works[i] = Work{d, profit[i]}
	}
	sort.Slice(works, func(i, j int) bool { return works[i][0] < works[j][0] })

	maxProfEZJob := works[0][1]
	for i := 1; i < len(works); i++ {
		if works[i][1] < maxProfEZJob {
			works[i][1] = maxProfEZJob
		} else {
			maxProfEZJob = works[i][1]
		}
	}

	maxProf := 0
	for _, w := range worker {
		i := sort.Search(len(works), func(i int) bool { return works[i][0] > w }) - 1
		if 0 <= i && i < len(works) {
			maxProf += works[i][1]
		}
	}
	return maxProf
}
