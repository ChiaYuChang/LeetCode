package leetcode

type Q0662 struct{}

type Q0662TreeNode struct {
	Val   int
	Left  *Q0662TreeNode
	Right *Q0662TreeNode
}

type Q0662IndexTreeNode struct {
	*Q0662TreeNode
	index int
}

func (q Q0662) AsIndexTreeNode(node *Q0662TreeNode, i int) *Q0662IndexTreeNode {
	return &Q0662IndexTreeNode{node, i}
}

func (q Q0662) WidthOfBinaryTree(root *Q0662TreeNode) int {
	prev := []*Q0662IndexTreeNode{q.AsIndexTreeNode(root, 0)}

	width := 0
	for len(prev) > 0 {
		if w := prev[len(prev)-1].index - prev[0].index + 1; w > width {
			width = w
		}

		curr := make([]*Q0662IndexTreeNode, 0, len(prev)*2)
		for _, n := range prev {
			if n.Left != nil {
				curr = append(curr, q.AsIndexTreeNode(n.Left, n.index*2))
			}

			if n.Right != nil {
				curr = append(curr, q.AsIndexTreeNode(n.Right, n.index*2+1))
			}
		}

		if len(curr) > 0 {
			shft := curr[0].index
			for i := range curr {
				curr[i].index -= shft
			}
		}

		prev = curr
	}
	return width
}
