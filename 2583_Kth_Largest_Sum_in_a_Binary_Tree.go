package leetcode

import "sort"

type Q2583 struct{}

func (q Q2583) KthLargestLevelSum(root *TreeNode, k int) int64 {
	var curr, next []*TreeNode
	curr = append(curr, root)

	levelSums := []int{}
	for len(curr) > 0 {
		levelSum := 0
		for i := range curr {
			levelSum += curr[i].Val
			if curr[i].Left != nil {
				next = append(next, curr[i].Left)
			}
			if curr[i].Right != nil {
				next = append(next, curr[i].Right)
			}
		}
		levelSums = append(levelSums, levelSum)
		curr, next = next, []*TreeNode{}
	}

	if k > len(levelSums) {
		return -1
	}

	sort.Sort(sort.Reverse(sort.IntSlice(levelSums)))
	return int64(levelSums[k-1])
}
