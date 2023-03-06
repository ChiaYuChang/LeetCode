package leetcode

import (
	"container/heap"
	"fmt"
	"sort"
)

const zero = 'A'

func (x *nchar) Less(y *nchar) bool {
	if x.count != y.count {
		return x.count > y.count
	}
	return x.char < y.char
}

func (c nchar) String() string {
	return fmt.Sprintf("%c: %d", c.char, c.count)
}

type nCharHeap []*nchar

func (h nCharHeap) Len() int           { return len(h) }
func (h nCharHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h nCharHeap) Less(i, j int) bool { return h[i].Less(h[j]) }

func (h *nCharHeap) Push(x any) {
	(*h) = append((*h), x.(*nchar))
}

func (h *nCharHeap) Pop() any {
	if h.Len() == 0 {
		return nil
	}
	old := (*h)
	n := len(old)
	x := old[n-1]
	(*h) = old[:n-1]
	return x
}

func LeastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	cs := nCharHeap(make([]*nchar, 26))
	for i := range cs {
		cs[i] = &nchar{char: byte(zero + i), count: 0}
	}

	for _, c := range tasks {
		cs[c-zero].count += 1
	}

	sort.Sort(cs)
	notZero := 0
	for i := range cs {
		// remove all zero nodes
		if cs[i].count == 0 {
			notZero = i
			break
		}
	}
	cs = cs[:notZero]
	heap.Init(&cs)

	var c1tn []*nchar
	nSeg := 0
	lSeg := n + 1
	for cs.Len() > 0 {
		c1tn = make([]*nchar, 0, n)
		for i := 0; i < lSeg && cs.Len() > 0; i++ {
			c := heap.Pop(&cs).(*nchar)
			c1tn = append(c1tn, c)
		}

		if cs.Len() < 1 {
			break
		}

		l := len(c1tn)
		nSeg += c1tn[l-1].count
		for i := 0; i < l-1; i++ {
			c1tn[i].count -= c1tn[l-1].count
			if c1tn[i].count > 0 {
				heap.Push(&cs, c1tn[i])
			}
		}
	}

	nSeg += c1tn[0].count - 1
	ans := nSeg * lSeg
	for _, c := range c1tn {
		if c.count == c1tn[0].count {
			ans++
		}
	}

	return ans
}
