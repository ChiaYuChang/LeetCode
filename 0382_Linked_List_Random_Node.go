package leetcode

import (
	"math/rand"
	"time"
)

type Q0382ListNode struct {
	Val  int
	Next *Q0382ListNode
}

type Q0382ListNodeSampler struct {
	Len  int
	Root *Q0382ListNode
	Rng  *rand.Rand
}

func NewQ0382ListNodeSampler(head *Q0382ListNode) Q0382ListNodeSampler {
	if head == nil {
		// should never happen
		panic("head should not be nil")
	}

	curr := head
	n := 0
	for curr != nil {
		n, curr = n+1, curr.Next
	}

	return Q0382ListNodeSampler{
		Len:  n,
		Root: head,
		Rng:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (q *Q0382ListNodeSampler) GetRandomProbability() int {
	curr := q.Root
	n := q.Len

	// p_{i} = (1 - p_{i})p_{i+1}
	// p_{1} = 1/n
	// p_{1}, p_{2}, p_{3} ..., p_{i}, ... = 1/n, 1/(n-1), 1/(n-3) ..., 1/(n-i), ...
	for ; n > 0; n-- {
		if q.Rng.Intn(n) == 0 {
			return curr.Val
		}
		curr = curr.Next
	}
	return curr.Val
}

func (q *Q0382ListNodeSampler) GetRandomCounter() int {
	curr := q.Root
	for n := q.Rng.Intn(q.Len); n > 0; n-- {
		curr = curr.Next
	}
	return curr.Val
}

// Reservoir Sampling
// i = 1 p_1 = 1
// i = 2 p_1 = 1 * 1/2; p_2 = 1/2
// i = 3 p_1 = 1 * 1/2 * 2/3; p_2 = 1/2 * 2/3; p_3 = 1/3
// i = 4 p_1 = 1 * 1/2 * 2/3 * 3/4 ... p_4 = 1/4
// q.Len is not necessary for RS
func (q *Q0382ListNodeSampler) GetRandomReservoirSampling() int {
	ans, n := q.Root.Val, 2
	curr := q.Root.Next
	for curr != nil {
		if q.Rng.Intn(n) == 0 {
			ans = curr.Val
		}
		curr, n = curr.Next, n+1
	}
	return ans
}
