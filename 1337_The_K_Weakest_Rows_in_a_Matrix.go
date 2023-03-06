package leetcode

import "sort"

type Q1337 struct{}

type q1337Data struct {
	nSoldier int
	order    int
}

type SortByNumSoldier []q1337Data

func (a SortByNumSoldier) Len() int      { return len(a) }
func (a SortByNumSoldier) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByNumSoldier) Less(i, j int) bool {
	if a[i].nSoldier == a[j].nSoldier {
		return a[i].order < a[j].order
	}
	return a[i].nSoldier < a[j].nSoldier
}

func (q Q1337) KWeakestRows(mat [][]int, k int) []int {
	data := make([]q1337Data, len(mat))
	for i, row := range mat {
		data[i] = q1337Data{nSoldier: q.FirstZeroIndex(row), order: i}
	}
	sort.Sort(SortByNumSoldier(data))

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = data[i].order
	}
	return ans
}

func (q Q1337) FirstZeroIndex(ary []int) int {
	lhs, rhs := 0, len(ary)
	if ary[rhs-1] == 1 {
		return rhs
	}

	if ary[lhs] == 0 {
		return 0
	}

	for rhs-lhs > 16 {
		mid := (rhs + lhs) / 2
		if ary[mid] == 0 {
			rhs = mid
		} else {
			lhs = mid
		}
	}

	for ; lhs < rhs; lhs++ {
		if ary[lhs] == 1 && ary[lhs+1] == 0 {
			break
		}
	}
	return lhs + 1
}
