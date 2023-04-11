package leetcode

type Q0892 struct{}

func (q Q0892) SurfaceArea(grid [][]int) int {
	c := q.NewCubes(grid)
	return c.CountSurfaceArea()
}

type Q0892Cubes struct {
	surface int
	grid    [][]int
	maxl    int
	l       int
}

func (q Q0892) NewCubes(grid [][]int) *Q0892Cubes {
	c := &Q0892Cubes{grid: grid}
	for row := 0; row < c.NRow(); row++ {
		for col := 0; col < c.NCol(); col++ {
			if c.grid[row][col] > 0 {
				c.surface += 2
			}
			if c.maxl < c.grid[row][col] {
				c.maxl = c.grid[row][col]
			}
		}
	}
	return c
}

func (c *Q0892Cubes) NRow() int {
	return len(c.grid)
}

func (c *Q0892Cubes) NCol() int {
	return len(c.grid[0])
}

func (c *Q0892Cubes) Get(col, row int) int {
	if col < 0 || row < 0 || col >= c.NCol() || row >= c.NRow() {
		return 0
	}
	return c.grid[row][col]
}

func (c *Q0892Cubes) IsZero(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}

func (c *Q0892Cubes) NZeroNeighbors(col, row int) int {
	u := c.IsZero(c.Get(col, row-1))
	d := c.IsZero(c.Get(col, row-1))
	l := c.IsZero(c.Get(col-1, row))
	r := c.IsZero(c.Get(col+1, row))
	return u + d + l + r
}

func (c *Q0892Cubes) HasNext() (ok bool) {
	return c.maxl > c.l
}

func (c *Q0892Cubes) Next() (ok bool) {
	if !c.HasNext() {
		return false
	}

	for row := 0; row < c.NRow(); row++ {
		for col := 0; col < c.NCol(); col++ {
			if c.grid[row][col] > 0 {
				c.grid[row][col]--
			}
		}
	}
	c.l++
	return true
}

func (c *Q0892Cubes) Scan() {
	for row := 0; row < c.NRow(); row++ {
		for col := 0; col < c.NCol(); col++ {
			if c.Get(row, col) > 0 {
				c.surface += c.NZeroNeighbors(row, col)
			}
		}
	}
}

func (c *Q0892Cubes) CountSurfaceArea() int {
	for c.HasNext() {
		c.Scan()
		c.Next()
	}
	return c.surface
}
