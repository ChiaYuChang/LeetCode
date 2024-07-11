package leetcode

import (
	"sort"
	"strings"
)

type Q2192 struct{}

func (q Q2192) GetAncestors(n int, edges [][]int) [][]int {
	// typological sort
	type Node struct {
		index     int
		inDegree  int
		offspring []int
		ancestor  map[int]bool
	}

	topologicalOrder := make([]int, 0, n)
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{i, 0, []int{}, map[int]bool{}}
	}

	for _, edge := range edges {
		nodes[edge[1]].inDegree += 1
		nodes[edge[0]].offspring = append(nodes[edge[0]].offspring, edge[1])
	}

	for _, node := range nodes {
		if node.inDegree == 0 {
			topologicalOrder = append(topologicalOrder, node.index)
		}
	}

	for i := 0; i < len(topologicalOrder); i++ {
		for _, to := range nodes[topologicalOrder[i]].offspring {
			nodes[to].inDegree -= 1
			if nodes[to].inDegree == 0 {
				topologicalOrder = append(topologicalOrder, nodes[to].index)
			}
		}
	}

	for _, i := range topologicalOrder {
		for _, to := range nodes[i].offspring {
			nodes[to].ancestor[nodes[i].index] = true
			for j := range nodes[i].ancestor {
				nodes[to].ancestor[j] = true
			}
		}
	}

	ancestor := make([][]int, n)
	for i, node := range nodes {
		ancestor[i] = make([]int, 0, len(node.ancestor))
		for a := range node.ancestor {
			ancestor[i] = append(ancestor[i], a)
		}
	}

	for i := range ancestor {
		sort.Ints(ancestor[i])
	}
	return ancestor
}

func (q Q2192) GetAncestorsSlow(n int, edges [][]int) [][]int {
	a := q.NewBinaryMatrix(n)
	for i := 0; i < a.Size(); i++ {
		a.Set(i, i, true)
	}

	for _, e := range edges {
		a.Set(e[0], e[1], true)
	}

	c := a.ToConverge()
	ancestor := make([][]int, n)
	for row := range c {
		for col := range c[row] {
			if c[col][row] && row != col {
				ancestor[row] = append(ancestor[row], col)
			}
		}
	}
	return ancestor
}

func (q Q2192) NewBinaryMatrix(n int) Q2192BinarySquareMatrix {
	m := make([][]bool, n)
	for i := 0; i < n; i++ {
		m[i] = make([]bool, n)
	}
	return m
}

type Q2192BinarySquareMatrix [][]bool

func (m *Q2192BinarySquareMatrix) Set(row, col int, val bool) *Q2192BinarySquareMatrix {
	(*m)[row][col] = val
	return m
}

func (m Q2192BinarySquareMatrix) Size() int {
	return len(m)
}

func (m Q2192BinarySquareMatrix) String() string {
	sb := strings.Builder{}
	for row := 0; row < m.Size(); row++ {
		sb.WriteRune('[')
		for col := 0; col < m.Size(); col++ {
			if m[row][col] {
				sb.WriteRune('1')
			} else {
				sb.WriteRune('0')
			}

			if col < m.Size()-1 {
				sb.WriteRune(' ')
			} else {
				sb.WriteRune(']')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (ml Q2192BinarySquareMatrix) Equal(mr Q2192BinarySquareMatrix) bool {
	if ml.Size() != mr.Size() {
		return false
	}

	n := ml.Size()
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if ml[row][col] != mr[row][col] {
				return false
			}
		}
	}
	return true
}

func (ml Q2192BinarySquareMatrix) Mut(mr Q2192BinarySquareMatrix) Q2192BinarySquareMatrix {
	n := ml.Size()
	m := make([][]bool, n)
	for row := 0; row < n; row++ {
		m[row] = make([]bool, n)
		for col := 0; col < n; col++ {
			for i := 0; i < n && !m[row][col]; i++ {
				m[row][col] = m[row][col] || (ml[row][i] && mr[i][col])
			}
		}
	}
	return m
}

func (m Q2192BinarySquareMatrix) ToConverge() Q2192BinarySquareMatrix {
	curr := m
	for {
		next := curr.Mut(curr)
		if curr.Equal(next) {
			break
		}
		curr = next
	}
	return curr
}
