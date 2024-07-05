package leetcode

type Q2181 struct{}

type Q2181ListNode struct {
	Val  int
	Next *Q2181ListNode
}

func (n *Q2181ListNode) ToList() []int {
	var l []int
	curr := n
	for curr != nil {
		l = append(l, curr.Val)
		curr = curr.Next
	}
	return l
}

func (q Q2181) NewListNodeFromList(vals []int) *Q2181ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &Q2181ListNode{}
	curr := head
	for i, v := range vals {
		curr.Val = v
		if i < len(vals)-1 {
			curr.Next = &Q2181ListNode{}
		}
		curr = curr.Next
	}
	return head
}

func (q Q2181) MergeNodes(head *Q2181ListNode) *Q2181ListNode {
	acc, curr := head, head.Next
	result := acc
	for curr != nil {
		if curr.Val == 0 {
			if curr.Next != nil {
				acc = acc.Next
				acc.Val = 0
			} else {
				acc.Next = nil
			}
		}
		acc.Val += curr.Val
		curr = curr.Next
	}
	return result
}
