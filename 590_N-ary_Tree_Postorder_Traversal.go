package leetcode

type q590NTreeNode struct {
	Val      int
	Children []*q590NTreeNode
}

func NTreePostOrder(root *q590NTreeNode) []int {
	po := make([]int, 0, 1000)
	NTreePostOrderTravel(root, &po)
	return po
}

func NTreePostOrderTravel(root *q590NTreeNode, po *[]int) {
	if root == nil {
		return
	}

	for i := range root.Children {
		NTreePostOrderTravel(root.Children[i], po)
	}
	(*po) = append((*po), root.Val)
}
