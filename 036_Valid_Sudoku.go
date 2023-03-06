package leetcode

func IsValidSudoku(board [][]byte) bool {
	mtrx := Matrix(board)
	for i := 0; i < 9; i++ {
		if !(check(mtrx.Row(i)...) && check(mtrx.Col(i)...) && check(mtrx.Block(i/3, i%3)...)) {
			return false
		}
	}
	return true
}

type Matrix [][]byte

func (m Matrix) Row(i int) []byte {
	row := make([]byte, 0, 9)
	for c := 0; c < 9; c++ {
		if m[i][c] != '.' {
			row = append(row, m[i][c])
		}
	}
	return row
}

func (m Matrix) Col(i int) []byte {
	col := make([]byte, 0, 9)
	for r := 0; r < 9; r++ {
		if m[r][i] != '.' {
			col = append(col, m[r][i])
		}
	}
	return col
}

func (m Matrix) Block(i, j int) []byte {
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

func check(bs ...byte) bool {
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
