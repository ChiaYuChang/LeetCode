package leetcode

type Q0109ListNode struct {
	Val  int
	Next *Q0109ListNode
}

type Q0109TreeNode struct {
	Val   int
	Left  *Q0109TreeNode
	Right *Q0109TreeNode
}

type Q0109 struct{}

func (q Q0109) SortedListToBST(head *Q0109ListNode) *Q0109TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &Q0109TreeNode{Val: head.Val}
	}

	p0, p1, p2 := head, head, head
	for p2 != nil && p2.Next != nil {
		p0, p1, p2 = p1, p1.Next, p2.Next.Next
	}

	p0.Next = nil
	n := &Q0109TreeNode{Val: p1.Val}
	n.Left = q.SortedListToBST(head)
	n.Right = q.SortedListToBST(p1.Next)
	p0.Next = p1
	return n
}
