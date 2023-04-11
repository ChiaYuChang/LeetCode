package leetcode

import "fmt"

type Q0785 struct{}

type Q0785TNodeColor int8

const (
	Q0785TColUnKnown Q0785TNodeColor = iota
	Q0785TColRed
	Q0785TColBlack
)

func (c Q0785TNodeColor) flip() Q0785TNodeColor {
	if c == Q0785TColBlack {
		return Q0785TColRed
	} else if c == Q0785TColRed {
		return Q0785TColBlack
	} else {
		return Q0785TColUnKnown
	}
}

func (c Q0785TNodeColor) String() string {
	if c == Q0785TColBlack {
		return "B"
	} else if c == Q0785TColRed {
		return "R"
	} else {
		return "?"
	}
}

type Q0785UndirectedGraph struct {
	NodeColor []Q0785TNodeColor
	Edge      *[][]int
}

func (q Q0785) NewUndirectedGraph(edges [][]int) Q0785UndirectedGraph {
	ug := Q0785UndirectedGraph{
		NodeColor: make([]Q0785TNodeColor, len(edges)),
		Edge:      &edges,
	}

	for i := 0; i < len(edges); i++ {
		for _, j := range edges[i] {
			(*ug.Edge)[j] = append((*ug.Edge)[j], i)
		}
	}

	fmt.Println("New Undirected Graph")
	for i := 0; i < len((*ug.Edge)); i++ {
		fmt.Printf("%d: %v\n", i, (*ug.Edge)[i])
	}

	return ug
}

func (g Q0785UndirectedGraph) Color(i int) Q0785TNodeColor {
	return g.NodeColor[i]
}

func (g *Q0785UndirectedGraph) Paint(i int, col Q0785TNodeColor) {
	g.NodeColor[i] = col
}

func (g Q0785UndirectedGraph) HasVisited(i int) bool {
	return g.Color(i) != Q0785TColUnKnown
}

func (g Q0785UndirectedGraph) Neighbors(i int) []int {
	return (*g.Edge)[i]
}

func (g Q0785UndirectedGraph) GetFirstUnvisitedNode() int {
	for i := range g.NodeColor {
		if g.NodeColor[i] == Q0785TColUnKnown {
			return i
		}
	}
	return -1
}

func (g Q0785UndirectedGraph) IsBipartite() bool {
	for seed := g.GetFirstUnvisitedNode(); seed >= 0; seed = g.GetFirstUnvisitedNode() {
		q1, q2 := []int{seed}, []int{}
		currColor := Q0785TColBlack
		for len(q1) > 0 {
			fmt.Println(g.NodeColor)
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

func (q Q0785) IsBipartite(graph [][]int) bool {
	ug := q.NewUndirectedGraph(graph)
	return ug.IsBipartite()
}
