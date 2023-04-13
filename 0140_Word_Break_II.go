package leetcode

import (
	"fmt"
	"strings"
)

type Q0140 struct{}

type Q0140Node struct {
	i     int
	nodes [27]*Q0140Node
}

func (n Q0140Node) Next(b byte) *Q0140Node {
	return n.nodes[b-'a'+1]
}

func (n Q0140Node) IsEnd() *Q0140Node {
	return n.nodes[0]
}

func (n *Q0140Node) Append(b byte) (next *Q0140Node) {
	i := b - 'a' + 1
	if n.nodes[i] == nil {
		n.nodes[i] = &Q0140Node{}
	}
	return n.nodes[i]
}

func (n Q0140Node) String() string {
	var ToString func(n *Q0140Node, char byte, lvl int, sb *strings.Builder)

	ToString = func(n *Q0140Node, char byte, lvl int, sb *strings.Builder) {
		sb.WriteString(strings.Repeat(" ", lvl))
		sb.WriteByte(char)
		sb.WriteRune('\n')

		if n.nodes[0] != nil {
			sb.WriteString(strings.Repeat(" ", lvl+1))
			sb.WriteString("end\n")
		}

		for i := 1; i < len(n.nodes); i++ {
			if n.nodes[i] != nil {
				ToString(n.nodes[i], byte('a'+i-1), lvl+1, sb)
			}
		}
	}

	sb := strings.Builder{}
	sb.WriteString("root\n")
	for i := 1; i < 27; i++ {
		if next := n.nodes[i]; next != nil {
			ToString(next, byte('a'+i-1), 1, &sb)
		}
	}
	return sb.String()
}

type Q0140PrefixTree struct {
	root *Q0140Node
}

func (q Q0140) NewPrefixTree() *Q0140PrefixTree {
	return &Q0140PrefixTree{root: &Q0140Node{}}
}

func (p *Q0140PrefixTree) Append(s string, i int) {
	curr := p.root
	for j := 0; j < len(s); j++ {
		curr = curr.Append(s[j])
	}
	curr.i = i
	curr.nodes[0] = p.root
}

func (p Q0140PrefixTree) String() string {
	return p.root.String()
}

func (p Q0140PrefixTree) Next(b byte) *Q0140Node {
	return p.root.Next(b)
}

type Q0140Tracer struct {
	wordsIdx []int
	node     *Q0140Node
}

func (t Q0140Tracer) String() string {
	return fmt.Sprintf("{%d}", t.wordsIdx)
}

func (q Q0140) WordBreak(s string, wordDict []string) (words []string) {
	pt := q.NewPrefixTree()
	for i, w := range wordDict {
		pt.Append(w, i)
	}
	// fmt.Println(pt)

	prev := []*Q0140Tracer{{wordsIdx: nil, node: pt.root}}
	for i := 0; i < len(s); i++ {
		// fmt.Printf("%c(%02d) %v\n", s[i], i, prev)
		curr := []*Q0140Tracer{}
		for j := range prev {
			t := prev[j]
			if next := t.node.Next(s[i]); next != nil {
				// fmt.Printf("\t- Append %c\n", s[i])
				t.node = next
				if i != len(s)-1 {
					// do not append to prev when the last element is selected
					curr = append(curr, t)
				}

				if root := t.node.IsEnd(); root != nil {
					newT := &Q0140Tracer{wordsIdx: make([]int, len(t.wordsIdx)+1), node: root}
					copy(newT.wordsIdx, t.wordsIdx)
					newT.wordsIdx[len(t.wordsIdx)] = t.node.i
					// fmt.Printf("\t- Add Word %s\n", wordDict[t.node.i])
					curr = append(curr, newT)
				}
			}
		}
		if len(curr) == 0 {
			// fmt.Println("Cound not reach the end")
			return words
		}
		prev = curr
	}

	// fmt.Println(prev)
	words = make([]string, len(prev))
	for i := range words {
		ss := make([]string, len(prev[i].wordsIdx))
		for j := range ss {
			ss[j] = wordDict[prev[i].wordsIdx[j]]
		}
		words[i] = strings.Join(ss, " ")
	}
	return words
}
