package leetcode

type Q0427 struct{}

type Q0427QutraTreeNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Q0427QutraTreeNode
	TopRight    *Q0427QutraTreeNode
	BottomLeft  *Q0427QutraTreeNode
	BottomRight *Q0427QutraTreeNode
}

type Q0427TwoDMatrix struct {
	nums                   *[][]int
	rstr, rend, cstr, cend int
	nrow, ncol             int
}

func (q Q0427) NewTwoDMatrix(nums [][]int) *Q0427TwoDMatrix {
	return &Q0427TwoDMatrix{
		nums: &nums,
		rstr: 0,
		rend: len(nums),
		cstr: 0,
		cend: len(nums[0]),
		nrow: len(nums),
		ncol: len(nums),
	}
}

// return the number of col of the matrix
func (m *Q0427TwoDMatrix) NCol() int {
	return m.ncol
}

// return the number of row of the matrix
func (m *Q0427TwoDMatrix) NRow() int {
	return m.nrow
}

// return a submarix of the original matrix, includes both boundaries [rstr, rend], [cstr, cend]
func (m *Q0427TwoDMatrix) SubMatrix(rstr, rend, cstr, cend int) (*Q0427TwoDMatrix, bool) {
	if rstr < 0 || cstr < 0 || rend >= m.NRow() || cend >= m.NCol() {
		return nil, false
	}

	return &Q0427TwoDMatrix{
		nums: m.nums,
		rstr: m.rstr + rstr,
		rend: m.rstr + rend + 1,
		cstr: m.cstr + cstr,
		cend: m.cstr + cend + 1,
		nrow: rend - rstr + 1,
		ncol: cend - cstr + 1,
	}, true
}

// return a submarix of the original matrix, and omit check
func (m *Q0427TwoDMatrix) MustSubMatrix(rstr, rend, cstr, cend int) *Q0427TwoDMatrix {
	subm, _ := m.SubMatrix(rstr, rend, cstr, cend)
	return subm
}

// return element m[row][col]
func (m *Q0427TwoDMatrix) Element(row, col int) (e int, ok bool) {
	if row < 0 || col < 0 || row >= m.NRow() || col >= m.NCol() {
		return 0, false
	}
	return (*m.nums)[m.rstr+row][m.cstr+col], true
}

// if m[row][col] == 0 return false, else return 1
func (m *Q0427TwoDMatrix) ElementBinary(row, col int) (be, ok bool) {
	e, ok := m.Element(row, col)
	return e == 1, ok
}

// check whether all elements in the matrix is the same
func (m *Q0427TwoDMatrix) IsHomogeneous() (bool, bool) {
	e0, ok := m.Element(0, 0)
	if !ok {
		return false, false
	}

	for row := 0; row < m.NRow(); row++ {
		for col := 0; col < m.NCol(); col++ {
			if e1, ok := m.Element(row, col); ok && e0 != e1 {
				return e0 == 1, false
			}
		}
	}
	return e0 == 1, true
}

func (qtb Q0427) Build(mtrx *Q0427TwoDMatrix) *Q0427QutraTreeNode {
	n := &Q0427QutraTreeNode{}
	n.Val, n.IsLeaf = mtrx.IsHomogeneous()

	if !n.IsLeaf {
		midrow := mtrx.NRow()/2 - 1
		midcol := mtrx.NCol()/2 - 1
		n.TopLeft = qtb.Build(
			mtrx.MustSubMatrix(0, midrow, 0, midcol))
		n.TopRight = qtb.Build(
			mtrx.MustSubMatrix(0, midrow, midcol+1, mtrx.NCol()-1))
		n.BottomLeft = qtb.Build(
			mtrx.MustSubMatrix(midrow+1, mtrx.NRow()-1, 0, midcol))
		n.BottomRight = qtb.Build(
			mtrx.MustSubMatrix(midrow+1, mtrx.NRow()-1, midcol+1, mtrx.NCol()-1))
	}
	return n
}

func (qtb Q0427) QutraTree(grid [][]int) *Q0427QutraTreeNode {
	mtrx := qtb.NewTwoDMatrix(grid)
	return qtb.Build(mtrx)
}
