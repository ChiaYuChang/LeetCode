package leetcode

type Q0990 struct{}

// connected component with 26 nodes
type Q0990CC26 struct {
	nodes [26]int
}

// new connected component with 26 nodes
func (q Q0990) NewCC26() *Q0990CC26 {
	cc := &Q0990CC26{nodes: [26]int{}}
	for i := 0; i < 26; i++ {
		cc.nodes[i] = i
	}
	return cc
}

// x and y should be with in a-z
func (cc *Q0990CC26) Connect(x, y byte) {
	i, j := int(x-'a'), int(y-'a')

	ihop := 0
	for cc.nodes[i] != i {
		i = cc.nodes[i]
		ihop++
	}

	jhop := 0
	for cc.nodes[j] != j {
		j = cc.nodes[j]
		jhop++
	}

	if ihop > jhop {
		cc.nodes[j] = i
	} else {
		cc.nodes[i] = j
	}
}

// x and y should be with in a-z
func (cc *Q0990CC26) IsConnected(x, y byte) bool {
	i, j := int(x-'a'), int(y-'a')

	for cc.nodes[i] != i {
		i = cc.nodes[i]
	}

	for cc.nodes[j] != j {
		j = cc.nodes[j]
	}
	return i == j
}

func (q Q0990) EquationsPossible(equations []string) bool {
	neqs := [][2]byte{}
	cc := q.NewCC26()

	for _, eq := range equations {
		v1, v2 := eq[0], eq[3]
		switch eq[1:3] {
		case "==":
			cc.Connect(v1, v2)
		case "!=":
			neqs = append(neqs, [2]byte{v1, v2})
		}
	}

	for _, neq := range neqs {
		if cc.IsConnected(neq[0], neq[1]) {
			return false
		}
	}
	return true
}
