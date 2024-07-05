package leetcode

type Q2039 struct{}

func (q Q2039) NetworkBecomesIdle(edges [][]int, patience []int) int {
	graph := make([][]int, len(patience))
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	ori := 0
	pLen := q.BSF(graph, ori) // path length

	idle := 0
	for i := 0; i < len(pLen); i++ {
		if i == ori {
			continue
		}
		pLen[i] *= 2
		r := pLen[i] % patience[i]
		q := pLen[i] / patience[i]

		if q != 0 {
			// q == 0 -> only sent message once
			if r == 0 {
				pLen[i] = 2*pLen[i] - patience[i]
			} else {
				pLen[i] = 2*pLen[i] - r
			}
		}

		if idle < pLen[i] {
			idle = pLen[i]
		}
	}
	return idle + 1
}

func (q Q2039) BSF(graph [][]int, ori int) []int {
	distFromOri := make([]int, len(graph))
	for i := range distFromOri {
		distFromOri[i] = -1
	}
	distFromOri[ori] = 0

	prev := graph[ori]
	step := 1
	for len(prev) > 0 {
		curr := []int{}
		for _, next := range prev {
			if distFromOri[next] >= 0 {
				// shortest path
				continue
			}
			distFromOri[next] = step
			curr = append(curr, graph[next]...)
		}
		prev = curr
		step++
	}
	return distFromOri
}
