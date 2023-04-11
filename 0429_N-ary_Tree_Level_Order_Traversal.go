package leetcode

type Q0429 struct{}

type Node struct {
	Val      int
	Children []*Node
}

func (q Q0429) LevelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	var curr, next []*Node
	curr = append(curr, root)

	level := [][]int{}
	for len(curr) > 0 {
		l := []int{}
		for i := range curr {
			l = append(l, curr[i].Val)
			for j := range curr[i].Children {
				next = append(next, curr[i].Children[j])
			}
		}
		level = append(level, l)
		curr, next = next, []*Node{}
	}
	return level
}
