package leetcode

import "container/heap"

type Q0373 struct{}

// priority, num1Index, num2Index
type Q0373Triple [3]int

func (q Q0373) NewTriple(x, y, xi, yi int) *Q0373Triple {
	return &Q0373Triple{x + y, xi, yi}
}

type Q0373CoupleHeap []*Q0373Triple

func (h Q0373CoupleHeap) Len() int           { return len(h) }
func (h Q0373CoupleHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h Q0373CoupleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Q0373CoupleHeap) Push(x any) {
	*h = append(*h, x.(*Q0373Triple))
}

func (h *Q0373CoupleHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (q Q0373) KSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	// to minimize the size of CoupleHeap, length nums1 should be less than or
	// equal to nums2
	flip := false
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
		flip = true
	}

	triples := Q0373CoupleHeap(make([]*Q0373Triple, len(nums1)))
	for i := range triples {
		triples[i] = q.NewTriple(nums1[i], nums2[0], i, 0)
	}

	heap.Init(&triples)
	k = min(k, len(nums1)*len(nums2))
	kMin := make([][]int, k)
	for i := 0; i < k; i++ {
		triple := heap.Pop(&triples).(*Q0373Triple)

		if flip {
			kMin[i] = []int{nums2[triple[2]], nums1[triple[1]]}
		} else {
			kMin[i] = []int{nums1[triple[1]], nums2[triple[2]]}
		}

		if triple[2]+1 < len(nums2) {
			heap.Push(&triples, q.NewTriple(
				nums1[triple[1]],
				nums2[triple[2]+1],
				triple[1],
				triple[2]+1,
			))
		}
	}
	return kMin
}
