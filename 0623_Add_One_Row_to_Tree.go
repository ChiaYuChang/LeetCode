package leetcode

type Q0623 struct{}

type Q0623TreeNode struct {
	Val   int
	Left  *Q0623TreeNode
	Right *Q0623TreeNode
}

func (q Q0623) AddOneRow(root *Q0623TreeNode, val int, depth int) *Q0623TreeNode {
	if root == nil {
		return nil
	}

	switch depth {
	case 1:
		// special case
		newRoot := &Q0623TreeNode{Val: val, Left: root}
		return newRoot
	case 2:
		root.Left = &Q0623TreeNode{Val: val, Left: root.Left}
		root.Right = &Q0623TreeNode{Val: val, Right: root.Right}
	default:
		_ = q.AddOneRow(root.Left, val, depth-1)
		_ = q.AddOneRow(root.Right, val, depth-1)
	}
	return root
}
