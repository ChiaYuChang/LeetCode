package leetcode

import (
	"fmt"
	"strings"
)

type Q0020 struct{}

type Q0020TypeBracket int8

const (
	TRoundBracket Q0020TypeBracket = iota
	TSquareBracket
	TCurlyBracket
)

func (t Q0020TypeBracket) String() string {
	switch t {
	case TRoundBracket:
		return "Round"
	case TSquareBracket:
		return "Square"
	case TCurlyBracket:
		return "Curly"
	default:
		panic("unknown bracket")
	}
}

type Q0020TypeLR bool

const (
	TL Q0020TypeLR = false
	TR             = true
)

func (t Q0020TypeLR) String() string {
	if t == TL {
		return "left"
	} else {
		return "right"
	}
}

type Q0020BracketToken struct {
	Type  Q0020TypeBracket
	LR    Q0020TypeLR
	Level int
	Pair  int
}

func (t Q0020BracketToken) String() string {
	return fmt.Sprintf(
		"%s- Type: %6v: LR: %6v, Lev: %d Pair %d",
		strings.Repeat("\t", t.Level),
		t.Type, t.LR, t.Level, t.Pair,
	)
}

type Q0020BracketNode struct {
	Type  Q0020TypeBracket
	Nodes []*Q0020BracketNode
}

func (n *Q0020BracketNode) ToString(lvl int, sb *strings.Builder) {
	if n == nil {
		return
	}

	sb.WriteString(fmt.Sprintf("%s- Type: %6v (%d)\n", strings.Repeat("\t", lvl), n.Type, len(n.Nodes)))
	for _, c := range n.Nodes {
		c.ToString(lvl+1, sb)
	}
}

func (n Q0020BracketNode) String() string {
	sb := strings.Builder{}
	sb.WriteString("BracketNodes:\n")
	n.ToString(0, &sb)
	return sb.String()
}

type Q0020BracketParser struct{}

func (p Q0020BracketParser) TokenType(b byte) (Q0020TypeBracket, Q0020TypeLR) {
	switch b {
	case '(':
		return TRoundBracket, TL
	case ')':
		return TRoundBracket, TR
	case '[':
		return TSquareBracket, TL
	case ']':
		return TSquareBracket, TR
	case '{':
		return TCurlyBracket, TL
	case '}':
		return TCurlyBracket, TR
	default:
		panic("unknown char")
	}
}

func (p Q0020BracketParser) Parse(s string) ([]*Q0020BracketToken, bool) {
	if len(s) == 0 {
		return nil, true
	}

	tokens := make([]*Q0020BracketToken, 0, len(s))
	l := 0
	m := map[int][]int{}
	for i := 0; i < len(s); i++ {
		t, lr := p.TokenType(s[i])

		if bool(lr) {
			// right
			l--
			if end := len(m[l]) - 1; end >= 0 {
				j := m[l][end]
				m[l] = m[l][:end]
				if tokens[j].Type != t {
					return nil, false
				}
				tokens[j].Pair = i - j
				// padding token
				tokens = append(tokens, nil)
			} else {
				return nil, false
			}
		} else {
			// left
			b := &Q0020BracketToken{Type: t, LR: lr}
			b.Level = l
			m[l] = append(m[l], i)
			l++
			tokens = append(tokens, b)
		}
	}

	for _, v := range m {
		if len(v) > 0 {
			return nil, false
		}
	}

	return tokens, true
}

func (p Q0020BracketParser) Lext(tokens []*Q0020BracketToken) []*Q0020BracketNode {
	nodes := []*Q0020BracketNode{}
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		n := &Q0020BracketNode{Type: token.Type}
		n.Nodes = p.Lext(tokens[i+1 : i+token.Pair])
		nodes = append(nodes, n)
		i = i + token.Pair
	}
	return nodes
}

func (q Q0020) IsValid(s string) bool {
	p := Q0020BracketParser{}
	_, ok := p.Parse(s)
	return ok
}
