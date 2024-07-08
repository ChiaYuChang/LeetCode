package leetcode

import "fmt"

type Q1823 struct{}

func (q Q1823) FindTheWinner(n int, k int) int {
	type node struct {
		value int
		next  *node
	}

	nextK := func(n *node, k int) *node {
		for i := 0; i < k; i++ {
			n = n.next
		}
		return n
	}

	if n == 1 || k == 1 {
		return n
	}

	var head, curr *node
	head = &node{0, nil}
	curr = head
	for i := 1; i <= n; i++ {
		curr.next = &node{i, nil}
		curr = curr.next
	}
	curr, curr.next = nextK(head, k-1), head.next
	for curr.value != curr.next.value {
		fmt.Println("remove:", curr.next.value)
		curr.next = curr.next.next
		curr = nextK(curr, k-1)
	}
	return curr.value
}
