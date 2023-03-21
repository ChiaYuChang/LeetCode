package leetcode

type Q0104 struct{}

type Q0104TreeNode struct {
	Val   int
	Left  *Q0104TreeNode
	Right *Q0104TreeNode
}

func (q Q0104) MaxDepthForBinaryTree(root *Q0104TreeNode) int {
	if root == nil {
		return 0
	}
	l := q.MaxDepthForBinaryTree(root.Left)
	r := q.MaxDepthForBinaryTree(root.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}
