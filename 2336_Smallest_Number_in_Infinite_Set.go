package leetcode

import (
	"fmt"
	"sort"
)

type Q2336 struct{}

type Q2336MinHeap struct {
	value []int
}

func (q Q2336) NewHeap(array []int) *Q2336MinHeap {
	v := make([]int, len(array)+1)
	copy(v[1:], array)
	if !sort.IsSorted(sort.IntSlice(v[1:])) {
		sort.Sort(sort.IntSlice(v[1:]))
	}
	return &Q2336MinHeap{v}
}

func (h Q2336MinHeap) String() string {
	return fmt.Sprintf("Heap: %v", h.value)
}

func (h Q2336MinHeap) More(i, j int) bool {
	return h.value[i] > h.value[j]
}

func (h Q2336MinHeap) Swap(i, j int) {
	h.value[i], h.value[j] = h.value[j], h.value[i]
}

func (h Q2336MinHeap) swim(k int) {
	for ; k > 1 && h.More(k/2, k); k /= 2 {
		h.Swap(k/2, k)
	}
}

func (h Q2336MinHeap) Len() int {
	return len(h.value) - 1
}

func (h *Q2336MinHeap) Push(n int) {
	h.value = append(h.value, n)
	h.swim(h.Len())
}

func (h Q2336MinHeap) sink(k int) {
	for k*2 <= h.Len() {
		j := k * 2
		if (j+1 < len(h.value)) && h.More(j, j+1) {
			j++
		}
		if !h.More(k, j) {
			break
		}
		h.Swap(k, j)
		k = j
	}
}

func (h *Q2336MinHeap) Peek() int {
	return h.value[1]
}

func (h *Q2336MinHeap) Pop() (int, bool) {
	if h.Len() < 1 {
		return 0, false
	}

	n := h.value[1]
	h.Swap(1, h.Len())
	h.value = h.value[:h.Len()]
	h.sink(1)
	return n, true
}

type Q2336SmallestInfiniteSet struct {
	set []bool
	h   *Q2336MinHeap
}

func (q Q2336) Constructor() Q2336SmallestInfiniteSet {
	set := make([]bool, 1000+1)
	val := make([]int, 1000+1)
	for i := 1; i <= 1000; i++ {
		set[i] = true
		val[i] = i
	}
	return Q2336SmallestInfiniteSet{
		set: set,
		h:   &Q2336MinHeap{value: val},
	}
}

func (this *Q2336SmallestInfiniteSet) PopSmallest() int {
	min, ok := this.h.Pop()
	if ok {
		this.set[min] = false
		return min
	}
	return 0
}

func (this *Q2336SmallestInfiniteSet) AddBack(num int) {
	if !this.set[num] {
		this.h.Push(num)
		this.set[num] = true
	}
}
