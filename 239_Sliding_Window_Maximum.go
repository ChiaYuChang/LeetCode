package leetcode

import (
	"container/heap"
)

// value, position
type Pair [2]int

type Heap []*Pair

func NewHeap() *Heap {
	h := Heap(make([]*Pair, 0, 100))
	heap.Init(&h)
	return &h
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i][0] != h[j][0] {
		return h[i][0] > h[j][0]
	}
	return h[i][1] > h[j][1]
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(*Pair))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MaxSlidingWindow(nums []int, k int) []int {
	hp := NewHeap()
	for i := 0; i < k-1; i++ {
		heap.Push(hp, &Pair{nums[i], i})
	}

	localmax := make([]int, 0, len(nums)-k+1)
	for i := 0; i < len(nums)-k+1; i++ {
		heap.Push(hp, &Pair{nums[i+k-1], i + k - 1})
		// position of max should be greater than start point
		max := &Pair{0, i - 1}
		for (*max)[1] < i {
			max = heap.Pop(hp).(*Pair)
		}
		heap.Push(hp, max)
		localmax = append(localmax, (*max)[0])
	}

	return localmax
}
