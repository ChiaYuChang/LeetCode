package leetcode

type Q0107 struct{}

type Q0107TreeNode struct {
	Val   int
	Left  *Q0107TreeNode
	Right *Q0107TreeNode
}

func (q Q0107) LevelOrderBottom(root *Q0107TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var curr, next []*Q0107TreeNode
	curr = append(curr, root)

	level := [][]int{}
	for len(curr) > 0 {
		l := []int{}
		for i := range curr {
			l = append(l, curr[i].Val)
			if curr[i].Left != nil {
				next = append(next, curr[i].Left)
			}
			if curr[i].Right != nil {
				next = append(next, curr[i].Right)
			}
		}
		level = append(level, l)
		curr, next = next, []*Q0107TreeNode{}
	}

	l := len(level)
	// swap
	for i := 0; i < l/2; i++ {
		level[i], level[l-i-1] = level[l-i-1], level[i]
	}
	return level
}
