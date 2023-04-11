package leetcode

type Q0110 struct{}

type Q0110TreeNode struct {
	Val   int
	Left  *Q0110TreeNode
	Right *Q0110TreeNode
}

func (q Q0110) IsBalanced(root *Q0110TreeNode) bool {
	return q.helper(root) != -1
}

func (q Q0110) helper(root *Q0110TreeNode) int {
	if root == nil {
		return 0
	}

	lhs := q.helper(root.Left)
	if lhs == -1 {
		return -1
	}

	rhs := q.helper(root.Right)
	if rhs == -1 {
		return -1
	}

	if rhs > lhs {
		rhs, lhs = lhs, rhs
	}

	if lhs-rhs > 1 {
		return -1
	}
	return lhs + 1
}
