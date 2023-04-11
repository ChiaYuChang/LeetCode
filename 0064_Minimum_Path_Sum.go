package leetcode

type Q0064 struct{}

func (q Q0064) MinPathSum(grid [][]int) int {
	var nrow, ncol int
	nrow = len(grid)
	if nrow == 0 {
		return 0
	}
	ncol = len(grid[0])

	for i := 1; i < ncol; i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}

	for row := 1; row < nrow; row++ {
		grid[row][0] = grid[row][0] + grid[row-1][0]
		for col := 1; col < ncol; col++ {
			if grid[row][col-1] < grid[row-1][col] {
				grid[row][col] = grid[row][col-1] + grid[row][col]
			} else {
				grid[row][col] = grid[row-1][col] + grid[row][col]
			}

		}
	}
	return grid[nrow-1][ncol-1]
}
