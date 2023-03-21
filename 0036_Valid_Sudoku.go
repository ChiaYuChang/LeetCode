package leetcode

type Q0036 struct{}

func (q Q0036) IsValidSudoku(board [][]byte) bool {
	mtrx := Q0036Matrix(board)
	for i := 0; i < 9; i++ {
		if !(q.check(mtrx.Row(i)...) && q.check(mtrx.Col(i)...) && q.check(mtrx.Block(i/3, i%3)...)) {
			return false
		}
	}
	return true
}

type Q0036Matrix [][]byte

func (m Q0036Matrix) Row(i int) []byte {
	row := make([]byte, 0, 9)
	for c := 0; c < 9; c++ {
		if m[i][c] != '.' {
			row = append(row, m[i][c])
		}
	}
	return row
}

func (m Q0036Matrix) Col(i int) []byte {
	col := make([]byte, 0, 9)
	for r := 0; r < 9; r++ {
		if m[r][i] != '.' {
			col = append(col, m[r][i])
		}
	}
	return col
}

func (m Q0036Matrix) Block(i, j int) []byte {
	bs := make([]byte, 0, 9)
	for row := i * 3; row < (i+1)*3; row++ {
		for col := j * 3; col < (j+1)*3; col++ {
			if m[row][col] != '.' {
				bs = append(bs, m[row][col])
			}
		}
	}
	return bs
}

func (q Q0036) check(bs ...byte) bool {
	checklist := make([]bool, 9)
	for _, b := range bs {
		if checklist[b-'1'] {
			return false
		} else {
			checklist[b-'1'] = true
		}
	}
	return true
}
