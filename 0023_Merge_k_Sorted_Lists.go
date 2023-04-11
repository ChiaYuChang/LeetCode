package leetcode

import (
	"container/heap"
	"fmt"
	"strings"
)

type Q0023 struct{}

type Q0023ListNode struct {
	Val  int
	Next *Q0023ListNode
}

type Q0023ListNodeHeap []*Q0023ListNode

func (h Q0023ListNodeHeap) Len() int           { return len(h) }
func (h Q0023ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h Q0023ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// add a node to the heap
func (h *Q0023ListNodeHeap) Push(x any) {
	*h = append(*h, x.(*Q0023ListNode))
}

// return the nodes with min val
func (h *Q0023ListNodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// print ListNodeHeap
func (h Q0023ListNodeHeap) String() string {
	LN2Slc := func(n *Q0023ListNode) []int {
		slc := []int{}
		for c := n; c != nil; c = c.Next {
			slc = append(slc, c.Val)
		}
		return slc
	}
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("ListNode Heap (%d):\n", h.Len()))
	for i, n := range h {
		sb.WriteString(fmt.Sprintf("\t - %d: %v\n", i, LN2Slc(n)))
	}
	return sb.String()
}

// merge k sorted ListNodes
func (q Q0023) MergeKLists(lists []*Q0023ListNode) *Q0023ListNode {
	lists = q.RMNilInList(lists)
	if len(lists) == 0 {
		return nil
	}
	h := Q0023ListNodeHeap(lists)
	heap.Init(&h)

	head := heap.Pop(&h).(*Q0023ListNode)
	curr := head
	if head.Next != nil {
		heap.Push(&h, head.Next)
	}

	for len(h) > 0 {
		n := heap.Pop(&h).(*Q0023ListNode)
		curr.Next = n
		curr = curr.Next
		if n.Next != nil {
			heap.Push(&h, n.Next)
		}
	}
	return head
}

// remove len(0) elements in lists
func (q Q0023) RMNilInList(ln []*Q0023ListNode) []*Q0023ListNode {
	i, j := 0, 0
	for {
		for j < len(ln) && ln[j] == nil {
			j++
		}
		if j < len(ln) {
			ln[i], i, j = ln[j], i+1, j+1
		} else {
			break
		}
	}
	return ln[:i]
}
