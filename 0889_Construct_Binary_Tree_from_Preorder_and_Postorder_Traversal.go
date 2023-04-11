package leetcode

type Q0889 struct{}

type Q0889TreeNode struct {
	Val   int
	Left  *Q0889TreeNode
	Right *Q0889TreeNode
}

func (q Q0889) PreOrder(n *Q0889TreeNode) []int {
	preorder := []int{}
	q.preOrder(n, &preorder)
	return preorder
}

func (q Q0889) preOrder(n *Q0889TreeNode, order *[]int) {
	if n == nil {
		return
	}
	*order = append(*order, n.Val)
	q.preOrder(n.Left, order)
	q.preOrder(n.Right, order)
}

func (q Q0889) PostOrder(n *Q0889TreeNode) []int {
	postorder := []int{}
	q.postOrder(n, &postorder)
	return postorder
}

func (q Q0889) postOrder(n *Q0889TreeNode, order *[]int) {
	if n == nil {
		return
	}
	q.postOrder(n.Left, order)
	q.postOrder(n.Right, order)
	*order = append(*order, n.Val)
}

func (q Q0889) ConstructFromPrePost(preorder []int, postorder []int) *Q0889TreeNode {
	// invalid input
	if len(preorder) != len(postorder) {
		return nil
	}

	l := len(preorder)
	if l == 0 {
		return nil
	}

	if preorder[0] != postorder[l-1] {
		return nil
	}
	node := &Q0889TreeNode{Val: preorder[0]}

	if l == 2 {
		// asign to node.Right is also ok
		node.Left = q.ConstructFromPrePost(preorder[1:], postorder[:1])
	}

	if l > 2 {
		// at least three nodes
		lVal := preorder[1]
		rVal := postorder[l-2]

		var preorderLHS, preorderRHS, postorderLHS, postorderRHS []int
		if lVal == rVal {
			preorderLHS = preorder[1:]
			postorderLHS = postorder[:(l - 1)]
		} else {
			i, j := 1, l-2
			for ; i < l-1; i++ {
				if preorder[i] == rVal {
					break
				}
			}

			for ; j > 1; j-- {
				if postorder[j-1] == lVal {
					break
				}
			}

			preorderLHS = preorder[1:i]
			preorderRHS = preorder[i:]
			postorderLHS = postorder[:j]
			postorderRHS = postorder[j:(l - 1)]
		}
		node.Left = q.ConstructFromPrePost(preorderLHS, postorderLHS)
		node.Right = q.ConstructFromPrePost(preorderRHS, postorderRHS)
	}
	return node
}
