package leetcode

func SurfaceArea(grid [][]int) int {
	c := NewCubes(grid)
	return c.CountSurfaceArea()
}

type Cubes struct {
	surface int
	grid    [][]int
	maxl    int
	l       int
}

func NewCubes(grid [][]int) *Cubes {
	c := &Cubes{grid: grid}
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

func (c *Cubes) NRow() int {
	return len(c.grid)
}

func (c *Cubes) NCol() int {
	return len(c.grid[0])
}

func (c *Cubes) Get(col, row int) int {
	if col < 0 || row < 0 || col >= c.NCol() || row >= c.NRow() {
		return 0
	}
	return c.grid[row][col]
}

func (c *Cubes) IsZero(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}

func (c *Cubes) NZeroNeighbors(col, row int) int {
	u := c.IsZero(c.Get(col, row-1))
	d := c.IsZero(c.Get(col, row-1))
	l := c.IsZero(c.Get(col-1, row))
	r := c.IsZero(c.Get(col+1, row))
	return u + d + l + r
}

func (c *Cubes) HasNext() (ok bool) {
	return c.maxl > c.l
}

func (c *Cubes) Next() (ok bool) {
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

func (c *Cubes) Scan() {
	for row := 0; row < c.NRow(); row++ {
		for col := 0; col < c.NCol(); col++ {
			if c.Get(row, col) > 0 {
				c.surface += c.NZeroNeighbors(row, col)
			}
		}
	}
}

func (c *Cubes) CountSurfaceArea() int {
	for c.HasNext() {
		c.Scan()
		c.Next()
	}
	return c.surface
}
