package leetcode

import (
	"fmt"
	"strings"
)

type Q0839 struct{}

type Q0839CC struct {
	node []int
}

func NewCC(n int) *Q0839CC {
	cc := &Q0839CC{make([]int, n)}
	for i := 0; i < n; i++ {
		cc.node[i] = i
	}
	return cc
}

func (cc *Q0839CC) Root(i int) (int, int) {
	nhop := 0
	for cc.node[i] != i {
		i = cc.node[i]
		nhop++
	}
	return i, nhop
}

func (cc *Q0839CC) Connect(i, j int) {
	i, ihop := cc.Root(i)
	j, jhop := cc.Root(j)

	if ihop > jhop {
		cc.node[j] = i
	} else {
		cc.node[i] = j
	}
}

func (cc *Q0839CC) NumSubGraphs() int {
	nSubGraph := 0
	for root, node := range cc.node {
		if root == node {
			nSubGraph++
		}
	}
	return nSubGraph
}

func (cc Q0839CC) String() string {
	sb := strings.Builder{}
	sb.WriteString("Connected Component:\n")
	sb.WriteString("\t- ")
	for i, node := range cc.node {
		sb.WriteString(fmt.Sprintf("%2d", node))
		if i < len(cc.node)-1 {
			sb.WriteString(", ")
		} else {
			sb.WriteRune('\n')
		}
	}

	sb.WriteString("\t- ")
	for i := range cc.node {
		node, _ := cc.Root(i)
		sb.WriteString(fmt.Sprintf("%2d", node))
		if i < len(cc.node)-1 {
			sb.WriteString(", ")
		} else {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func (q Q0839) NumSimilarGroups(strs []string) int {
	cc := NewCC(len(strs))
	for i := 0; i < len(strs)-1; i++ {
		for j := i + 1; j < len(strs); j++ {
			if q.IsConnected(strs[i], strs[j]) {
				// fmt.Printf("Connect %s(%d) and %s(%d)\n", strs[i], i, strs[j], j)
				cc.Connect(i, j)
			}
		}
	}
	// fmt.Println(cc)
	return cc.NumSubGraphs()
}

func (q Q0839) IsConnected(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	if str1 == str2 {
		return true
	}

	mismatchAt := make([]int, 0, 2)
	for i := 0; i < len(str1) && len(mismatchAt) < 2; i++ {
		if str1[i] != str2[i] {
			mismatchAt = append(mismatchAt, i)
		}
	}

	// ...s1...B...s2...A(...s3...)
	// ...s1...A...s2...B(...s3...)
	if len(mismatchAt) == 2 &&
		str1[mismatchAt[0]] == str2[mismatchAt[1]] &&
		str1[mismatchAt[1]] == str2[mismatchAt[0]] &&
		str1[mismatchAt[1]+1:] == str2[mismatchAt[1]+1:] {
		return true
	}

	return false
}
