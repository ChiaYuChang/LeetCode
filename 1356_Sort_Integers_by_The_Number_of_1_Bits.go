package leetcode

import (
	"sort"
)

type Q1356 struct{}

type SortByNOneBit struct {
	num  []int
	nOne []int
}

func (a SortByNOneBit) Len() int {
	return len(a.num)
}

func (a SortByNOneBit) Swap(i, j int) {
	a.num[i], a.num[j] = a.num[j], a.num[i]
	a.nOne[i], a.nOne[j] = a.nOne[j], a.nOne[i]
}
func (a SortByNOneBit) Less(i, j int) bool {
	if a.nOne[i] != a.nOne[j] {
		return a.nOne[i] < a.nOne[j]
	}
	return a.num[i] < a.num[j]
}

func SortByBits(arr []int) []int {
	// bits.OnesCount(n uint)
	onesCount := func(n int) int {
		nb := 0
		for n > 0 {
			if n&0b1 > 0 {
				nb++
			}
			n >>= 1
		}
		return nb
	}

	n1bit := make([]int, len(arr))
	for i, n := range arr {
		n1bit[i] = onesCount(n)
	}
	sort.Sort(SortByNOneBit{arr, n1bit})
	return arr
}
