package leetcode

type Q0559 struct{}

type Q0559TreeNode struct {
	Val      int
	Children []*Q0559TreeNode
}

func (q Q0559) MaxDepthForNaryTree(root *Q0559TreeNode) int {
	if root == nil {
		return 0
	}

	maxLen := 0
	for _, c := range root.Children {
		l := q.MaxDepthForNaryTree(c)
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen + 1
}
