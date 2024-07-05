package leetcode

type Q0498 struct{}

type Q0498Matrix struct {
	nums       [][]int
	nRow, nCol int
	row, col   int
	direction  bool
}

func (q Q0498) NewMatrix(mat [][]int) *Q0498Matrix {
	if len(mat) < 1 {
		return &Q0498Matrix{}
	}

	return &Q0498Matrix{
		nums: mat,
		nRow: len(mat),
		nCol: len(mat[0]),
	}
}

func (m Q0498Matrix) Len() int {
	return m.nRow * m.nCol
}

func (m *Q0498Matrix) Next() (int, bool) {
	if m.row >= m.nRow || m.col >= m.nCol {
		return 0, false
	}

	element := m.nums[m.row][m.col]
	if m.direction {
		m.row += 1
		m.col -= 1
	} else {
		m.row -= 1
		m.col += 1
	}

	if m.row == m.nRow {
		m.col += 2
		m.row -= 1
		m.direction = false
	}

	if m.col == m.nCol {
		m.row += 2
		m.col -= 1
		m.direction = true
	}

	if m.row == -1 {
		m.row += 1
		m.direction = true
	}

	if m.col == -1 {
		m.col += 1
		m.direction = false
	}

	return element, true
}

func (q Q0498) FindDiagonalOrder(mat [][]int) []int {
	mtrx := q.NewMatrix(mat)
	diagonalOrder := make([]int, mtrx.Len())
	for i := range diagonalOrder {
		if next, ok := mtrx.Next(); ok {
			diagonalOrder[i] = next
		} else {
			panic("length not match")
		}
	}
	return diagonalOrder
}
