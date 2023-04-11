package leetcode

import "fmt"

type Q1466 struct{}

func (q Q1466) MinReorder(n int, connections [][]int) int {
	g := NewQ1466Graph(n, connections)
	ans := 0
	g.Travel(0, &ans)
	return ans
}

type Q1466Edge struct {
	nodes   *[]int
	hasUsed bool
}

func (e Q1466Edge) String() string {
	return fmt.Sprintf("%d -> %d (%v)", (*e.nodes)[0], (*e.nodes)[1], e.hasUsed)
}

type Q1466Graph struct {
	edge [][]*Q1466Edge
}

func NewQ1466Graph(n int, edges [][]int) *Q1466Graph {
	g := &Q1466Graph{edge: make([][]*Q1466Edge, n)}

	for i := 0; i < len(edges); i++ {
		e := &Q1466Edge{nodes: &(edges[i]), hasUsed: false}
		g.edge[edges[i][0]] = append(g.edge[edges[i][0]], e)
		g.edge[edges[i][1]] = append(g.edge[edges[i][1]], e)
	}
	return g
}

func (g *Q1466Graph) Travel(i int, nRev *int) {
	for len(g.edge[i]) > 0 {
		e := g.edge[i][0]
		g.edge[i] = g.edge[i][1:]

		if e.hasUsed == true {
			continue
		}
		e.hasUsed = true

		if (*e.nodes)[0] == i {
			(*nRev)++
			g.Travel((*e.nodes)[1], nRev)
		} else {
			g.Travel((*e.nodes)[0], nRev)
		}
	}
	return
}
