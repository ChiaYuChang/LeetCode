package leetcode

type Q0101TreeNode struct {
	Val   int
	Left  *Q0101TreeNode
	Right *Q0101TreeNode
}

type Q0101 struct{}

func (q Q0101) IsSymmetric(root *Q0101TreeNode) bool {
	if root == nil {
		return true
	}
	return q.SymmetricTravel(root.Left, root.Right)
}

func (q Q0101) SymmetricTravel(lhs, rhs *Q0101TreeNode) bool {
	lhsIsNil, rhsIsNil := lhs == nil, rhs == nil
	if lhsIsNil && rhsIsNil {
		return true
	}

	if (lhsIsNil != rhsIsNil) || (lhs.Val != rhs.Val) {
		return false
	}
	return q.SymmetricTravel(lhs.Left, rhs.Right) && q.SymmetricTravel(lhs.Right, rhs.Left)
}
