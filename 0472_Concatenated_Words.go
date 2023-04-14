package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

type Q0472 struct{}

const (
	ROOT rune = 'z' + 1
	ZERO rune = 'a'
)

type Q0472Node struct {
	char  rune
	child [27]*Q0472Node
}

func (n *Q0472Node) IsRoot() bool {
	return n.char == ROOT
}

func (n *Q0472Node) IsEnd() bool {
	return n.Child(ROOT) != nil
}

func (n *Q0472Node) Child(r rune) *Q0472Node {
	return n.child[r-ZERO]
}

func (n *Q0472Node) Append(r rune) (next *Q0472Node, ok bool) {
	i := r - ZERO
	if n.child[i] == nil {
		ok = true
		n.child[i] = &Q0472Node{r, [27]*Q0472Node{}}
	}
	return n.child[i], ok
}

func (n Q0472Node) String() string {
	var ToString func(n *Q0472Node, char byte, lvl int, sb *strings.Builder)

	ToString = func(n *Q0472Node, char byte, lvl int, sb *strings.Builder) {
		sb.WriteString(strings.Repeat(" ", lvl))
		sb.WriteByte(char)
		sb.WriteRune('\n')

		if n.child[ROOT-ZERO] != nil {
			sb.WriteString(strings.Repeat(" ", lvl+1))
			sb.WriteString("end\n")
		}

		for i := 0; i < 26; i++ {
			if n.child[i] != nil {
				ToString(n.child[i], byte('a'+i), lvl+1, sb)
			}
		}
	}

	sb := strings.Builder{}
	sb.WriteString("root\n")
	for i := 0; i < 26; i++ {
		if next := n.child[i]; next != nil {
			ToString(next, byte('a'+i), 4, &sb)
		}
	}
	return sb.String()
}

type Q0472Visiter struct {
	n          int
	index      int // nth rune
	*Q0472Node     // current node
}

func (v Q0472Visiter) Fork() *Q0472Visiter {
	if v.IsEnd() {
		return &Q0472Visiter{v.n + 1, v.index, v.child[ROOT-ZERO]}
	}
	panic("not the end of a word")
}

func (v Q0472Visiter) String() string {
	return fmt.Sprintf("%d: %c(%d)", v.n, v.Q0472Node.char, v.index)
}

func (n *Q0472Node) AppendString(s string) (nNewNodes int) {
	appendAt := &Q0472Visiter{n: 0, index: -1, Q0472Node: n}
	reachEnd := true

	if n.IsRoot() {
		prev := []*Q0472Visiter{appendAt}
		for i, r := range s {
			// fmt.Printf("\t%c(%d)\n", r, i)
			curr := []*Q0472Visiter{}
			for _, v := range prev {
				if next := v.Child(r); next != nil {
					// fmt.Printf("\t - Append: %c\n", next.char)

					v.Q0472Node, v.index = next, i
					curr = append(curr, v)

					if v.IsEnd() {
						// fmt.Println("\t - Fork")
						curr = append(curr, v.Fork())
					}
				}
			}
			if len(curr) < 1 {
				reachEnd = false
				break
			}
			prev = curr
			// fmt.Println("\t - q:", prev)
		}
		// fmt.Println("ReachEnd:", reachEnd)
		if reachEnd {
			reachEnd = false
			for _, n := range prev {
				reachEnd = reachEnd || n.IsEnd()
			}
		}

		if reachEnd {
			return 0
		}
	} else {
		panic("the given node is not a root node")
	}

	appendAt.index++
	curr := appendAt.Q0472Node
	// fmt.Println("\t- Add word freqement:", s[appendAt.index:])
	for _, r := range s[appendAt.index:] {
		curr, _ = curr.Append(r)
		nNewNodes++
	}
	curr.child[ROOT-ZERO] = n
	return nNewNodes
}

func (q Q0472) FindAllConcatenatedWordsInADictsPrefixTree(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) != len(words[j]) {
			return len(words[i]) < len(words[j])
		}
		return words[i] < words[j]
	})

	cWordsInDict := []string{}
	root := &Q0472Node{char: ROOT}
	for _, w := range words {
		if w == "" {
			// cWordsInDict = append(cWordsInDict, w)
			continue
		}
		// fmt.Println("Add word:", w)
		if n := root.AppendString(w); n == 0 {
			// fmt.Println("-> Append word to list:", w)
			cWordsInDict = append(cWordsInDict, w)
		}
		// fmt.Println(root)
	}

	return cWordsInDict
}

// FROM https://github.com/chai2010/LeetCode-in-Go/
func (q Q0472) FindAllConcatenatedWordsInADictsDP(words []string) []string {
	hasLen := make(map[int]bool, len(words))
	hasWord := make(map[string]bool, len(words))
	for _, word := range words {
		hasLen[len(word)] = true
		hasWord[word] = true
	}

	res := make([]string, 0, len(words))
	for _, word := range words {
		if q.isConcatenated(word, hasLen, hasWord, false) {
			res = append(res, word)
		}
	}

	return res
}

func (q Q0472) isConcatenated(word string, hasLen map[int]bool, hasWord map[string]bool, isCutted bool) bool {
	for wLen := 1; wLen < len(word); wLen++ {
		if hasLen[wLen] &&
			hasWord[word[:wLen]] &&
			q.isConcatenated(word[wLen:], hasLen, hasWord, true) {
			return true
		}
	}
	return isCutted &&
		hasLen[len(word)] &&
		hasWord[word]
}
