package leetcode

type Q0399 struct{}

type Q0399Edge struct {
	To  int
	Val float64
}

type Q0399Graph map[int][]Q0399Edge

func (g Q0399Graph) DFS(i, j int) float64 {
	var dfs func(i, j int, v float64, ans *float64, stop *bool)
	var hasVisited []bool

	dfs = func(i, j int, v float64, ans *float64, stop *bool) {
		edges := g[i]
		for k := 0; k < len(edges) && !(*stop); k++ {
			if edges[k].To == j {
				// fmt.Printf("\t %d -> %d\n", i, j)
				(*ans) = v * edges[k].Val
				(*stop) = true
			} else {
				if !hasVisited[edges[k].To] {
					// fmt.Printf("\t %d -> %d\n", i, edges[k].To)
					hasVisited[edges[k].To] = true
					dfs(edges[k].To, j, v*edges[k].Val, ans, stop)
				}
			}
		}
	}

	ans, stop := -1.0, false
	hasVisited = make([]bool, len(g))
	hasVisited[i] = true
	dfs(i, j, 1, &ans, &stop)
	return ans
}

func (g Q0399Graph) AppendEdge(from, to int, val float64) {
	g[from] = append(g[from], Q0399Edge{to, val})
}

func (q Q0399) CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := Q0399Graph{}
	m := map[string]int{}

	k := 0
	for l, eq := range equations {
		var i, j int
		var ok bool
		if i, ok = m[eq[0]]; !ok {
			i, m[eq[0]], k = k, k, k+1
			graph.AppendEdge(i, i, 1.0)
		}

		if j, ok = m[eq[1]]; !ok {
			j, m[eq[1]], k = k, k, k+1
			if values[l] != 0.0 {
				graph.AppendEdge(j, j, 1.0)
			}
		}

		graph.AppendEdge(i, j, values[l])
		if values[l] != 0.0 {
			graph.AppendEdge(j, i, 1/values[l])
		}
	}

	ans := make([]float64, len(queries))
	for k, q := range queries {
		// fmt.Println("Qurey:", q)
		var i, j int
		var ok bool
		ans[k] = -1
		if i, ok = m[q[0]]; !ok {
			continue
		}

		if j, ok = m[q[1]]; !ok {
			continue
		}
		ans[k] = graph.DFS(i, j)
	}

	return ans
}
