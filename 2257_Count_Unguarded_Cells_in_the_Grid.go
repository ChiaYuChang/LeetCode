package leetcode

type Q2257 struct{}

type Q2257Grid struct {
	NRow, NCol int
	IsGuarded  [][]bool
	IsOccupied [][]bool
}

var Q2257directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func (q Q2257) NewGrid(nrow, ncol int) Q2257Grid {
	grid := Q2257Grid{
		NRow:       nrow,
		NCol:       ncol,
		IsGuarded:  make([][]bool, nrow),
		IsOccupied: make([][]bool, nrow),
	}

	for i := 0; i < nrow; i++ {
		grid.IsGuarded[i] = make([]bool, ncol)
		grid.IsOccupied[i] = make([]bool, ncol)
	}
	return grid
}

func (g Q2257Grid) IsNotOutOfRange(row, col int) bool {
	return row >= 0 && col >= 0 && row < g.NRow && col < g.NCol
}

func (g Q2257Grid) CountUnguarded() int {
	nUnguarded := 0
	for row := 0; row < g.NRow; row++ {
		for col := 0; col < g.NCol; col++ {
			if !(g.IsGuarded[row][col] || g.IsOccupied[row][col]) {
				nUnguarded++
			}
		}
	}
	return nUnguarded
}

func (q Q2257) CountUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := q.NewGrid(m, n)
	for _, wall := range walls {
		grid.IsOccupied[wall[0]][wall[1]] = true
	}

	for _, guard := range guards {
		grid.IsOccupied[guard[0]][guard[1]] = true
	}

	for _, guard := range guards {
		grid.IsGuarded[guard[0]][guard[1]] = true
		for _, dir := range Q2257directions {
			row, col := guard[0]+dir[0], guard[1]+dir[1]
			for grid.IsNotOutOfRange(row, col) {
				if grid.IsOccupied[row][col] {
					break
				}
				grid.IsGuarded[row][col] = true
				row += dir[0]
				col += dir[1]
			}
		}
	}

	return grid.CountUnguarded()
}
