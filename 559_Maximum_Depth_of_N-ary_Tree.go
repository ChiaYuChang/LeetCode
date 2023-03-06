package leetcode

type q559TreeNode struct {
	Val      int
	Children []*q559TreeNode
}

func MaxDepthForNaryTree(root *q559TreeNode) int {
	if root == nil {
		return 0
	}

	maxLen := 0
	for _, c := range root.Children {
		l := MaxDepthForNaryTree(c)
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen + 1
}
