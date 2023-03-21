package leetcode

import (
	"fmt"
	"sort"
)

type Q0347Count struct {
	n     int
	count int
}

type Q0347 struct{}

func (c1 *Q0347Count) Less(c2 *Q0347Count) bool {
	return c1.count < c2.count
}

func (c Q0347Count) String() string {
	return fmt.Sprintf("%2d: %3d", c.n, c.count)
}

type SortQ0347ByCount []*Q0347Count

func (a SortQ0347ByCount) Len() int           { return len(a) }
func (a SortQ0347ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortQ0347ByCount) Less(i, j int) bool { return a[i].Less(a[j]) }

func (q Q0347) TopKFrequent(nums []int, k int) []int {
	c := make([]*Q0347Count, 0, len(nums))
	m := make(map[int]*Q0347Count)

	for _, n := range nums {
		_, ok := m[n]
		if !ok {
			e := &Q0347Count{n: n, count: 0}
			c = append(c, e)
			m[n] = e
		}
		(*m[n]).count += 1
	}

	sort.Sort(sort.Reverse(SortQ0347ByCount(c)))
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
