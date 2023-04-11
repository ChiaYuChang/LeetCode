package leetcode

import (
	"container/heap"
	"fmt"
)

type Q0264 struct{}

func (q Q0264) SlowNthUglyNumberII(n int) int {
	h := IntHeap([]int{1})
	heap.Init(&h)

	var ce, pe int
	for n > 0 {
		for ce == pe {
			ce = heap.Pop(&h).(int)
		}
		fmt.Println(ce)

		for _, i := range []int{2, 3, 5} {
			heap.Push(&h, i*ce)
		}
		pe = ce
		n--
	}
	return ce
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
