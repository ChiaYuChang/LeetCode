package leetcode

type Q0429 struct{}

type Q0429Node struct {
	Val      int
	Children []*Q0429Node
}

func (q Q0429) LevelOrder(root *Q0429Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	var curr, next []*Q0429Node
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
		curr, next = next, []*Q0429Node{}
	}
	return level
}
