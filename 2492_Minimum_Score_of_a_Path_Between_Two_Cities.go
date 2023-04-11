package leetcode

import (
	"fmt"
	"math"
	"strings"
)

type Q2492 struct{}

func (q Q2492) MinScore(n int, roads [][]int) int {
	cc := q.NewCC(n)
	for _, e := range roads {
		cc.Connect(e[0], e[1], e[2])
	}
	return cc.MinScore(1, n)
}

type Q2492CC struct {
	node []int
	mScr []int
}

func (q Q2492) NewCC(n int) *Q2492CC {
	cc := &Q2492CC{
		node: make([]int, n),
		mScr: make([]int, n),
	}
	for i := 0; i < n; i++ {
		cc.node[i] = i
		// cc.mScr[i] = -1
		cc.mScr[i] = math.MaxInt
	}
	return cc
}

func (cc *Q2492CC) Root(i int) (int, int) {
	nhop := 0
	for cc.node[i] != i {
		i = cc.node[i]
		nhop++
	}
	return i, nhop
}

func (cc *Q2492CC) Connect(i, j, dist int) {
	i, ihop := cc.Root(i - 1)
	j, jhop := cc.Root(j - 1)
	dist = cc.Min(dist, cc.mScr[i])
	dist = cc.Min(dist, cc.mScr[j])

	if ihop > jhop {
		cc.node[j] = i
		cc.mScr[i] = cc.Min(cc.mScr[i], dist)
	} else {
		cc.node[i] = j
		cc.mScr[j] = cc.Min(cc.mScr[j], dist)
	}
}

func (cc *Q2492CC) Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (cc *Q2492CC) MinScore(i, j int) int {
	i, _ = cc.Root(i - 1)
	j, _ = cc.Root(j - 1)

	if i != j {
		return -1
	}

	return cc.mScr[i]
}

func (cc Q2492CC) String() string {
	sb := strings.Builder{}
	sb.WriteString("Connected Component:\n")
	for i, c := range cc.node {
		sb.WriteString(fmt.Sprintf("%2d: %2d (%d)\n", i+1, c+1, cc.mScr[c]))
	}
	return sb.String()
}
