package leetcode

import (
	"fmt"
	"strings"
)

type Q0140 struct{}

type Q0140Node struct {
	wIdx  int
	IsEnd bool
	nodes [26]*Q0140Node
}

func (n Q0140Node) Next(b byte) *Q0140Node {
	return n.nodes[b-'a']
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
				ToString(n.nodes[i], byte('a'+i), lvl+1, sb)
			}
		}
	}

	sb := strings.Builder{}
	sb.WriteString("root\n")
	for i := 0; i < 26; i++ {
		if next := n.nodes[i]; next != nil {
			ToString(next, byte('a'+i), 1, &sb)
		}
	}
	return sb.String()
}

func (n *Q0140Node) Append(b byte) (next *Q0140Node) {
	i := b - 'a'
	if n.nodes[i] == nil {
		n.nodes[i] = &Q0140Node{}
	}
	return n.nodes[i]
}

type Q0140Tracer struct {
	wordsIdx []int
	node     *Q0140Node
}

func (t Q0140Tracer) Fork(next *Q0140Node) *Q0140Tracer {
	if !t.node.IsEnd {
		return nil
	}

	newT := &Q0140Tracer{
		wordsIdx: make([]int, len(t.wordsIdx)+1),
		node:     next}
	copy(newT.wordsIdx, t.wordsIdx)
	newT.wordsIdx[len(t.wordsIdx)] = t.node.wIdx
	return newT
}

func (t Q0140Tracer) String() string {
	return fmt.Sprintf("{%d}", t.wordsIdx)
}

func (t Q0140Tracer) ToString(wordDict []string, sep string) string {
	ss := make([]string, len(t.wordsIdx))
	for j := range ss {
		ss[j] = wordDict[t.wordsIdx[j]]
	}
	return strings.Join(ss, sep)
}

func (q Q0140) WordBreak(s string, wordDict []string) (words []string) {
	root := &Q0140Node{} // root of a prefix tree
	for i, w := range wordDict {
		curr := root
		for j := 0; j < len(w); j++ {
			curr = curr.Append(w[j])
		}
		curr.wIdx = i
		curr.IsEnd = true
	}
	// fmt.Println(root)

	prev := []*Q0140Tracer{{wordsIdx: nil, node: root}}
	for i := range s {
		// fmt.Printf("%c(%02d) %v\n", s[i], i, prev)
		curr := []*Q0140Tracer{}
		for _, t := range prev {
			if t.node = t.node.Next(s[i]); t.node != nil {
				// fmt.Printf("\t- Append %c\n", s[i])
				if i != len(s)-1 {
					// do not append to prev when the last element is selected
					curr = append(curr, t)
				}

				if t.node.IsEnd {
					// fmt.Printf("\t- Add Word %s\n", wordDict[t.node.i])
					curr = append(curr, t.Fork(root))
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
		words[i] = prev[i].ToString(wordDict, " ")
	}
	return words
}
