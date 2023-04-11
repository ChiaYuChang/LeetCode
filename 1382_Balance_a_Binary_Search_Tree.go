package leetcode

type Q1382 struct{}

type Q1382TreeNode struct {
	Val   int
	Left  *Q1382TreeNode
	Right *Q1382TreeNode
}

func (q Q1382) BalanceBST(root *Q1382TreeNode) *Q1382TreeNode {
	nodes := []*Q1382TreeNode{}
	q.inOrder(root, &nodes) // sorted nodes

	// clean up
	for i := range nodes {
		nodes[i].Left, nodes[i].Right = nil, nil
	}

	// rebuild the tree
	return q.buildBST(nodes, 0, len(nodes)-1)
}

func (q Q1382) buildBST(nodes []*Q1382TreeNode, lhs, rhs int) *Q1382TreeNode {
	if lhs > rhs {
		return nil
	}
	mid := (lhs + rhs) / 2
	node := nodes[mid]
	node.Left = q.buildBST(nodes, lhs, mid-1)
	node.Right = q.buildBST(nodes, mid+1, rhs)
	return node
}

func (q Q1382) inOrder(node *Q1382TreeNode, list *[]*Q1382TreeNode) {
	if node == nil {
		return
	}
	q.inOrder(node.Left, list)
	(*list) = append((*list), node)
	q.inOrder(node.Right, list)
}

// func balancing(pNode, cNode *TreeNode) int {
// 	if cNode == nil {
// 		return 0
// 	}
// 	dLhs := balancing(cNode, cNode.Left)
// 	dRhs := balancing(cNode, cNode.Right)

// 	if AbsDiff(dLhs, dRhs) <= 1 {
// 		// already balance
// 		return Max(dLhs, dRhs) + 1
// 	}

// 	if dRhs > dLhs {
// 		pNode.Right, cNode.Right = cNode.Right, nil
// 		Swing(pNode.Right, cNode)
// 	} else {
// 		pNode.Left, cNode.Left = cNode.Left, nil
// 		Swing(pNode.Left, cNode)
// 	}

// }

// func Swing(root, node *TreeNode) {
// 	if node.Val == root.Val {
// 		panic("duplicated values")
// 	}
// 	if node.Val > root.Val {
// 		if root.Right != nil {
// 			Swing(root.Right, node)
// 		} else {
// 			root.Right = node
// 		}
// 	} else {
// 		if root.Left != nil {
// 			Swing(root.Left, node)
// 		} else {
// 			root.Left = node
// 		}
// 	}
// }

// func Max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func AbsDiff(x, y int) int {
// 	if x > y {
// 		return x - y
// 	}
// 	return y - x
// }
