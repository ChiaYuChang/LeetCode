package leetcode

type Q1314 struct{}

type Q1314Matrix struct {
	value *[][]int
	nrow  int
	ncol  int
}

func NewQ1314Matrix(m *[][]int) Q1314Matrix {
	nrow, ncol := len(*m), 0
	if len(*m) > 0 {
		ncol = len((*m)[0])
	}

	return Q1314Matrix{
		value: m, nrow: nrow, ncol: ncol,
	}
}

func (m Q1314Matrix) Element(row, col int) (e int, ok bool) {
	if row < 0 || row >= m.nrow || col < 0 || col >= m.ncol {
		return 0, false
	}
	return (*m.value)[row][col], true
}

func (m Q1314Matrix) BlockSum(row, col, k int) int {
	s := 0
	for i := row - k; i <= row+k; i++ {
		for j := col - k; j <= col+k; j++ {
			e, _ := m.Element(i, j)
			s += e
		}
	}
	return s
}

func (q Q1314) MatrixBlockSum(mat [][]int, k int) [][]int {
	if k == 0 {
		return mat
	}

	m := NewQ1314Matrix(&mat)
	sm := make([][]int, m.nrow)
	for i := 0; i < m.nrow; i++ {
		sm[i] = make([]int, m.ncol)
		for j := 0; j < m.ncol; j++ {
			sm[i][j] = m.BlockSum(i, j, k)
		}
	}
	return sm
}
