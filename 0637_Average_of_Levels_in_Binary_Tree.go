package leetcode

type Q0637 struct{}

type Q0637TreeNode struct {
	Val   int
	Left  *Q0637TreeNode
	Right *Q0637TreeNode
}

func (q Q0637) AverageOfLevels(root *Q0637TreeNode) []float64 {
	var curr, next []*Q0637TreeNode
	var levelAvg []float64

	curr = append(curr, root)
	for len(curr) > 0 {
		sum := 0
		n := 0
		for i := range curr {
			sum, n = sum+curr[i].Val, n+1
			if curr[i].Left != nil {
				next = append(next, curr[i].Left)
			}
			if curr[i].Right != nil {
				next = append(next, curr[i].Right)
			}
		}
		levelAvg = append(levelAvg, q.Division(sum, n))
		curr, next = next, []*Q0637TreeNode{}
	}
	return levelAvg
}

func (q Q0637) Division(n, d int) float64 {
	return float64(n) / float64(d)
}
