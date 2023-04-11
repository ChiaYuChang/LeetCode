package leetcode

import (
	"strings"
)

type Q1462 struct{}

func (q Q1462) CheckIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	m := NewBooleanMatrix(numCourses, numCourses)
	m.SetDiagonal(true)
	for _, e := range prerequisites {
		m.Set(e[0], e[1], true)
	}
	m, _ = m.ToConverge(100) // converge in 2^100 steps

	isRequired := make([]bool, len(queries))
	for i, q := range queries {
		isRequired[i] = m.Get(q[0], q[1])
	}
	return isRequired
}

type BooleanMatrix struct {
	data       [][]bool
	nrow, ncol int
}

func NewBooleanMatrix(nrow, ncol int) BooleanMatrix {
	data := make([][]bool, nrow)
	for row := 0; row < nrow; row++ {
		data[row] = make([]bool, ncol)
	}
	return BooleanMatrix{data: data, nrow: nrow, ncol: ncol}
}

func (m *BooleanMatrix) SetDiagonal(b bool) {
	n := m.ncol
	if n > m.nrow {
		n = m.nrow
	}
	for i := 0; i < n; i++ {
		m.data[i][i] = b
	}
}

func (m *BooleanMatrix) Set(row, col int, b bool) {
	m.data[row][col] = b
}

func (m BooleanMatrix) Get(row, col int) bool {
	return m.data[row][col]
}

func (m BooleanMatrix) ToConverge(maxitr int) (BooleanMatrix, bool) {
	m1 := m
	m2 := m.Mult(m)
	for maxitr >= 0 && !m1.Equal(m2) {
		m1, m2 = m2, m2.Mult(m2)
		maxitr--
	}

	return m1, maxitr > 0
}

func (m1 BooleanMatrix) Transpose() BooleanMatrix {
	m2 := NewBooleanMatrix(m1.nrow, m1.ncol)
	for row := 0; row < m1.nrow; row++ {
		for col := 0; row < m1.ncol; col++ {
			m2.data[col][row] = m1.data[row][col]
		}
	}
	return m2
}

func (m1 BooleanMatrix) Mult(m2 BooleanMatrix) BooleanMatrix {
	if m1.nrow != m2.ncol || m1.ncol != m2.nrow {
		panic("non-conformable arguments")
	}
	m := NewBooleanMatrix(m1.nrow, m2.ncol)

	for row := 0; row < m.nrow; row++ {
		for col := 0; col < m.ncol; col++ {
			for i := 0; !m.data[row][col] && i < m1.ncol; i++ {
				m.data[row][col] = m.data[row][col] || (m1.data[row][i] && m2.data[i][col])
			}
		}
	}
	return m
}

func (m1 BooleanMatrix) Equal(m2 BooleanMatrix) bool {
	if m1.nrow != m2.nrow || m1.ncol != m2.ncol {
		return false
	}

	for row := 0; row < m1.nrow; row++ {
		for col := 0; col < m1.ncol; col++ {
			if m1.data[row][col] != m2.data[row][col] {
				return false
			}
		}
	}
	return true
}

func (m BooleanMatrix) String() string {
	sb := strings.Builder{}
	sb.WriteString("Boolean Matrix:\n")
	for row := 0; row < m.nrow; row++ {
		for col := 0; col < m.ncol; col++ {
			if m.data[row][col] {
				sb.WriteString("1, ")
			} else {
				sb.WriteString("0, ")
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
