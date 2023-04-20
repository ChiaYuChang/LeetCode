package leetcode

import "sort"

type Q2099 struct{}

type Q2099MaxK struct {
	num []int
	idx []int
}

func (m Q2099MaxK) FirstK(k int) Q2099MaxK {
	return Q2099MaxK{m.num[:k], m.idx[:k]}
}

type Q2099SortByNum Q2099MaxK

func (a Q2099SortByNum) Len() int { return len(a.num) }
func (a Q2099SortByNum) Swap(i, j int) {
	a.num[i], a.num[j] = a.num[j], a.num[i]
	a.idx[i], a.idx[j] = a.idx[j], a.idx[i]
}

func (a Q2099SortByNum) Less(i, j int) bool {
	return a.num[i] < a.num[j]
}

type Q2099SortByIndex Q2099MaxK

func (a Q2099SortByIndex) Len() int { return len(a.num) }
func (a Q2099SortByIndex) Swap(i, j int) {
	a.num[i], a.num[j] = a.num[j], a.num[i]
	a.idx[i], a.idx[j] = a.idx[j], a.idx[i]
}
func (a Q2099SortByIndex) Less(i, j int) bool {
	return a.idx[i] < a.idx[j]
}

func (q Q2099) MaxSubsequence(nums []int, k int) []int {
	maxK := Q2099MaxK{num: nums, idx: make([]int, len(nums))}
	for i := range nums {
		maxK.idx[i] = i
	}
	sort.Sort(sort.Reverse(Q2099SortByNum(maxK)))
	sort.Sort(Q2099SortByIndex(maxK.FirstK(k)))
	return maxK.num[:k]
}
