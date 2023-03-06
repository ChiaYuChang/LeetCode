package leetcode

type q237ListNode struct {
	Val  int
	Next *q237ListNode
}

func DeleteCurrentNodeVal(node *q237ListNode) {
	if node.Next == nil {
		// should not be the last node
		return
	}

	// method 1: remove current node
	(*node) = (*node.Next)

	// method 2: copy and remove next node
	// node.Val = node.Next.Val
	// node.Next = node.Next.Next

	// method 3: shift values
	// n0, n1 := node, node.Next
	// for n1 != nil {
	//     n0.Val = n1.Val

	//     if n1.Next == nil {
	//         // n1 is the last node
	//         n0.Next = nil
	//         n1 = nil // remove the last node
	//     } else {
	//         n0, n1 = n1, n1.Next
	//     }
	// }
}
