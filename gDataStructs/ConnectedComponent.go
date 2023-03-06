package gdatastructs

type ConnectedComponent struct {
	nodes []int
}

func NewConnectedComponent(n int) *ConnectedComponent {
	cc := &ConnectedComponent{nodes: make([]int, n)}
	for i := range cc.nodes {
		cc.nodes[i] = i
	}
	return cc
}

func (cc *ConnectedComponent) Len() int {
	return len(cc.nodes)
}

func (cc *ConnectedComponent) Connect(x, y int) {
	var i, j, ihop, jhop int

	i, ihop = cc.Root(x)
	j, jhop = cc.Root(y)

	if ihop > jhop {
		cc.nodes[j] = i
	} else {
		cc.nodes[i] = j
	}
}

func (cc *ConnectedComponent) IsConnected(x, y int) bool {
	i, j := x, y

	for cc.nodes[i] != i {
		i = cc.nodes[i]
	}

	for cc.nodes[j] != j {
		j = cc.nodes[j]
	}
	return i == j
}

func (cc *ConnectedComponent) Root(x int) (i int, hop int) {
	i = x
	for cc.nodes[i] != i {
		i = cc.nodes[i]
		hop++
	}
	return i, hop
}

func (cc *ConnectedComponent) Groups() [][]int {
	mp := make(map[int][]int)

	for i, j := range cc.nodes {
		r, _ := cc.Root(j)
		mp[r] = append(mp[r], i)
	}

	group := make([][]int, 0, len(mp))
	for _, v := range mp {
		group = append(group, v)
	}
	return group
}
