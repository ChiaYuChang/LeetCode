package leetcode

type Q0061 struct{}

type Q0061ListNode struct {
	Val  int
	Next *Q0061ListNode
}

func (q Q0061) RotateRight(head *Q0061ListNode, k int) *Q0061ListNode {
	if head == nil {
		return head
	}

	p1, p2 := head, head
	n := 1
	for i := k - 1; i >= 0; i-- {
		if p1.Next != nil {
			p1 = p1.Next
			n++
		} else {
			p1 = head
			i = i % n // speed up
		}
	}

	if p1 == p2 {
		// p1 and p2 point to the same node
		return head
	}

	for p1.Next != nil {
		p1, p2 = p1.Next, p2.Next
	}

	p1.Next, p2.Next, head = head, nil, p2.Next
	return head
}
