package leetcode

type Q0999 struct{}

var Q0999Movements = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

const (
	Q0999WhiteRook   byte = 'R'
	Q0999WhiteBishop byte = 'B'
	Q0999BlackPawn   byte = 'p'
	Q0999Empty       byte = '.'
)

const (
	Q0999NRow int = 8
	Q0999NCol     = 8
)

type Q0999Board [][]byte

func (b Q0999Board) IsInRange(row, col int) bool {
	return row >= 0 && col >= 0 && row < Q0999NRow && col < Q0999NCol
}

func (b Q0999Board) FindRook() (int, int) {
	rookRow, rookCol := -1, -1
	for row := 0; row < Q0999NRow; row++ {
		for col := 0; col < Q0999NCol; col++ {
			if b[row][col] == Q0999WhiteRook {
				rookRow, rookCol = row, col
				break
			}
		}
	}
	return rookRow, rookCol
}

func (q Q0999) NumRookCaptures(board [][]byte) int {
	brd := Q0999Board(board)
	rookRow, rookCol := brd.FindRook()

	nAvailableCapture := 0
	for _, move := range Q0999Movements {
		row, col := rookRow+move[0], rookCol+move[1]
		for brd.IsInRange(row, col) {
			if brd[row][col] != Q0999Empty {
				if brd[row][col] == Q0999BlackPawn {
					nAvailableCapture++
				}
				break
			}
			row, col = row+move[0], col+move[1]
		}
	}
	return nAvailableCapture
}
