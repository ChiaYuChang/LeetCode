package leetcode

import (
	"fmt"
	"strings"
)

type Q0794 struct{}

type Q0794Token struct {
	b byte
	i int8
}

var (
	Q0794Empty  = Q0794Token{' ', 0b00}
	Q0794Nought = Q0794Token{'O', 0b10}
	Q0794Cross  = Q0794Token{'X', 0b01}
)

type Q0794Board struct {
	token                         [][]int8
	numEmpty, numNought, numCross int
	Size                          int
}

func (q Q0794) NewBoard(board []string) *Q0794Board {
	size := len(board)
	b := &Q0794Board{
		token: make([][]int8, size),
		Size:  size,
	}

	for row := 0; row < size; row++ {
		b.token[row] = make([]int8, size)
		for col := range board[row] {
			switch board[row][col] {
			case Q0794Nought.b:
				b.numNought++
				b.token[row][col] = Q0794Nought.i
			case Q0794Cross.b:
				b.numCross++
				b.token[row][col] = Q0794Cross.i
			case Q0794Empty.b:
				b.numEmpty++
				b.token[row][col] = Q0794Empty.i
			default:
				panic("unknown character")
			}
		}
	}
	return b
}

func (b Q0794Board) String() string {
	sb := strings.Builder{}
	sb.WriteString("Tic Tac Toe Board:\n")
	for row := 0; row < b.Size; row++ {
		sb.WriteRune('[')
		for col := 0; col < b.Size; col++ {
			switch b.token[row][col] {
			case Q0794Empty.i:
				sb.WriteByte(Q0794Empty.b)
			case Q0794Nought.i:
				sb.WriteByte(Q0794Nought.b)
			case Q0794Cross.i:
				sb.WriteByte(Q0794Cross.b)
			default:
				panic("unknown char")
			}
			if col != b.Size-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("]  ")
		switch row {
		case 0:
			sb.WriteString(fmt.Sprintf("# of Cross : %d\n", b.numCross))
		case 1:
			sb.WriteString(fmt.Sprintf("# of Nought: %d\n", b.numNought))
		case 2:
			sb.WriteString(fmt.Sprintf("# of Empty : %d\n", b.numEmpty))
		}
	}
	return sb.String()
}

func (b Q0794Board) CheckRowI(row int) int8 {
	winner := b.token[row][0]
	for col := 0; winner > 0 && col < b.Size; col++ {
		winner &= b.token[row][col]
	}
	return winner
}

func (b Q0794Board) CheckColI(col int) int8 {
	winner := b.token[0][col]
	for row := 0; winner > 0 && row < b.Size; row++ {
		winner &= b.token[row][col]
	}
	return winner
}

func (b Q0794Board) CheckDiag(direction bool) int8 {
	col, next := 0, 1
	if direction {
		col, next = b.Size-1, -1
	}

	winner := b.token[0][col]
	for row := 0; winner > 0 && row < b.Size; row, col = row+1, col+next {
		winner &= b.token[row][col]
	}
	return winner
}

func (b Q0794Board) Validate() bool {
	var winner int8
	for col := 0; col < b.Size; col++ {
		winner |= b.CheckColI(col)
	}

	for row := 0; row < b.Size; row++ {
		winner |= b.CheckRowI(row)
	}

	winner |= b.CheckDiag(true)
	winner |= b.CheckDiag(false)

	switch winner {
	case Q0794Empty.i:
		if b.numCross-b.numNought == 0 || b.numCross-b.numNought == 1 {
			return true
		}
	case Q0794Cross.i:
		if b.numCross-b.numNought == 1 {
			return true
		}
	case Q0794Nought.i:
		if b.numCross-b.numNought == 0 {
			return true
		}
	}
	return false
}

func (q Q0794) ValidTicTacToe(board []string) bool {
	return q.NewBoard(board).Validate()
}
