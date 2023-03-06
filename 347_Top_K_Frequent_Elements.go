package leetcode

import (
	"fmt"
	"sort"
)

type Count struct {
	n     int
	count int
}

func (c1 *Count) Less(c2 *Count) bool {
	return c1.count < c2.count
}

func (c Count) String() string {
	return fmt.Sprintf("%2d: %3d", c.n, c.count)
}

type SortByCount []*Count

func (a SortByCount) Len() int           { return len(a) }
func (a SortByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByCount) Less(i, j int) bool { return a[i].Less(a[j]) }

func TopKFrequent(nums []int, k int) []int {
	c := make([]*Count, 0, len(nums))
	m := make(map[int]*Count)

	for _, n := range nums {
		_, ok := m[n]
		if !ok {
			e := &Count{n: n, count: 0}
			c = append(c, e)
			m[n] = e
		}
		(*m[n]).count += 1
	}

	sort.Sort(sort.Reverse(SortByCount(c)))
	for i := range c {
		fmt.Println(c[i])
	}

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = c[i].n
	}

	min := c[k-1].count
	for i := k; i < len(c); i++ {
		if c[i].count == min {
			ans = append(ans, c[i].n)
		}
	}

	return ans
}
