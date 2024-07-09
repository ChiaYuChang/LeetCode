package leetcode

import (
	"strconv"
	"strings"
)

type Q1038 struct{}

type Q1038TreeNode struct {
	Val   int
	Left  *Q1038TreeNode
	Right *Q1038TreeNode
}

func (t Q1038TreeNode) String() string {
	ll := 1
	curr := &t
	for curr != nil {
		ll *= 2
		curr = curr.Right
	}

	l := make([]string, ll)
	for i := 0; i < ll; i++ {
		l[i] = "null"
	}

	var travel func(t *Q1038TreeNode, index int)
	travel = func(t *Q1038TreeNode, index int) {
		if t == nil {
			return
		}

		l[index] = strconv.Itoa(t.Val)
		travel(t.Left, index*2)
		travel(t.Right, index*2+1)
	}
	travel(&t, 1)
	return "[" + strings.Join(l[1:], ",") + "]"
}

func (q Q1038) BSTToGst(root *Q1038TreeNode) *Q1038TreeNode {
	var helper func(*Q1038TreeNode, int) int
	helper = func(root *Q1038TreeNode, acc int) int {
		if root == nil {
			return acc
		}

		root.Val += helper(root.Right, acc)
		if root.Left != nil {
			return helper(root.Left, root.Val)
		}
		return root.Val
	}
	helper(root, 0)
	return root
}
