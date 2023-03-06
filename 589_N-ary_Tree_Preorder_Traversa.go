package leetcode

type q589NTreeNode struct {
	Val      int
	Children []*q589NTreeNode
}

func NTreePreOrder(root *q589NTreeNode) []int {
	po := make([]int, 0, 10)
	NTreeDFS(root, &po)
	return po
}

func NTreeDFS(root *q589NTreeNode, po *[]int) {
	if root == nil {
		return
	}

	(*po) = append((*po), root.Val)
	for i := range root.Children {
		NTreeDFS(root.Children[i], po)
	}
}
