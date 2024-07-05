package leetcode

import "sort"

type Q2248 struct{}

const Q2248MaxLen int = 1_000

func (q Q2248) Intersection(nums [][]int) []int {
	numLength := len(nums)
	if numLength == 1 {
		sort.Ints(nums[0])
		return nums[0]
	}

	minLen := Q2248MaxLen
	minPos := -1
	for i, n := range nums {
		if len(n) < minLen {
			minLen = len(n)
			minPos = i
		}
	}

	set := make(map[int]int, minLen)
	for _, n := range nums[minPos] {
		set[n] = 0
	}

	for _, num := range nums {
		for _, n := range num {
			if _, ok := set[n]; ok {
				set[n] += 1
			}
		}
	}

	intrsctn := make([]int, 0, minLen)
	for k, v := range set {
		if v == numLength {
			intrsctn = append(intrsctn, k)
		}
	}

	sort.Ints(intrsctn)
	return intrsctn
}
