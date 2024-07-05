package leetcode

//! Time Limit Exceeded

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

type Q2040 struct{}

type Quintuple struct {
	priority   int
	n1WhichSeg WhichSegment
	n1Index    int
	n2WhichSeg WhichSegment
	n2Index    int
}

func NewQuintuple(x, y, xi, yi int, wx, wy WhichSegment) *Quintuple {
	return &Quintuple{
		priority:   x * y,
		n1WhichSeg: wx,
		n1Index:    xi,
		n2WhichSeg: wy,
		n2Index:    yi,
	}
}

func (t Quintuple) String() string {
	return fmt.Sprintf("%3d ([%v] %d, [%v] %d)", t.priority, t.n1WhichSeg, t.n1Index, t.n1WhichSeg, t.n2Index)
}

type CoupleHeap []*Quintuple

func (h CoupleHeap) Len() int           { return len(h) }
func (h CoupleHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h CoupleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CoupleHeap) Push(x any) { *h = append(*h, x.(*Quintuple)) }

func (h *CoupleHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type WhichSegment int8

const (
	All WhichSegment = iota
	Pos
	Neg
)

func (ws WhichSegment) String() string {
	switch ws {
	case Pos:
		return "P"
	case Neg:
		return "N"
	default:
		return "A"
	}
}

type Segment struct {
	Neg []int
	Pos []int
}

func (s Segment) Len(which WhichSegment) int {
	switch which {
	case Pos:
		return len(s.Pos)
	case Neg:
		return len(s.Neg)
	default:
		return len(s.Pos) + len(s.Neg)
	}
}

func (s Segment) Get(which WhichSegment) []int {
	switch which {
	case Pos:
		return s.Pos
	case Neg:
		return s.Neg
	default:
		total := make([]int, s.Len(All))
		copy(total[:s.Len(Neg)], s.Neg)
		copy(total[s.Len(Neg):], s.Pos)
		return total
	}
}

func (s Segment) String() string {
	sb := strings.Builder{}
	sb.WriteString("Segments:\n")
	sb.WriteString(fmt.Sprintf("\t- Pos: %v\n", s.Pos))
	sb.WriteString(fmt.Sprintf("\t- Neg: %v\n", s.Neg))
	return sb.String()
}

func KthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	// to minimize the size of CoupleHeap, length nums1 should be less than or
	// equal to nums2
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	n1NegEnd, n1PosStr := sort.SearchInts(nums1, 0), sort.SearchInts(nums1, 1)
	n2NegEnd, n2PosStr := sort.SearchInts(nums2, 0), sort.SearchInts(nums2, 1)

	nNegProds := n1NegEnd*(len(nums2)-n2PosStr) + (len(nums1)-n1PosStr)*n2NegEnd
	nPosProds := n1NegEnd*n2NegEnd + (len(nums1)-n1PosStr)*(len(nums2)-n2PosStr)
	nZeros := len(nums1)*len(nums2) - nNegProds - nPosProds

	// fmt.Println("- # of positive products:", nPosProds)
	// fmt.Println("- # of negative products:", nNegProds)
	// fmt.Println("- # of zeros:", nZeros)

	n1Segs := Segment{nums1[:n1NegEnd], nums1[n1PosStr:]}
	n2Segs := Segment{nums2[:n2NegEnd], nums2[n2PosStr:]}
	if n1Segs.Len(All) > n2Segs.Len(All) {
		n1Segs, n2Segs = n2Segs, n1Segs
	}
	// fmt.Println(n1Segs)
	// fmt.Println(n2Segs)

	var kthMinIsNeg bool
	if k <= int64(nNegProds) {
		kthMinIsNeg = true
	} else if k <= int64(nNegProds+nZeros) {
		// fmt.Println("kthMin: 0")
		return 0
	} else {
		kthMinIsNeg = false
	}

	quintuples := make(CoupleHeap, n1Segs.Len(All))
	if kthMinIsNeg {
		i := 0
		yi := 0
		if n2Segs.Len(Neg) > 0 {
			for xi := range n1Segs.Pos {
				quintuples[i] = NewQuintuple(
					n1Segs.Pos[xi], n2Segs.Neg[yi], xi, yi, Pos, Neg)
				i++
			}
		}

		yi = n2Segs.Len(Pos) - 1
		if n2Segs.Len(Pos) > 0 {
			for xi := range n1Segs.Neg {
				quintuples[i] = NewQuintuple(n1Segs.Neg[xi], n2Segs.Pos[yi], xi, yi, Neg, Pos)
				i++
			}
		}
		// fmt.Printf("k: %d Negative Number\n", k)
	} else {
		k -= int64(nNegProds + nZeros)
		i := 0
		yi := n2Segs.Len(Neg) - 1
		if n2Segs.Len(Neg) > 0 {
			for xi := range n1Segs.Neg {
				quintuples[i] = NewQuintuple(n1Segs.Neg[xi], n2Segs.Neg[yi], xi, yi, Neg, Neg)
				i++
			}
		}
		yi = 0
		if n2Segs.Len(Pos) > 0 {
			for xi := range n1Segs.Pos {
				quintuples[i] = NewQuintuple(n1Segs.Pos[xi], n2Segs.Pos[yi], xi, yi, Pos, Pos)
				i++
			}
		}
		// fmt.Printf("k: %d Positive Number\n", k)
	}

	heap.Init(&quintuples)
	var kthMin int64
	for i := int64(0); i < k; i++ {
		// fmt.Println(quintuples)

		quintuple := heap.Pop(&quintuples).(*Quintuple)
		kthMin = int64(quintuple.priority)

		next := quintuple.n2Index + 1
		if quintuple.n1WhichSeg == Neg {
			next = quintuple.n2Index - 1
		}

		if 0 <= next && next < n2Segs.Len(quintuple.n2WhichSeg) {
			heap.Push(&quintuples, NewQuintuple(
				n1Segs.Get(quintuple.n1WhichSeg)[quintuple.n1Index],
				n2Segs.Get(quintuple.n2WhichSeg)[next],
				quintuple.n1Index,
				next,
				quintuple.n1WhichSeg,
				quintuple.n2WhichSeg,
			))
		}
	}

	// fmt.Printf("kthMin: %d\n", kthMin)
	return kthMin
}
