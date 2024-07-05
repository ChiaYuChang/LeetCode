package leetcode

import (
	"container/heap"
	"fmt"
	"strings"
)

type Q1438 struct{}

type Q1438TwoQueue struct {
	Start   int
	End     int
	Limit   int
	MinHeap *Q1438Heap
	MaxHeap *Q1438Heap
	Nums    []int
}

func (q Q1438TwoQueue) Len() int {
	return q.End - q.Start + 1
}

func (q Q1438TwoQueue) Min() int {
	return q.MinHeap.Peek().Value
}

func (q Q1438TwoQueue) Max() int {
	return q.MaxHeap.Peek().Value
}

func (q *Q1438TwoQueue) Add(val, pos int) {
	q.End = pos
	item := &Q1438Item{val, pos}
	change := 0
	if q.MinHeap.Len() > 0 && q.MinHeap.Peek().Value > item.Value {
		change = -1
	}

	if q.MaxHeap.Len() > 0 && q.MaxHeap.Peek().Value < item.Value {
		change = 1
	}
	heap.Push(q.MinHeap, item)
	heap.Push(q.MaxHeap, item)
	if change == 0 {
		return
	}

	diff := q.MaxHeap.Peek().AbsDiff(q.MinHeap.Peek())
	// fmt.Println("Diff: ", diff, "Change:", change)
	if diff <= q.Limit {
		return
	}

	var hp *Q1438Heap = q.MaxHeap
	if change > 0 {
		hp = q.MinHeap
	}

	var prev, curr *Q1438Item
	prev = heap.Pop(hp).(*Q1438Item)
	for {
		curr = hp.Peek()
		// fmt.Printf("\t- %2d...", curr.Value)
		if curr.Index > prev.Index {
			if curr.AbsDiff(item) <= q.Limit {
				// fmt.Println("Keep")
				break
			}
			prev = curr
		}
		// fmt.Println("Pop")
		heap.Pop(hp)
	}

	// fmt.Println("Calc Start")
	for i := curr.Index; i > prev.Index; i-- {
		diff := q.Nums[i] - item.Value
		if diff < 0 {
			diff = -diff
		}
		// fmt.Printf("\t- %2d: |%2d - %2d| = %2d\n",
		// 	i, q.Nums[i], item.Value, diff)
		if diff > q.Limit {
			break
		}
		q.Start = i
	}
}

type Q1438Item struct {
	Value int
	Index int
}

func (n Q1438Item) String() string {
	return fmt.Sprintf("{Val: %d, Pos: %d}", n.Value, n.Index)
}

func (i1 *Q1438Item) AbsDiff(i2 any) int {
	i := i2.(*Q1438Item)
	d := i1.Value - i.Value
	if d < 0 {
		return -d
	}
	return d
}

type Q1438Heap struct {
	Order bool // asc: true , desc: false
	Items []*Q1438Item
}

func NewQ1438Heap(order bool) *Q1438Heap {
	return &Q1438Heap{Order: order}
}

func (h Q1438Heap) Len() int {
	return len(h.Items)
}

func (h Q1438Heap) Less(i, j int) bool {
	ii, ij := h.Items[i], h.Items[j]
	if ii.Value == ij.Value {
		return ii.Index > ij.Index
	}

	if h.Order {
		// min heap
		return ii.Value < ij.Value
	}
	// max heap
	return ii.Value > ij.Value
}

func (h *Q1438Heap) Push(x any) {
	h.Items = append(h.Items, x.(*Q1438Item))
}

func (h *Q1438Heap) Pop() any {
	items := h.Items
	n := len(items)
	x := items[n-1]
	h.Items = items[:n-1]
	return x
}

func (h *Q1438Heap) Peek() *Q1438Item {
	return h.Items[0]
}

func (h Q1438Heap) Swap(i, j int) {
	h.Items[i], h.Items[j] = h.Items[j], h.Items[i]
}

func (q Q1438) TwoQueuesLongestSubarray(nums []int, limit int) int {
	twoQueue := Q1438TwoQueue{
		Start:   0,
		End:     0,
		Limit:   limit,
		MinHeap: NewQ1438Heap(true),
		MaxHeap: NewQ1438Heap(false),
		Nums:    nums,
	}
	var ans = 0
	for i := range nums {
		// fmt.Println("Add ", nums[i])
		twoQueue.Add(nums[i], i)
		// fmt.Printf("Index: [%d, %d] %d, Range: [%d, %d] %d\n",
		// 	twoQueue.Start, twoQueue.End, twoQueue.Len(), twoQueue.Min(), twoQueue.Max(), twoQueue.Max()-twoQueue.Min())
		ans = q.max(ans, twoQueue.Len())
	}
	return ans
}

type Deques struct {
	Order bool // asc: true , desc: false
	Data  []int
}

func (dq Deques) String() string {
	sb := strings.Builder{}
	if dq.Order {
		sb.WriteString("MinQueue:")
	} else {
		sb.WriteString("MaxQueue:")
	}
	sb.WriteString(fmt.Sprintf("%v", dq.Data))
	return sb.String()
}

func (dq Deques) Len() int {
	return len(dq.Data)
}

func (dq Deques) Last() int {
	return dq.Data[dq.Len()-1]
}

func (dq Deques) First() int {
	return dq.Data[0]
}

func (dq *Deques) Enqueue(x int) {
	if dq.Order {
		for dq.Len() > 0 && dq.Last() > x {
			dq.Data = dq.Data[:dq.Len()-1]
		}
	} else {
		for dq.Len() > 0 && dq.Last() < x {
			dq.Data = dq.Data[:dq.Len()-1]
		}
	}
	dq.Data = append(dq.Data, x)
}

func (dq *Deques) Dequeue() {
	dq.Data = dq.Data[1:]
}

func (q Q1438) TwoDequesLongestSubarray(nums []int, limit int) int {
	maxQ := &Deques{Order: false}
	minQ := &Deques{Order: true}
	maxLen := 0

	for lhs, rhs := 0, 0; rhs < len(nums); rhs++ {
		// fmt.Printf("[%d, %d] %d\n", lhs, rhs, rhs-lhs+1)
		// fmt.Println("add:", nums[rhs])
		maxQ.Enqueue(nums[rhs])
		minQ.Enqueue(nums[rhs])
		// fmt.Printf("\temaxQ: %v\n", maxQ)
		// fmt.Printf("\teminQ: %v\n", minQ)
		for maxQ.First()-minQ.First() > limit {
			if maxQ.First() == nums[lhs] {
				maxQ.Dequeue()
			}
			if minQ.First() == nums[lhs] {
				minQ.Dequeue()
			}
			lhs++
		}
		// fmt.Printf("\tsmaxQ: %v\n", maxQ)
		// fmt.Printf("\tsminQ: %v\n", minQ)
		maxLen = q.max(maxLen, rhs-lhs+1)
	}
	return maxLen
}

func (q Q1438) max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
