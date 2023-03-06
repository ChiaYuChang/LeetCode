package leetcode

type q783TreeNode struct {
	Val   int
	Left  *q783TreeNode
	Right *q783TreeNode
}

func MinDiffInBST(root *q783TreeNode) int {
	if root == nil {
		return 0
	}
	minDiff := -1
	FindLMinRMaxInBinaryTree(root, &minDiff)

	return minDiff
}

func FindLMinRMaxInBinaryTree(root *q783TreeNode, minDiff *int) (int, int) {
	if root == nil {
		return -1, -1 // all valid node value should be greather or equal to 0
	}

	lmin, lmax := FindLMinRMaxInBinaryTree(root.Left, minDiff)
	if lmax >= 0 && ((*minDiff) < 0 || root.Val-lmax < (*minDiff)) {
		(*minDiff) = root.Val - lmax
	}

	rmin, rmax := FindLMinRMaxInBinaryTree(root.Right, minDiff)
	if rmin >= 0 && ((*minDiff) < 0 || rmin-root.Val < (*minDiff)) {
		(*minDiff) = rmin - root.Val
	}

	if lmin < 0 {
		lmin = root.Val
	}

	if rmax < 0 {
		rmax = root.Val
	}

	return lmin, rmax
}
