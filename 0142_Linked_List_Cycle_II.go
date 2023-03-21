package leetcode

type Q0142 struct{}

type Q0142ListNode struct {
	Val  int
	Next *Q0142ListNode
}

func (q Q0142) DetectCycle(head *Q0142ListNode) *Q0142ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	p1, p2 := head.Next, head.Next.Next
	for p1 != p2 {
		if (p2.Next != nil) && (p2.Next.Next != nil) {
			p1, p2 = p1.Next, p2.Next.Next
		} else {
			return nil
		}
	}

	p2 = head
	for p1 != p2 {
		p1, p2 = p1.Next, p2.Next
	}

	return p1
}
