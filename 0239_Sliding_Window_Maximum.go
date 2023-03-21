package leetcode

import (
	"container/heap"
)

type Q0239 struct{}

// value, position
type Q0239Pair [2]int

type Q0239Heap []*Q0239Pair

func (q Q0239) NewHeap() *Q0239Heap {
	h := Q0239Heap(make([]*Q0239Pair, 0, 100))
	heap.Init(&h)
	return &h
}

func (h Q0239Heap) Len() int {
	return len(h)
}

func (h Q0239Heap) Less(i, j int) bool {
	if h[i][0] != h[j][0] {
		return h[i][0] > h[j][0]
	}
	return h[i][1] > h[j][1]
}
func (h Q0239Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Q0239Heap) Push(x any) {
	*h = append(*h, x.(*Q0239Pair))
}

func (h *Q0239Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (q Q0239) MaxSlidingWindow(nums []int, k int) []int {
	hp := q.NewHeap()
	for i := 0; i < k-1; i++ {
		heap.Push(hp, &Q0239Pair{nums[i], i})
	}

	localmax := make([]int, 0, len(nums)-k+1)
	for i := 0; i < len(nums)-k+1; i++ {
		heap.Push(hp, &Q0239Pair{nums[i+k-1], i + k - 1})
		// position of max should be greater than start point
		max := &Q0239Pair{0, i - 1}
		for (*max)[1] < i {
			max = heap.Pop(hp).(*Q0239Pair)
		}
		heap.Push(hp, max)
		localmax = append(localmax, (*max)[0])
	}

	return localmax
}
