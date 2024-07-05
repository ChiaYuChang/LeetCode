package leetcode

import (
	"fmt"
	"sort"
)

type Q2285 bool

func (q Q2285) MaximumImportance(n int, roads [][]int) int64 {
	degree := make([]int, n)
	for _, edge := range roads {
		degree[edge[0]] += 1
		degree[edge[1]] += 1
	}
	sort.Sort(sort.IntSlice(degree))

	importance := 0
	for i := 0; i < n; i++ {
		fmt.Printf("Index %d, Degree: %d\n", i, degree[i])
		importance += (i + 1) * degree[i]
	}
	return int64(importance)
}
