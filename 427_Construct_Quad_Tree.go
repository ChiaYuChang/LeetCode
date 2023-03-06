package leetcode

type QutraTreeNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QutraTreeNode
	TopRight    *QutraTreeNode
	BottomLeft  *QutraTreeNode
	BottomRight *QutraTreeNode
}

type QutraTreeBuilder struct{}

type TwoDMatrix struct {
	nums                   *[][]int
	rstr, rend, cstr, cend int
	nrow, ncol             int
}

func NewTwoDMatrix(nums [][]int) *TwoDMatrix {
	return &TwoDMatrix{
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
func (m *TwoDMatrix) NCol() int {
	return m.ncol
}

// return the number of row of the matrix
func (m *TwoDMatrix) NRow() int {
	return m.nrow
}

// return a submarix of the original matrix, includes both boundaries [rstr, rend], [cstr, cend]
func (m *TwoDMatrix) SubMatrix(rstr, rend, cstr, cend int) (*TwoDMatrix, bool) {
	if rstr < 0 || cstr < 0 || rend >= m.NRow() || cend >= m.NCol() {
		return nil, false
	}

	return &TwoDMatrix{
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
func (m *TwoDMatrix) MustSubMatrix(rstr, rend, cstr, cend int) *TwoDMatrix {
	subm, _ := m.SubMatrix(rstr, rend, cstr, cend)
	return subm
}

// return element m[row][col]
func (m *TwoDMatrix) Element(row, col int) (e int, ok bool) {
	if row < 0 || col < 0 || row >= m.NRow() || col >= m.NCol() {
		return 0, false
	}
	return (*m.nums)[m.rstr+row][m.cstr+col], true
}

// if m[row][col] == 0 return false, else return 1
func (m *TwoDMatrix) ElementBinary(row, col int) (be, ok bool) {
	e, ok := m.Element(row, col)
	return e == 1, ok
}

// check whether all elements in the matrix is the same
func (m *TwoDMatrix) IsHomogeneous() (bool, bool) {
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

func (qtb QutraTreeBuilder) Build(mtrx *TwoDMatrix) *QutraTreeNode {
	n := &QutraTreeNode{}
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

func (qtb QutraTreeBuilder) QutraTree(grid [][]int) *QutraTreeNode {
	mtrx := NewTwoDMatrix(grid)
	return qtb.Build(mtrx)
}
