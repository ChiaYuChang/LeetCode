package leetcode

type Q0235 struct{}

type Q0235TreeNode struct {
	Val   int
	Left  *Q0235TreeNode
	Right *Q0235TreeNode
}

func (this Q0235) LowestCommonAncestor(root, p, q *Q0235TreeNode) *Q0235TreeNode {
	return this.FindFstBetween(root, p.Val, q.Val, true, true)
}

func (this Q0235) FindFstBetween(node *Q0235TreeNode, x, y int, lb, ub bool) *Q0235TreeNode {
	if x < y {
		x, y = y, x
	}

	if (x > node.Val && node.Val > y) ||
		(x == node.Val && ub) ||
		(y == node.Val && lb) {
		return node
	}

	// right branch
	if x < node.Val {
		return this.FindFstBetween(node.Left, x, y, lb, ub)
	}
	// left branch
	return this.FindFstBetween(node.Right, x, y, lb, ub)
}
