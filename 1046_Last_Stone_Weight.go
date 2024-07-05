package leetcode

import (
	"container/heap"
)

type Q1046 struct{}

type Q1046IntHeap []int

func (h Q1046IntHeap) Len() int           { return len(h) }
func (h Q1046IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h Q1046IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Q1046IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Q1046IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (q Q1046) LastStoneWeight(stones []int) int {
	h := Q1046IntHeap(stones)
	heap.Init(&h)
	for len(h) > 1 {
		n1 := heap.Pop(&h).(int)
		n2 := heap.Pop(&h).(int)
		if n1 > n2 {
			heap.Push(&h, n1-n2)
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(&h).(int)
}
