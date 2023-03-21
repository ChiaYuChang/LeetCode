package leetcode

type Q0543 struct{}

type Q0543TreeNode struct {
	Val   int
	Left  *Q0543TreeNode
	Right *Q0543TreeNode
}

func (q Q0543) DiameterOfBinaryTree(root *Q0543TreeNode) int {
	if root == nil {
		return 0
	}

	var diameter int
	q.Travel(root, 0, &diameter)

	return diameter
}

func (q Q0543) Travel(n *Q0543TreeNode, depth int, diameter *int) int {
	if n == nil {
		return depth
	}

	depth++

	ll := q.Travel(n.Left, depth, diameter)
	rl := q.Travel(n.Right, depth, diameter)

	if ll+rl-2*depth > *diameter {
		*diameter = ll + rl - 2*depth
	}

	if ll > rl {
		return ll
	}
	return rl
}
