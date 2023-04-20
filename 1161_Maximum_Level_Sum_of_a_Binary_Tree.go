package leetcode

type Q1161TreeNode struct {
	Val   int
	Left  *Q1161TreeNode
	Right *Q1161TreeNode
}

type Q1161 struct{}

func (q Q1161) MaxLevelSum(root *Q1161TreeNode) int {
	if root == nil {
		return 0
	}

	var curr, next []*Q1161TreeNode
	curr = append(curr, root)

	maxLevelSum := -1 << 32
	maxLevel, curLevel := 0, 1
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
		if levelSum > maxLevelSum {
			maxLevelSum = levelSum
			maxLevel = curLevel
		}
		curLevel++
		curr, next = next, []*Q1161TreeNode{}
	}
	return maxLevel
}
