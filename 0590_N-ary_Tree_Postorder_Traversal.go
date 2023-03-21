package leetcode

type Q0590 struct{}

type Q0590NTreeNode struct {
	Val      int
	Children []*Q0590NTreeNode
}

func (q Q0590) NTreePostOrder(root *Q0590NTreeNode) []int {
	po := make([]int, 0, 1000)
	q.NTreePostOrderTravel(root, &po)
	return po
}

func (q Q0590) NTreePostOrderTravel(root *Q0590NTreeNode, po *[]int) {
	if root == nil {
		return
	}

	for i := range root.Children {
		q.NTreePostOrderTravel(root.Children[i], po)
	}
	(*po) = append((*po), root.Val)
}
