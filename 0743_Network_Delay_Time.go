package leetcode

type Q0743 struct{}

func (q Q0743) NetworkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]int, n+1)
	for i, e := range times {
		graph[e[0]] = append(graph[e[0]], i)
	}

	// BFS
	distFromOri := make([]int, len(graph))
	for i := range distFromOri {
		distFromOri[i] = -1
	}
	distFromOri[k] = 0

	prev := graph[k]
	for len(prev) > 0 {
		curr := []int{}
		for _, i := range prev {
			next := times[i]
			if distFromOri[next[1]] < 0 || distFromOri[next[1]] > distFromOri[next[0]]+next[2] {
				distFromOri[next[1]] = distFromOri[next[0]] + next[2]
				curr = append(curr, graph[next[1]]...)
			}
		}
		prev = curr
	}

	maxDistFromOri := 0
	for i := 1; i < len(distFromOri); i++ {
		if distFromOri[i] < 0 {
			return -1
		}

		if maxDistFromOri < distFromOri[i] {
			maxDistFromOri = distFromOri[i]
		}
	}
	return maxDistFromOri
}
