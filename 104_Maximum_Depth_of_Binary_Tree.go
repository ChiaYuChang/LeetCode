package leetcode

type q104TreeNode struct {
	Val   int
	Left  *q104TreeNode
	Right *q104TreeNode
}

func MaxDepthForBinaryTree(root *q104TreeNode) int {
	if root == nil {
		return 0
	}
	l := MaxDepthForBinaryTree(root.Left)
	r := MaxDepthForBinaryTree(root.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}
