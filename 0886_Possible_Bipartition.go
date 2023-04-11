package leetcode

// similar to Q0785
type Q0886 struct{}

func (q Q0886) PossibleBipartition(n int, dislikes [][]int) bool {
	ug := q.NewUndirectedGraph(n, dislikes)
	return ug.IsBipartite()
}

type Q0886TNodeColor int8

const (
	TColUnKnown Q0886TNodeColor = iota
	TColRed
	TColBlack
)

func (c Q0886TNodeColor) flip() Q0886TNodeColor {
	if c == TColBlack {
		return TColRed
	} else if c == TColRed {
		return TColBlack
	} else {
		return TColUnKnown
	}
}

func (c Q0886TNodeColor) String() string {
	if c == TColBlack {
		return "B"
	} else if c == TColRed {
		return "R"
	} else {
		return "?"
	}
}

type Q0886UndirectedGraph struct {
	NodeColor []Q0886TNodeColor
	Edge      [][]int
}

func (q Q0886) NewUndirectedGraph(n int, edges [][]int) Q0886UndirectedGraph {
	ug := Q0886UndirectedGraph{
		NodeColor: make([]Q0886TNodeColor, n+1),
		Edge:      make([][]int, n+1),
	}

	for _, edge := range edges {
		ug.Edge[edge[0]] = append(ug.Edge[edge[0]], edge[1])
		ug.Edge[edge[1]] = append(ug.Edge[edge[1]], edge[0])
	}
	return ug
}

func (g Q0886UndirectedGraph) Color(i int) Q0886TNodeColor {
	return g.NodeColor[i]
}

func (g *Q0886UndirectedGraph) Paint(i int, col Q0886TNodeColor) {
	g.NodeColor[i] = col
}

func (g Q0886UndirectedGraph) HasVisited(i int) bool {
	return g.Color(i) != TColUnKnown
}

func (g Q0886UndirectedGraph) Neighbors(i int) []int {
	return g.Edge[i]
}

func (g Q0886UndirectedGraph) GetFirstUnvisitedNode() int {
	for i := range g.NodeColor {
		if g.NodeColor[i] == TColUnKnown {
			return i
		}
	}
	return -1
}

func (g Q0886UndirectedGraph) IsBipartite() bool {
	for seed := g.GetFirstUnvisitedNode(); seed >= 0; seed = g.GetFirstUnvisitedNode() {
		q1, q2 := []int{seed}, []int{}
		currColor := TColBlack
		for len(q1) > 0 {
			for _, i := range q1 {
				if g.HasVisited(i) {
					if g.Color(i) != currColor {
						return false
					}
				} else {
					g.Paint(i, currColor)
					q2 = append(q2, g.Neighbors(i)...)
				}
			}
			currColor = currColor.flip()
			q1, q2 = q2, []int{}
		}
	}
	return true
}
