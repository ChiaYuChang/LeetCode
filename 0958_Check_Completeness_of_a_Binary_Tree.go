package leetcode

type Q0958 struct{}

type Q0958TreeNode struct {
	Val   int
	Left  *Q0958TreeNode
	Right *Q0958TreeNode
}

func (q Q0958) IsCompleteTree(root *Q0958TreeNode) bool {
	if root == nil {
		return true
	}

	// init two slices as queues
	var q1, q2 []*Q0958TreeNode
	q1 = append(q1, root)

	for len(q1) > 0 {
		for _, n := range q1 {
			q2 = append(q2, n.Left)
			q2 = append(q2, n.Right)
		}

		for i, n := range q2 {
			if n == nil {
				return q.WithNoChild(q2[:i]) && q.IsAllNil(q2[i:])
			}
		}
		q1, q2 = q2, []*Q0958TreeNode{}
	}

	// should never reach here
	panic("unknown error")
}

func (q Q0958) IsAllNil(nodes []*Q0958TreeNode) bool {
	for _, n := range nodes {
		if n != nil {
			return false
		}
	}
	return true
}

func (q Q0958) WithNoChild(nodes []*Q0958TreeNode) bool {
	for _, n := range nodes {
		if n.Left != nil || n.Right != nil {
			return false
		}
	}
	return true
}
