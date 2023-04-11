package leetcode

type Q0131 struct{}

type Q0131Graph struct {
	edge [][]int
	s    string
}

func (g Q0131Graph) FindAllPalindromeCombination() [][]string {
	combination := [][]string{}
	g.dfs(0, []string{}, &combination)
	return combination
}

func (g Q0131Graph) dfs(i int, pre []string, ans *[][]string) {
	if i >= len(g.s) {
		(*ans) = append((*ans), pre)
		return
	}
	for _, j := range g.edge[i] {
		s := make([]string, len(pre), len(pre)+1)
		copy(s, pre)
		s = append(s, g.s[i:j])
		g.dfs(j, s, ans)
	}
}

func (q Q0131) Partition(s string) [][]string {
	if len(s) == 0 {
		return nil
	}

	interval := make([][]int, len(s))
	// Dynamic Programming
	prev := make([]int, len(s)+1)
	last := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		curr := make([]int, last+1)
		// margin
		if s[i] == s[0] {
			curr[0] = 1
		}
		for j := 1; j <= last && j < len(s); j++ {
			if s[i] == s[j] && j > 0 {
				curr[j] = prev[j-1] + 1
			}
		}

		if curr[last] > 0 {
			for d := curr[last]; d > 1; d-- {
				interval[last-d+1] = append(interval[last-d+1], last+d-1)
			}
		}

		if curr[last-1] > 0 {
			for d := curr[last-1]; d > 0; d-- {
				interval[last-d] = append(interval[last-d], last+d-1)
			}
		}
		prev = curr
		last--
	}
	return Q0131Graph{edge: interval, s: s}.FindAllPalindromeCombination()
}
