package leetcode

type Q1319 struct{}

func (q Q1319) MakeConnected(n int, connections [][]int) int {
	cc := q.NewCC(n)
	for _, e := range connections {
		cc.Connect(e[0], e[1])
	}

	roots := cc.Roots()
	nE := cc.NE(roots)
	nV := cc.NV(roots)

	nRedundant := 0
	for i := 0; i < len(roots); i++ {
		nRedundant += nE[i] - nV[i] + 1
	}

	// at least len(roots)-1 is needed to connect len(roots) cc
	if nRedundant-(len(roots)-1) < 0 {
		return -1
	}
	return len(roots) - 1
}

// connected component
type Q1319CC struct {
	node []int // labels
	nE   []int // number of edges in each cc
	nV   []int // number of vertices in each cc
}

// connected compenent constructor
func (q Q1319) NewCC(n int) *Q1319CC {
	cc := &Q1319CC{
		node: make([]int, n),
		nE:   make([]int, n),
		nV:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		cc.node[i] = i // if cc.node[i] == i -> root node
		cc.nV[i] = 1
	}
	return cc
}

// return the root node of the given node
func (cc *Q1319CC) Root(i int) (int, int) {
	nhop := 0
	for cc.node[i] != i {
		i = cc.node[i]
		nhop++
	}
	return i, nhop
}

// connect node i and j
func (cc *Q1319CC) Connect(i, j int) {
	i, ihop := cc.Root(i)
	j, jhop := cc.Root(j)

	if i != j {
		// merging cc
		if ihop > jhop {
			cc.node[j] = i
			cc.nE[i] += cc.nE[j] + 1
			cc.nV[i] += cc.nV[j]
		} else {
			cc.node[i] = j
			cc.nE[j] += cc.nE[i] + 1
			cc.nV[j] += cc.nV[i]
		}
	} else {
		// if node i and node j are already in the same cc
		cc.nE[i] += 1
	}
}

// return root nodes of each cc
func (cc *Q1319CC) Roots() []int {
	r := make([]int, 0)
	for i := 0; i < len(cc.node); i++ {
		if i == cc.node[i] {
			r = append(r, i)
		}
	}
	return r
}

// return the number of vertices in a cc
func (c *Q1319CC) NV(roots []int) []int {
	n := make([]int, len(roots))
	for i, r := range roots {
		if c.node[r] == r {
			n[i] = c.nV[r]
		}
	}
	return n
}

// return the number of edges in a cc
func (c *Q1319CC) NE(roots []int) []int {
	n := make([]int, len(roots))
	for i, r := range roots {
		if c.node[r] == r {
			n[i] = c.nE[r]
		}
	}
	return n
}
