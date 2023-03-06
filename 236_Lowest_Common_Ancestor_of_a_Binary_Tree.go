package leetcode

type q236TreeNode struct {
	Val   int
	Left  *q236TreeNode
	Right *q236TreeNode
}

type q236Ans struct {
	node *q236TreeNode
}

func lowestCommonAncestor(root, p, q *q236TreeNode) *q236TreeNode {
	var ans = q236Ans{}
	ans.find(root, p, q)
	return ans.node
}

func (a *q236Ans) find(root, p, q *q236TreeNode) bool {
	if root == nil || a.node != nil {
		return false
	}

	var l, r, m int
	if a.find(root.Left, p, q) {
		l = 1
	}

	if a.find(root.Right, p, q) {
		r = 1
	}

	if root.Val == p.Val || root.Val == q.Val {
		m = 1
	}

	if l+r+m >= 2 {
		a.node = root
	}

	return l+r+m > 0
}
