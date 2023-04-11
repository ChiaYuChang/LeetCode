package leetcode

import (
	"fmt"
	"strings"
)

type Q0211 struct{}

type Q0211WordDictionary struct {
	root *Q0211Node
}

func (q Q0211) Constructor() Q0211WordDictionary {
	return Q0211WordDictionary{root: q.NewNode()}
}

func (this *Q0211WordDictionary) AddWord(word string) int {
	curr := this.root
	n := 0
	ok := false
	for i := 0; i < len(word); i++ {
		curr, ok = curr.Append(word[i])
		if ok {
			n++
		}
	}
	curr.IsEnd = true
	return n
}

func (this *Q0211WordDictionary) Search(word string) (ok bool) {
	bs := []byte(word)
	if bs[0] == '.' {
		for i := 0; i < 26; i++ {
			if child := this.root.Children[i]; child != nil {
				ok = ok || child.Search(bs[1:])
			}
		}
	} else {
		if child := this.root.Children[bs[0]-'a']; child != nil {
			ok = ok || child.Search(bs[1:])
		}
	}
	return ok
}

func (this Q0211WordDictionary) String() string {
	sb := strings.Builder{}
	sb.WriteString("WordDictionary:\n")
	for i := 0; i < 26; i++ {
		if child := this.root.Children[i]; child != nil {
			child.Visit(1, byte(i+'a'), &sb)
		}
	}
	return sb.String()
}

type Q0211Node struct {
	IsEnd    bool
	Children [26]*Q0211Node
}

func (q Q0211) NewNode() *Q0211Node {
	return &Q0211Node{Children: [26]*Q0211Node{}}
}

func (n *Q0211Node) IsLower(b byte) bool {
	return 'a' <= b && b <= 'z'
}

func (n *Q0211Node) Append(b byte) (child *Q0211Node, ok bool) {
	if !n.IsLower(b) {
		return nil, false
	}

	i := b - 'a'
	if n.Children[i] == nil {
		n.Children[i] = &Q0211Node{Children: [26]*Q0211Node{}}
		ok = true
	}
	return n.Children[i], ok
}

func (n *Q0211Node) Search(bs []byte) (ok bool) {
	if len(bs) == 0 {
		return n.IsEnd
	}

	if char := bs[0]; char == '.' {
		for i := 0; i < 26 && !ok; i++ {
			c := n.Children[i]
			if c == nil {
				continue
			}
			ok = ok || c.Search(bs[1:])
		}
	} else {
		if c := n.Children[char-'a']; c != nil {
			ok = c.Search(bs[1:])
		}
	}
	return ok
}

func (n *Q0211Node) Visit(nIndent int, char byte, sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%s%c (IsEnd: %v)\n", strings.Repeat("\t", nIndent), char, n.IsEnd))

	for i, child := range n.Children {
		if child != nil {
			child.Visit(nIndent+1, byte('a'+i), sb)
		}
	}
}
