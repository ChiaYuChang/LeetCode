package leetcode

type Q2316 struct{}

func (q Q2316) CountPairs(n int, connections [][]int) int {
	cc := NewQ2316CC(n)
	for _, e := range connections {
		cc.Connect(e[0], e[1])
	}

	roots := cc.Roots()
	nV := cc.NV(roots)

	nPairs := 0
	for i := 0; i < len(nV)-1; i++ {
		for j := i + 1; j < len(nV); j++ {
			nPairs += nV[i] * nV[j]
		}
	}
	return nPairs
}

// connected component
type Q2316CC struct {
	node []int // labels
	nV   []int // number of vertices in each cc
}

// connected compenent constructor
func NewQ2316CC(n int) *Q2316CC {
	cc := &Q2316CC{
		node: make([]int, n),
		nV:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		cc.node[i] = i // if cc.node[i] == i -> root node
		cc.nV[i] = 1
	}
	return cc
}

// return the root node of the given node
func (cc *Q2316CC) Root(i int) (int, int) {
	nhop := 0
	for cc.node[i] != i {
		i = cc.node[i]
		nhop++
	}
	return i, nhop
}

// connect node i and j
func (cc *Q2316CC) Connect(i, j int) {
	i, ihop := cc.Root(i)
	j, jhop := cc.Root(j)

	if i != j {
		// merging cc
		if ihop > jhop {
			cc.node[j] = i
			cc.nV[i] += cc.nV[j]
		} else {
			cc.node[i] = j
			cc.nV[j] += cc.nV[i]
		}
	}
}

// return root nodes of each cc
func (cc *Q2316CC) Roots() []int {
	r := make([]int, 0)
	for i := 0; i < len(cc.node); i++ {
		if i == cc.node[i] {
			r = append(r, i)
		}
	}
	return r
}

// return the number of vertices in a cc
func (c *Q2316CC) NV(roots []int) []int {
	n := make([]int, len(roots))
	for i, r := range roots {
		if c.node[r] == r {
			n[i] = c.nV[r]
		}
	}
	return n
}
