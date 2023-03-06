package leetcode

// import (
// 	"errors"
// )

// type SudokuChecker struct {
// 	size          int
// 	bLength       int
// 	bSize         int
// 	board         [][]byte
// 	rowCandidates []int
// 	colCandidates []int
// 	mapb2i        map[byte]int
// 	mapi2b        map[int]byte
// 	candidates    []byte
// }

// type q036Set struct {
// 	len int
// 	set map[byte]struct{}
// }

// // add one element to the set, return false if the element has
// // already existed
// func (s *q036Set) AddOne(b byte) bool {
// 	_, ok := s.set[b]
// 	if ok {
// 		return false
// 	}
// 	s.set[b] = struct{}{}
// 	s.len++
// 	return true
// }

// // add elements to the set
// func (s *q036Set) Add(i ...byte) {
// 	for _, j := range i {
// 		s.AddOne(j)
// 	}
// }

// // remove given element from the set if exist
// func (s *q036Set) RemoveOne(b byte) {
// 	if s.IsEmpty() {
// 		return
// 	}

// 	_, ok := s.set[b]
// 	if !ok {
// 		return
// 	}
// 	delete(s.set, b)
// 	s.len--
// }

// // remove given elements from the set if exit
// func (s *q036Set) Remove(b ...byte) {
// 	for _, j := range b {
// 		s.RemoveOne(j)
// 	}
// }

// // return whether the set is empty
// func (s *q036Set) IsEmpty() bool {
// 	return s.len == 0
// }

// func (s *q036Set) In(b byte) bool {
// 	_, ok := s.set[b]
// 	return ok
// }

// func (s *q036Set) Len() int {
// 	return s.len
// }

// // convert the set to an array
// func (s *q036Set) ToList() []byte {
// 	bs := make([]byte, 0, s.len)
// 	for b := range s.set {
// 		bs = append(bs, b)
// 	}
// 	return bs
// }

// func (s1 *q036Set) Intersection(s2 *q036Set) *q036Set {
// 	bs := Newq036Set()
// 	for b := range s2.set {
// 		if s1.In(b) {
// 			bs.AddOne(b)
// 		}
// 	}
// 	return bs
// }

// func (s1 *q036Set) Difference(s2 *q036Set) *q036Set {
// 	bs := Newq036Set()
// 	for b := range s2.set {
// 		if !s1.In(b) {
// 			bs.AddOne(b)
// 		}
// 	}
// 	return bs
// }

// func (s1 *q036Set) SymmetricDifference(s2 *q036Set) []byte {
// 	bs := Newq036Set()
// 	bs.Add(s1.Difference(s2).ToList()...)
// 	bs.Add(s2.Difference(s1).ToList()...)
// 	return bs.ToList()
// }

// // create a new set struct
// func Newq036Set() *q036Set {
// 	return &q036Set{len: 0, set: map[byte]struct{}{}}
// }

// func ByteToInt(b byte) int {
// 	return int(b - '0')
// }

// func NewSudokuChecker(size int, board [][]byte, candidates []byte) (*SudokuChecker, error) {
// 	if len(board) == 0 {
// 		return nil, nil
// 	}

// 	if len(board) != len(board[0]) {
// 		return nil, errors.New("A Valid Sudoku should be an nxn matrix")
// 	}

// 	if len(board) != size*size {
// 		return nil, errors.New("given size and board do not match")
// 	}

// 	mb2i := map[byte]int{}
// 	for i, b := range candidates {
// 		mb2i[b] = i
// 	}

// 	mi2b := map[int]byte{}
// 	for i, b := range candidates {
// 		mi2b[i] = b
// 	}

// 	return &SudokuChecker{
// 		board:         board,
// 		size:          size,
// 		rowCandidates: nil,
// 		colCandidates: nil,
// 		mapb2i:        mb2i,
// 		mapi2b:        mi2b,
// 		candidates:    candidates,
// 	}, nil
// }

// func (sc *SudokuChecker) Size() int {
// 	return sc.size
// }

// func (sc *SudokuChecker) BLength() int {
// 	if sc.bLength == 0 {
// 		sc.bLength = len(sc.board)
// 	}
// 	return sc.bLength
// }

// func (sc *SudokuChecker) BSize() int {
// 	if sc.bSize == 0 {
// 		sc.bSize = sc.BLength() * sc.BLength()
// 	}
// 	return sc.bSize
// }

// func (sc *SudokuChecker) calcCandidatesIndex(ns ...byte) []int {
// 	check := make([]bool, sc.BLength())
// 	nc := sc.BLength()
// 	for _, n := range ns {
// 		check[sc.mapb2i[n]] = true
// 		nc--
// 	}

// 	cs := make([]int, 0, nc)
// 	for i, c := range check {
// 		if !c {
// 			cs = append(cs, i)
// 		}
// 	}
// 	return cs
// }

// func (sc *SudokuChecker) CalcCandidates(ns ...byte) *q036Set {
// 	idxs := sc.calcCandidatesIndex(ns...)
// 	bs := Newq036Set()
// 	for _, c := range idxs {
// 		bs.AddOne(sc.mapi2b[c])
// 	}
// 	return bs
// }

// func (sc *SudokuChecker) Row(i int) []byte {
// 	if i < 0 || i > sc.BLength() {
// 		return nil
// 	}
// 	return sc.board[i]
// }

// func (sc *SudokuChecker) Col(i int) []byte {
// 	if i < 0 || i > sc.BLength() {
// 		return nil
// 	}

// 	bs := make([]byte, sc.BLength())
// 	for row := 0; row < sc.BLength(); row++ {
// 		bs[row] = sc.board[row][i]
// 	}
// 	return bs
// }

// func (sc *SudokuChecker) Block(i int, j int) []byte {
// 	if i < 0 || i > sc.Size() || j < 0 || j > sc.Size() {
// 		return nil
// 	}

// 	bs := make([]byte, 0, sc.BLength())
// 	for row := i * sc.Size(); row < (i+1)*sc.Size(); row++ {
// 		for col := j * sc.Size(); col < (j+1)*sc.Size(); col++ {
// 			bs = append(bs, sc.board[row][col])
// 		}
// 	}
// 	return bs
// }

// func (sc *SudokuChecker) Valid() bool {
// 	row := make([]*q036Set, sc.BLength())
// 	for i := 0; i < sc.BLength(); i++ {
// 		row[i] = sc.CalcCandidates(sc.Row(i)...)
// 	}

// 	col := make([]*q036Set, sc.BLength())
// 	for i := 0; i < sc.BLength(); i++ {
// 		col[i] = sc.CalcCandidates(sc.Col(i)...)
// 	}

// 	block := make([][]*q036Set, sc.Size())
// 	for i := 0; i < sc.Size(); i++ {
// 		for j := 0; i < sc.Size(); i++ {
// 			block[i] = append(block[i], sc.CalcCandidates(sc.Block(i, j)...))
// 		}
// 	}

// 	for i, r := range row {
// 		for j, c := range col {
// 			if sc.board[i][j] == '0' && r.Intersection(c).Intersection(block[i][j]).Len() < 1 {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
