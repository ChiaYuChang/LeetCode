package leetcode

import "fmt"

type Q0415 struct{}

type Q0415Token struct {
	s *string
	i int
}

func (t *Q0415Token) Next() (b byte) {
	b = '0'
	if t.HasNext() {
		b = (*t.s)[t.i]
		t.i--
	}
	return b
}

func (t *Q0415Token) HasNext() bool {
	return t.i >= 0
}

func (t *Q0415Token) Len() int {
	return len(*t.s)
}

func (q Q0415) Parse(s string) *Q0415Token {
	return &Q0415Token{
		s: &s,
		i: len(s) - 1,
	}
}

func (q Q0415) AddByte(b1, b2 byte, carry bool) (b byte, c bool) {
	c1 := b1 - '0'
	c2 := b2 - '0'

	s := c1 + c2
	if carry {
		s += 1
	}

	if s >= 10 {
		return s - 10 + '0', true
	}
	return s + '0', false
}

func (q Q0415) Add(t1, t2 *Q0415Token) string {
	maxlen := 1
	if t1.Len() > t2.Len() {
		maxlen += t1.Len()
	} else {
		maxlen += t2.Len()
	}

	s := make([]byte, maxlen)
	c := false
	for i := maxlen - 1; i >= 0; i-- {
		s[i], c = q.AddByte(t1.Next(), t2.Next(), c)
	}

	if s[0] > '0' {
		fmt.Printf("%c\n", s)
		return string(s)
	}
	return string(s[1:])
}

func (q Q0415) AddStrings(num1 string, num2 string) string {
	tkn1 := q.Parse(num1)
	tkn2 := q.Parse(num2)
	return q.Add(tkn1, tkn2)
}
