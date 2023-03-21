package leetcode

import (
	"fmt"
	"strings"
)

type Q0606 struct{}

type Q0606TreeNode struct {
	Val   int
	Left  *Q0606TreeNode
	Right *Q0606TreeNode
}

func (q Q0606) Tree2Str(root *Q0606TreeNode) string {
	if root == nil {
		return ""
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprint(root.Val))
	q.Travel(root, &sb)
	return sb.String()
}

func (q Q0606) Travel(node *Q0606TreeNode, sb *strings.Builder) {
	hasR, hasL := node.Right != nil, node.Left != nil

	if !hasL && !hasR {
		return
	} else if hasL && !hasR {
		sb.WriteString(fmt.Sprintf("(%d", node.Left.Val))
		q.Travel(node.Left, sb)
	} else if !hasL && hasR {
		sb.WriteString(fmt.Sprintf("()(%d", node.Right.Val))
		q.Travel(node.Right, sb)
	} else {
		sb.WriteString(fmt.Sprintf("(%d", node.Left.Val))
		q.Travel(node.Left, sb)
		sb.WriteString(fmt.Sprintf(")(%d", node.Right.Val))
		q.Travel(node.Right, sb)
	}
	sb.WriteRune(')')
}
