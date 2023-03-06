package leetcode

type q543TreeNode struct {
	Val   int
	Left  *q543TreeNode
	Right *q543TreeNode
}

func diameterOfBinaryTree(root *q543TreeNode) int {
	if root == nil {
		return 0
	}

	var diameter int
	travel(root, 0, &diameter)

	return diameter
}

func travel(n *q543TreeNode, depth int, diameter *int) int {
	if n == nil {
		return depth
	}

	depth++

	ll := travel(n.Left, depth, diameter)
	rl := travel(n.Right, depth, diameter)

	if ll+rl-2*depth > *diameter {
		*diameter = ll + rl - 2*depth
	}

	if ll > rl {
		return ll
	}
	return rl
}
