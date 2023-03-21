package leetcode

import (
	"container/heap"
	"fmt"
	"strings"
)

type Q0023ListNode struct {
	Val  int
	Next *Q0023ListNode
}

type Q0023LinkedList struct {
	head *Q0023ListNode
}

func (ll *Q0023LinkedList) HasNext() bool {
	return ll.head != nil
}

func (ll *Q0023LinkedList) Next() *Q0023ListNode {
	if !ll.HasNext() {
		return nil
	}
	n := ll.head
	ll.head = n.Next
	return n
}

func (ll Q0023LinkedList) String() string {
	sb := strings.Builder{}
	sb.WriteString("LinkedList: ")
	for c := ll.head; c != nil; c = c.Next {
		sb.WriteString(fmt.Sprintf("%d -> ", c.Val))
	}
	sb.WriteString("end")
	return sb.String()
}

type Q0023Data struct {
	Node  *Q0023ListNode
	Index int
}

func (d Q0023Data) Val() int {
	return d.Node.Val
}

// An Q0023DataHeap is a min-heap of ints.
type Q0023DataHeap []*Q0023Data

func (h Q0023DataHeap) Len() int      { return len(h) }
func (h Q0023DataHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h Q0023DataHeap) Less(i, j int) bool {
	if h[i] != nil && h[j] != nil {
		return h[i].Node.Val < h[j].Node.Val
	}
	if h[i] == nil && h[j] != nil {
		return true
	}
	return false
}

func (h *Q0023DataHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Q0023Data))
}

func (h *Q0023DataHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Q0023 struct{}

func (q Q0023) MergeLists(lists ...*Q0023ListNode) *Q0023ListNode {
	ll := make([]*Q0023LinkedList, len(lists))
	dt := make([]*Q0023Data, 0, len(lists))
	hp := Q0023DataHeap(dt)
	for i := range lists {
		ll[i] = &Q0023LinkedList{head: lists[i]}
		if ll[i].HasNext() {
			fmt.Println("Push:", ll[i].Next().Val)
			hp = append(hp, &Q0023Data{Node: ll[i].Next(), Index: i})
		}
	}
	fmt.Println(hp.Len())
	heap.Init(&hp)
	root := heap.Pop(&hp).(*Q0023Data).Node
	curr := root
	for hp.Len() > 0 {
		data := heap.Pop(&hp).(*Q0023Data)
		fmt.Println("Add: ", data.Val())
		curr.Next = data.Node
		curr = curr.Next
		if ll[data.Index].HasNext() {
			fmt.Println("Push:", ll[data.Index].Next().Val)
			heap.Push(&hp, &Q0023Data{Node: ll[data.Index].Next(), Index: data.Index})
		}
	}
	return root
}
