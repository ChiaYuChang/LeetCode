package leetcode

type Q0113 struct{}

type Q0113Node struct {
	Val       int
	Neighbors []*Q0113Node
}

func (q Q0113) CloneGraph1(node *Q0113Node) *Q0113Node {
	if node == nil {
		return nil
	}

	root := &Q0113Node{Val: node.Val}
	dstQueue := []*Q0113Node{root}
	srcQueue := []*Q0113Node{node}
	mVal2Idx := map[int]int{root.Val: 0}

	for i := 0; i < len(dstQueue); i++ {
		dst := dstQueue[i]
		src := srcQueue[i]

		if src.Neighbors != nil {
			dst.Neighbors = make([]*Q0113Node, len(src.Neighbors))
			for j, nbr := range src.Neighbors {
				var nd *Q0113Node
				if indx, ok := mVal2Idx[nbr.Val]; ok {
					nd = dstQueue[indx]
				} else {
					nd = &Q0113Node{Val: nbr.Val}
					srcQueue = append(srcQueue, src.Neighbors[j])
					dstQueue = append(dstQueue, nd)
					mVal2Idx[nd.Val] = len(dstQueue) - 1
				}
				dst.Neighbors[j] = nd
			}
		}
	}
	return root
}

func (q Q0113) CloneGraph2(node *Q0113Node) *Q0113Node {
	if node == nil {
		return nil
	}

	adjLst := map[int][]int{}
	q.Travel(node, adjLst)

	newNodes := make([]*Q0113Node, len(adjLst))
	if len(newNodes) == 0 {
		return &Q0113Node{Val: node.Val}
	}

	for i := 1; i <= len(adjLst); i++ {
		newNodes[i-1] = &Q0113Node{Val: i}
	}

	for k, v := range adjLst {
		newNodes[k-1].Neighbors = make([]*Q0113Node, len(v))
		for i, j := range v {
			newNodes[k-1].Neighbors[i] = newNodes[j-1]
		}
	}

	return newNodes[0]
}

func (q Q0113) Travel(node *Q0113Node, m map[int][]int) {
	if node == nil {
		return
	}

	m[node.Val] = make([]int, len(node.Neighbors))
	for i, ng := range node.Neighbors {
		m[node.Val][i] = ng.Val
		if _, ok := m[ng.Val]; !ok {
			q.Travel(ng, m)
		}
	}
}
