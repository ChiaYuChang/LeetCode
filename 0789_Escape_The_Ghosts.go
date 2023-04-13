package leetcode

type Q0789 struct{}

func (q Q0789) EscapeGhosts(ghosts [][]int, target []int) bool {
	ManhattanDist := func(x, y []int) int {
		// len(x) should be equal len(y)
		d := 0
		for i := 0; i < len(x); i++ {
			if x[i] > y[i] {
				d += x[i] - y[i]
			} else {
				d += y[i] - x[i]
			}
		}
		return d
	}

	minDist := ManhattanDist([]int{0, 0}, target)
	for _, g := range ghosts {
		if minDist >= ManhattanDist(g, target) {
			return false
		}
	}
	return true
}

// type position [2]int

// func ToPosition(slc []int) position {
// 	return position{slc[0], slc[1]}
// }

// type graph struct {
// 	nRow, nCol int
// 	nodes      [][][]position
// }

// func NewGraph(nRow, nCol int) *graph {
// 	nodes := make([][][]position, nRow)
// 	for i := range nodes {
// 		nodes[i] = make([][]position, nCol)
// 	}
// 	return &graph{nRow: nRow, nCol: nCol, nodes: nodes}
// }

// func (g graph) Connect(p1, p2 position) {
// 	if !g.IsValid(p1) || !g.IsValid(p2) {
// 		return
// 	}
// 	g.nodes[p1[0]][p1[1]] = append(g.nodes[p1[0]][p1[1]], p2)
// 	g.nodes[p2[0]][p2[1]] = append(g.nodes[p2[0]][p2[1]], p1)
// }

// func (g graph) Neighbors(p position) []position {
// 	if g.IsValid(p) {
// 		return g.nodes[p[0]][p[1]]
// 	}
// 	return nil
// }

// func (g graph) IsValid(p position) bool {
// 	if p[0] < 0 || p[0] >= g.nRow {
// 		return false
// 	}

// 	if p[1] < 0 || p[1] >= g.nCol {
// 		return false
// 	}
// 	return true
// }

// func (g graph) BFS(dst position, src ...position) int {
// 	hasVisited := make([][]bool, g.nRow)
// 	for i := range hasVisited {
// 		hasVisited[i] = make([]bool, g.nCol)
// 	}

// 	step := 0
// 	prev := []position{}
// 	for _, s := range src {
// 		for _, ngbr := range g.Neighbors(s) {
// 			prev = append(prev, ngbr)
// 			hasVisited[ngbr[0]][ngbr[1]] = true
// 		}
// 	}

// 	for len(prev) == 0 {
// 		step++
// 		curr := []position{}
// 		for _, p := range prev {
// 			if hasVisited[p[0]][p[1]] {
// 				continue
// 			}
// 			for _, ngbr := range g.Neighbors(p) {
// 				if ngbr == dst {
// 					return step
// 				}
// 				curr = append(curr, ngbr)
// 				hasVisited[ngbr[0]][ngbr[1]] = true
// 			}
// 		}
// 		prev = curr
// 	}
// 	return -1
// }
