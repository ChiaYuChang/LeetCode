package leetcode

import (
	"fmt"
)

type Q1020 struct{}

type Grid struct {
	nRow, nCol int
	grid       [][]int
}

type Position [2]int

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p[0], p[1])
}

func NewGrid(grid [][]int) Grid {
	g := Grid{grid: grid}
	if len(grid) == 0 {
		return g
	}

	g.nRow = len(g.grid)
	g.nCol = len(g.grid[0])
	return g
}

func (g Grid) IsOutOfRange(p Position) bool {
	return p[0] < 0 || p[0] >= g.nRow || p[1] < 0 || p[1] >= g.nCol
}

func (g Grid) Neighbors(p Position) []Position {
	if g.IsOutOfRange(p) {
		return []Position{}
	}

	ns := make([]Position, 0, 4)
	if p[0] > 0 {
		ns = append(ns, Position{p[0] - 1, p[1]})
	}

	if p[1] > 0 {
		ns = append(ns, Position{p[0], p[1] - 1})
	}

	if p[0] < g.nRow-1 {
		ns = append(ns, Position{p[0] + 1, p[1]})
	}

	if p[1] < g.nCol-1 {
		ns = append(ns, Position{p[0], p[1] + 1})
	}

	return ns
}

func (g Grid) IsLand(p Position) bool {
	if g.IsOutOfRange(p) {
		return false
	}
	return g.grid[p[0]][p[1]] == 1
}

func NumEnclaves(grid [][]int) int {
	g := NewGrid(grid)
	hasVisited := make([][]bool, g.nRow)
	for i := 0; i < g.nRow; i++ {
		hasVisited[i] = make([]bool, g.nCol)
	}

	// add margin to queue
	q1, q2 := map[Position]struct{}{}, map[Position]struct{}{}
	for row := 0; row < g.nRow; row++ {
		pfc, plc := Position{row, 0}, Position{row, g.nCol - 1}
		if g.IsLand(pfc) {
			q1[pfc] = struct{}{}
		}
		if g.IsLand(plc) {
			q1[plc] = struct{}{}
		}
	}

	for col := 0; col < g.nCol; col++ {
		pfr, plr := Position{0, col}, Position{g.nRow - 1, col}
		if g.IsLand(pfr) {
			q1[pfr] = struct{}{}
		}
		if g.IsLand(plr) {
			q1[plr] = struct{}{}
		}
	}

	// DFS
	for len(q1) > 0 {
		for p := range q1 {
			hasVisited[p[0]][p[1]] = true
			for _, n := range g.Neighbors(p) {
				if g.IsLand(n) && !hasVisited[n[0]][n[1]] {
					q2[n] = struct{}{}
				}
			}
		}
		q1, q2 = q2, map[Position]struct{}{}
	}

	nEnclaves := 0
	for row := 0; row < g.nRow; row++ {
		for col := 0; col < g.nCol; col++ {
			if g.IsLand(Position{row, col}) && !hasVisited[row][col] {
				nEnclaves++
			}
		}
	}
	return nEnclaves
}
