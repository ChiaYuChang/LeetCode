package leetcode

import (
	"fmt"
)

type MiniExprParser struct{}

type TokenType uint8

const (
	tInt TokenType = iota
	tLBracket
	tRBracket
	tComma
)

type OperatorType uint8

const (
	tAppend OperatorType = iota
)

type Element interface {
	Value() any // should be either int or []int
}

type Integer[T int | []int] struct {
	value T
}

func NewInteger[T int | []int](val T) Integer[T] {
	return Integer[T]{value: val}
}

func (i Integer[T]) Value() any {
	return i.value
}

type Token struct {
	Type TokenType
	Text string
}

func (t Token) String() string {
	return fmt.Sprintf("`%s`", t.Text)
}

func (t Token) Is(tp TokenType) bool {
	return t.Type == tp
}

func IsNumber(b byte) bool {
	if '0' <= b && b <= '9' {
		return true
	}
	return false
}

func (p MiniExprParser) Lext(expr string) []*Token {
	t := make([]*Token, 0)
	for i := 0; i < len(expr); i++ {
		switch expr[i] {
		case '[':
			t = append(t, &Token{tLBracket, "["})
		case ']':
			t = append(t, &Token{tRBracket, "]"})
		case ',':
			t = append(t, &Token{tComma, ","})
		default:
			// should be integers
			var j int
			for j = i; j < len(expr); j++ {
				if !IsNumber(expr[j]) {
					break
				}
			}
			t = append(t, &Token{tInt, expr[i:j]})
			i = j - 1
		}
	}
	return t
}

// func (p MiniExprParser) Parse(tokens []*Token) []Element {
// 	e := []Element{}
// 	for i := 0; i < len(tokens); i++ {
// 		switch tokens[i].Type {
// 		case tLBracket:
// 			j := i
// 			for l := 1; j < len(tokens) && l > 0; j++ {
// 				if tokens[i].Is(tLBracket) {
// 					l++
// 				} else if tokens[i].Is(tRBracket) {
// 					l--
// 				}
// 			}
// 			subexpr := p.Parse(tokens[i+1 : j-1])
// 		case tInt:
// 			i, _ := strconv.Atoi(tokens[i].Text)
// 			e = append(e, NewInteger(i))
// 		case tComma:
// 			continue
// 		default:
// 			panic("Not supported operator")
// 		}
// 	}
// }

// func deserialize(s string) *NestedInteger {

// }
