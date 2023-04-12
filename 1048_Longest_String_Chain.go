package leetcode

import (
	"sort"
)

type Q1048 struct{}

func (q Q1048) LongestStrChain(words []string) int {
	IsSingleInsertion := func(w1, w2 string) bool {
		// whether is single insertion/deletion between w1 and w2
		// len(w1) - len(w2) should equal 1
		i := 0
		for ; i < len(w1); i++ {
			if w2[i] != w1[i] {
				break
			}
		}
		return w2[i+1:] == w1[i:]
	}

	Max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	strLenMap := make(map[int][]int)
	for i, w := range words {
		strLenMap[len(w)] = append(strLenMap[len(w)], i)
	}

	keys := make([]int, 0, len(strLenMap))
	for k := range strLenMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.IntSlice(keys))

	// DP
	prev := keys[0]
	node := make([]int, len(words))
	maxHop := 0
	for i := 1; i < len(keys); i++ {
		if keys[i]-prev > 1 {
			prev = keys[i]
			continue
		}
		for _, w1i := range strLenMap[prev] {
			for _, w2i := range strLenMap[keys[i]] {
				if IsSingleInsertion(words[w1i], words[w2i]) {
					node[w2i] = Max(node[w2i], node[w1i]+1)
					maxHop = Max(node[w2i], maxHop)
				}
			}
		}
		prev = keys[i]
	}

	// length = hops + 1
	return maxHop + 1
}
