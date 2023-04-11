package leetcode

import (
	"fmt"
	"strings"
)

type Q1254 struct{}

type Q1254CC struct {
	node []int
}

func (q Q1254) NewCC(n int) *Q1254CC {
	cc := &Q1254CC{
		node: make([]int, n),
	}

	for i := 0; i < n; i++ {
		cc.node[i] = i
	}
	return cc
}

func (cc *Q1254CC) Root(i int) (int, int) {
	nhop := 0
	for cc.node[i] != i {
		i = cc.node[i]
		nhop++
	}
	return i, nhop
}

func (cc *Q1254CC) Connect(i, j int) {
	i, ihop := cc.Root(i)
	j, jhop := cc.Root(j)

	if i != j {
		if ihop > jhop {
			cc.node[j] = i
		} else {
			cc.node[i] = j
		}
	}
}

func (cc *Q1254CC) NSubGraph() int {
	n := 0
	for i := 0; i < len(cc.node); i++ {
		if i == cc.node[i] {
			n++
		}
	}
	return n
}

type Q1254Grid struct {
	grid          [][]int
	nRow, nCol, n int
	water, mland  int
	Q1254CC
}

func (q Q1254) NewGrid(grid [][]int) Q1254Grid {
	if len(grid) == 0 {
		return Q1254Grid{}
	}

	g := Q1254Grid{
		grid:  grid,
		nRow:  len(grid),
		nCol:  len(grid[0]),
		water: -1,
		mland: -1,
	}
	g.n = g.nCol * g.nRow
	g.Q1254CC = *q.NewCC(g.n)
	return g
}

func (g Q1254Grid) PositionToIndex(row, col int) int {
	return row*g.nCol + col
}

func (g Q1254Grid) Margin(which string) []int {
	switch which {
	case "row":
		return []int{0, g.nRow - 1}
	case "col":
		return []int{0, g.nCol - 1}
	default:
		panic("which should be either row or col")
	}
}

func (g Q1254Grid) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Grid (water %d margin land %d):", g.water, g.mland))
	for i := 0; i < g.n; i++ {
		if i%g.nCol == 0 {
			sb.WriteRune('\n')
		}
		j, _ := g.Root(g.Q1254CC.node[i])
		switch j {
		case g.mland:
			sb.WriteString("  -")
		case g.water:
			sb.WriteString("  .")
		default:
			sb.WriteString(fmt.Sprintf("%3d", j))
		}
	}
	sb.WriteRune('\n')
	return sb.String()
}

func (q Q1254) ClosedIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	g := q.NewGrid(grid)

	for row := 0; row < g.nRow; row++ {
		for _, col := range g.Margin("col") {
			if grid[row][col] == 0 {
				j := g.PositionToIndex(row, col)
				if g.mland < 0 {
					g.mland = j
				} else {
					g.node[j] = g.mland
				}
			}
		}
	}

	for col := 0; col < g.nCol; col++ {
		for _, row := range g.Margin("row") {
			if grid[row][col] == 0 {
				j := g.PositionToIndex(row, col)
				if g.mland < 0 {
					g.mland = j
				} else {
					g.node[j] = g.mland
				}
			}
		}
	}

	for row := 0; row < g.nRow; row++ {
		for col := 0; col < g.nCol; col++ {
			if grid[row][col] == 0 {
				if row+1 < g.nRow && grid[row][col] == grid[row+1][col] {
					g.Connect(g.PositionToIndex(row, col), g.PositionToIndex(row+1, col))
				}
				if col+1 < g.nCol && grid[row][col] == grid[row][col+1] {
					g.Connect(g.PositionToIndex(row, col), g.PositionToIndex(row, col+1))
				}
			} else {
				j := g.PositionToIndex(row, col)
				if g.water < 0 {
					g.water = j
				} else {
					g.node[g.PositionToIndex(row, col)] = g.water
				}
			}
		}
	}

	// fmt.Println(g)
	nCC := g.NSubGraph()
	if g.water >= 0 {
		nCC--
	}

	if g.mland >= 0 {
		nCC--
	}

	return nCC
}
