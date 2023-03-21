package leetcode

type Q0236TreeNode struct {
	Val   int
	Left  *Q0236TreeNode
	Right *Q0236TreeNode
}

type Q0236Tree struct {
	node *Q0236TreeNode
}

type Q0236 struct{}

func (this Q0236) LowestCommonAncestor(root, p, q *Q0236TreeNode) *Q0236TreeNode {
	var tree = Q0236Tree{}
	tree.Find(root, p, q)
	return tree.node
}

func (tree *Q0236Tree) Find(root, p, q *Q0236TreeNode) bool {
	if root == nil || tree.node != nil {
		return false
	}

	var l, r, m int
	if tree.Find(root.Left, p, q) {
		l = 1
	}

	if tree.Find(root.Right, p, q) {
		r = 1
	}

	if root.Val == p.Val || root.Val == q.Val {
		m = 1
	}

	if l+r+m >= 2 {
		tree.node = root
	}

	return l+r+m > 0
}
