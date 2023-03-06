package leetcode

type q234ListNode struct {
	Val  int
	Next *q234ListNode
}

func IsPalindrome(head *q234ListNode) bool {
	if head.Next == nil {
		return true
	}

	// get mid point
	p1, p2 := head, head
	cutsite := p1
	for p2 != nil && p2.Next != nil {
		cutsite, p1, p2 = p1, p1.Next, p2.Next.Next
	}

	if p2 != nil {
		cutsite, p1 = p1, p1.Next
	}

	p2 = head
	p1 = p1.reverse_()
	rHead := p1

	ans := true
	for p1 != nil && ans {
		ans = p1.Val == p2.Val
		p1 = p1.Next
		p2 = p2.Next
	}

	// reorder linklist
	cutsite.Next = rHead.reverse_()
	return ans
}

func (n *q234ListNode) reverse_() *q234ListNode {
	if n == nil {
		return n
	}

	var n0, n1, n2 *q234ListNode
	n1 = n
	if n1 == nil {
		n.Next = n1
		return n
	}
	n2 = n1.Next
	for n1 != nil && n2 != nil {
		n1.Next = n0
		n0 = n1
		n1 = n2
		n2 = n1.Next
	}
	n1.Next = n0
	n = n1
	return n
}
