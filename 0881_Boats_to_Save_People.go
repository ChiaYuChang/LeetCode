package leetcode

import (
	"sort"
)

type Q0881 struct{}

func (q Q0881) NumRescueBoats(people []int, limit int) int {
	if !sort.IsSorted(sort.IntSlice(people)) {
		sort.Sort(sort.IntSlice(people))
	}
	nBoat := 0

	p1, p2 := len(people)-1, 0
	for p1 >= p2 {
		if people[p1]+people[p2] <= limit {
			p2++
		}
		p1--
		nBoat++
	}

	return nBoat
}
