package leetcode

type Q2075 struct{}

type Q2075Grid struct {
	text       string
	nRow, nCol int
	row, col   int
}

func (q Q2075) NewGrid(text string, nRow int) *Q2075Grid {
	return &Q2075Grid{
		text: text,
		nRow: nRow,
		nCol: len(text) / nRow,
		row:  0,
		col:  0,
	}
}

func (g *Q2075Grid) Next() (byte, bool) {
	if g.col >= g.nCol {
		return ' ', false
	}
	b := g.text[g.row*g.nCol+g.col]
	g.row++
	g.col++
	if g.row%g.nRow == 0 {
		g.row = 0
		g.col = g.col - g.nRow + 1
	}
	return b, true
}

func (q Q2075) DecodeCiphertext(encodedText string, rows int) string {
	g := q.NewGrid(encodedText, rows)
	originalText := make([]byte, 0, len(encodedText))
	for next, ok := g.Next(); ok; next, ok = g.Next() {
		originalText = append(originalText, next)
	}

	lastNotSpace := len(originalText) - 1
	for ; lastNotSpace >= 0; lastNotSpace-- {
		if originalText[lastNotSpace] != ' ' {
			break
		}
	}

	return string(originalText[:lastNotSpace+1])
}
