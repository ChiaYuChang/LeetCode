package leetcode

import (
	"fmt"
	"strings"
)

type Q0984 struct {
	cache Q0984Cache
}

func NewQ0984() Q0984 {
	q := Q0984{}

	c := Q0984Cache{
		data: map[int][]*Q0894TreeNode{},
		dcp:  q.TreeNodeDeepCopy,
	}
	c.data[1] = []*Q0894TreeNode{{Val: 0}}
	q.cache = c
	return q
}

type Q0894TreeNode struct {
	Val   int
	Left  *Q0894TreeNode
	Right *Q0894TreeNode
}

func (n *Q0894TreeNode) String() string {
	var ToString func(Nindent int, n *Q0894TreeNode, sb *strings.Builder)

	ToString = func(Nindent int, n *Q0894TreeNode, sb *strings.Builder) {
		if n == nil {
			return
		}
		ToString(Nindent+1, n.Left, sb)
		sb.WriteString(fmt.Sprintf("%s%d\n", strings.Repeat("\t", Nindent), n.Val))
		ToString(Nindent+1, n.Right, sb)
	}
	sb := strings.Builder{}
	sb.WriteString("TreeNode:\n")
	ToString(0, n, &sb)
	return sb.String()
}

func (q Q0984) AllPossibleFBT(n int) []*Q0894TreeNode {
	if n%2 == 0 {
		return nil
	}
	return q.cache.Get(n)
}

func (q Q0984) TreeNodeDeepCopy(root *Q0894TreeNode) *Q0894TreeNode {
	if root == nil {
		return nil
	}

	var cp func(dst, src *Q0894TreeNode)
	cp = func(dst, src *Q0894TreeNode) {
		// set current node value
		dst.Val = src.Val

		// copy left branch
		if src.Left != nil {
			dst.Left = &Q0894TreeNode{}
			cp(dst.Left, src.Left)
		}

		// copy right branch
		if src.Right != nil {
			dst.Right = &Q0894TreeNode{}
			cp(dst.Right, src.Right)
		}
	}

	dst := &Q0894TreeNode{Val: root.Val}
	cp(dst, root)
	return dst
}

type Q0984Cache struct {
	data map[int][]*Q0894TreeNode
	dcp  func(root *Q0894TreeNode) *Q0894TreeNode
}

func (l *Q0984Cache) GetFromCache(n int) []*Q0894TreeNode {
	roots := l.data[n]
	entry := make([]*Q0894TreeNode, len(roots))
	for i, root := range roots {
		entry[i] = l.dcp(root)
	}
	return entry
}

func (l *Q0984Cache) Get(n int) []*Q0894TreeNode {
	if n%2 == 0 {
		return nil
	}

	if _, ok := l.data[n]; ok {
		return l.GetFromCache(n)
	}

	entry := []*Q0894TreeNode{}
	for n1, n2 := 1, n-2; n2 > 0; {
		lhs := l.Get(n1)
		rhs := l.Get(n2)
		for _, l := range lhs {
			for _, r := range rhs {
				root := &Q0894TreeNode{Val: 0}
				root.Left = l
				root.Right = r
				entry = append(entry, root)
			}
		}
		n1 += 2
		n2 -= 2
	}
	l.data[n] = entry

	return l.GetFromCache(n)
}

func (q Q0984) AllPossibleFBT2(n int) []*Q0894TreeNode {
	if n%2 == 0 {
		return nil
	}

	if n == 1 {
		return []*Q0894TreeNode{{Val: 0}}
	}

	entry := []*Q0894TreeNode{}
	for n1, n2 := 1, n-2; n2 > 0; {
		lhs := q.AllPossibleFBT2(n1)
		rhs := q.AllPossibleFBT2(n2)
		for _, l := range lhs {
			for _, r := range rhs {
				root := &Q0894TreeNode{Val: 0}
				root.Left = l
				root.Right = r
				entry = append(entry, root)
			}
		}
		n1 += 2
		n2 -= 2
	}
	return entry
}
