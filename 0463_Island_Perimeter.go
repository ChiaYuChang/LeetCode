package leetcode

type Q0463 struct{}

type Q0463Grid struct {
	grid       [][]int
	nrow, ncol int
}

func (q Q0463) NewGrid(grid [][]int) *Q0463Grid {
	g := &Q0463Grid{grid: grid}
	g.nrow = len(grid)
	if g.nrow < 1 {
		return nil
	}
	g.ncol = len(grid[0])
	return g
}

func (g Q0463Grid) IsLand(row, col int) bool {
	return !g.IsWalter(row, col)
}

func (g Q0463Grid) IsWalter(row, col int) bool {
	if row < 0 || row >= g.nrow || col < 0 || col >= g.ncol {
		return true
	}
	return g.grid[row][col] == 0
}

func (g Q0463Grid) NAdjacentCellIsWater(row, col int) int {
	nWater := 0
	if g.IsWalter(row-1, col) {
		nWater++
	}
	if g.IsWalter(row+1, col) {
		nWater++
	}
	if g.IsWalter(row, col-1) {
		nWater++
	}
	if g.IsWalter(row, col+1) {
		nWater++
	}
	return nWater
}

func (q Q0463) IslandPerimeter(grid [][]int) int {
	g := q.NewGrid(grid)
	perimeter := 0

	for row := 0; row < g.nrow; row++ {
		for col := 0; col < g.ncol; col++ {
			if g.IsLand(row, col) {
				perimeter += g.NAdjacentCellIsWater(row, col)
			}
		}
	}
	return perimeter
}
