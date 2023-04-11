package leetcode

import "sort"

type Q0455 struct{}

func (q Q0455) FindContentChildren(g []int, s []int) int {
	sort.Sort(sort.IntSlice(g))
	sort.Sort(sort.IntSlice(s))

	var i, j int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}
