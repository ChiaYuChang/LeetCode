package leetcode

type Q0021 struct{}

type Q0021ListNode struct {
	Val  int
	Next *Q0021ListNode
}

func (q Q0021) MergeTwoLists(list1 *Q0021ListNode, list2 *Q0021ListNode) *Q0021ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var curr_node, merged_list *Q0021ListNode
	init := true
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			if init {
				curr_node = &Q0021ListNode{Val: list1.Val, Next: nil}
				merged_list = curr_node
				init = false
			} else {
				curr_node.Next = &Q0021ListNode{Val: list1.Val, Next: nil}
				curr_node = curr_node.Next
			}
			list1 = list1.Next
		} else {
			if init {
				curr_node = &Q0021ListNode{Val: list2.Val, Next: nil}
				merged_list = curr_node
				init = false
			} else {
				curr_node.Next = &Q0021ListNode{Val: list2.Val, Next: nil}
				curr_node = curr_node.Next
			}
			list2 = list2.Next
		}
	}

	if list1 != nil {
		curr_node.Next = list1
	}

	if list2 != nil {
		curr_node.Next = list2
	}

	return merged_list
}
