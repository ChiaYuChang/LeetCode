package leetcode

type Q0589 struct{}

type Q0589NTreeNode struct {
	Val      int
	Children []*Q0589NTreeNode
}

func (q Q0589) NTreePreOrder(root *Q0589NTreeNode) []int {
	po := make([]int, 0, 10)
	q.NTreeDFS(root, &po)
	return po
}

func (q Q0589) NTreeDFS(root *Q0589NTreeNode, po *[]int) {
	if root == nil {
		return
	}

	(*po) = append((*po), root.Val)
	for i := range root.Children {
		q.NTreeDFS(root.Children[i], po)
	}
}
