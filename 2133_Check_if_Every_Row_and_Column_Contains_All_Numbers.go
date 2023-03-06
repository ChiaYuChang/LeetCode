package leetcode

func checkValid(matrix [][]int) bool {
	m := q2133Matrix(matrix)

	for i := 0; i < m.N(); i++ {
		if !(m.Check(m.Row(i)...) && m.Check(m.Col(i)...)) {
			return false
		}
	}
	return true
}

type q2133Matrix [][]int

func (m q2133Matrix) N() int {
	return len(m)
}

func (m q2133Matrix) Row(i int) []int {
	return m[i]
}

func (m q2133Matrix) Col(i int) []int {
	col := make([]int, m.N())
	for r := 0; r < m.N(); r++ {
		col[r] = m[r][i]
	}
	return col
}

func (m q2133Matrix) Check(is ...int) bool {
	checklist := make([]bool, m.N())
	for _, i := range is {
		if checklist[i-1] {
			return false
		} else {
			checklist[i-1] = true
		}
	}
	return true
}
