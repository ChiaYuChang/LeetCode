package leetcode

import (
	"strconv"
)

type Q043Token struct {
	s *[]byte
	i int
}

type Q043 struct{}

func (t *Q043Token) Next() (b byte) {
	b = '0'
	if t.HasNext() {
		b = (*t.s)[t.i]
		t.i--
	}
	return b
}

func (t *Q043Token) HasNext() bool {
	return t.i >= 0
}

func (t *Q043Token) Len() int {
	return len(*t.s)
}

func (t *Q043Token) Rewind() {
	t.i = t.Len() - 1
}

func (t Q043Token) String() string {
	return string(*t.s)
}

func (q Q043) ParseString(s string) *Q043Token {
	b := []byte(s)
	return &Q043Token{
		s: &b,
		i: len(s) - 1,
	}
}

func (q Q043) ParseBytes(b []byte) *Q043Token {
	return &Q043Token{
		s: &b,
		i: len(b) - 1,
	}
}

func (q Q043) AddByte(b1, b2 byte, carry bool) (b byte, c bool) {
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

func (q Q043) MultByte(b1, b2 byte, d1, d2 int) *Q043Token {
	c1 := b1 - '0'
	c2 := b2 - '0'

	p := int(c1) * int(c2)
	if p == 0 {
		return q.ParseString("0")
	}

	sp := strconv.Itoa(p)
	bs := make([]byte, len(sp)+d1+d2)
	for i := len(sp); i < len(sp)+d1+d2; i++ {
		bs[i] = '0'
	}
	copy(bs[:len(sp)], []byte(sp))
	return q.ParseBytes(bs)
}

func (q Q043) MultString(s1, s2 string) string {
	tkn1 := q.ParseString(s1)
	tkn2 := q.ParseString(s2)

	if tkn1.Len() < tkn2.Len() {
		tkn1, tkn2 = tkn2, tkn1
	}

	tkns := []*Q043Token{}
	var d1, d2 = 0, 0
	for tkn2.HasNext() {
		b2 := tkn2.Next()
		for tkn1.HasNext() {
			b1 := tkn1.Next()
			tkns = append(tkns, q.MultByte(b1, b2, d1, d2))
			d1++
		}

		tkn1.Rewind()
		d1 = 0
		d2++
	}

	ans := q.ParseString("0")
	for _, tkn := range tkns {
		ans = q.Add(ans, tkn)
	}
	return string(*ans.s)
}

func (q Q043) Add(t1, t2 *Q043Token) *Q043Token {
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
		return q.ParseBytes(s)
	}
	return q.ParseBytes(s[1:])
}
