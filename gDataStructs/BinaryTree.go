package gdatastructs

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type BinaryTree[TKey constraints.Ordered, TVal any] interface {
	Append(key TKey, val TVal)
	Find(key TKey) (TVal, bool)
	Size() int
	Diameter() int
	LowestCommonAncestor(x, y TKey) (TKey, TVal)
}

// Binary Tree implemented by nodes
type nBTData[TKey constraints.Ordered, TVal any] struct {
	key TKey
	val TVal
}

// Node for binary trees
type nBTNode[TKey constraints.Ordered, TVal any] *Node[nBTData[TKey, TVal]]

// Create a node for binary trees
func newBTNode[TKey constraints.Ordered, TVal any](key TKey, val TVal) nBTNode[TKey, TVal] {
	return nBTNode[TKey, TVal](
		&Node[nBTData[TKey, TVal]]{
			data: &nBTData[TKey, TVal]{key, val},
		},
	)
}

// Insert given node into a binary tree
func appendBTNode[TKey constraints.Ordered, TVal any](btn nBTNode[TKey, TVal], key TKey, val TVal) bool {
	if btn == nil {
		return true
	}

	if btn.data.key > key {
		if btn.prev == nil {
			btn.prev = newBTNode(key, val)
			fmt.Println(key, "|--", btn.data.key)
		} else {
			fmt.Println("<--", btn.data.key)
			appendBTNode(btn.prev, key, val)
		}
	} else if btn.data.key < key {
		if btn.next == nil {
			btn.next = newBTNode(key, val)
			fmt.Println(btn.data.key, "--|", key)
		} else {
			fmt.Println(btn.data.key, "-->")
			appendBTNode(btn.next, key, val)
		}
	} else {
		btn.data.val = val
		return false
	}
	return true
}

// Find given node in a binary tree by key
func findBTNode[TKey constraints.Ordered, TVal any](btn nBTNode[TKey, TVal], key TKey) (TVal, bool) {
	if btn == nil {
		return *new(TVal), false
	}

	if btn.data.key > key {
		// fmt.Println("<--")
		return findBTNode(btn.prev, key)
	} else if btn.data.key < key {
		// fmt.Println("-->")
		return findBTNode(btn.next, key)
	} else {
		return btn.data.val, true
	}
}

func findFstBetween[TKey constraints.Ordered, TVal any](btn nBTNode[TKey, TVal], x, y TKey, lb, ub bool) nBTNode[TKey, TVal] {
	if x < y {
		x, y = y, x
	}

	if (x > btn.data.key && btn.data.key > y) ||
		(x == btn.data.key && ub) ||
		(y == btn.data.key && lb) {
		return btn
	}

	// right branch
	if x < btn.data.key {
		return findFstBetween(btn.prev, x, y, lb, ub)
	}
	// left branch
	return findFstBetween(btn.next, x, y, lb, ub)
}

// Binary tree implemented by BTnodes
type nBT[TKey constraints.Ordered, TVal any] struct {
	root nBTNode[TKey, TVal]
	size int
}

// Insert given node into a binary tree
func (btr *nBT[TKey, TVal]) Append(key TKey, val TVal) {
	if btr.size == 0 {
		btr.root = newBTNode(key, val)
		btr.size++
		return
	}

	if appendBTNode(btr.root, key, val) {
		btr.size++
	}
}

// Find given node in a binary tree by key
func (btr *nBT[TKey, TVal]) Find(key TKey) (TVal, bool) {
	if btr.size == 0 {
		return *new(TVal), false
	}
	return findBTNode(btr.root, key)
}

// Return the size of the binary tree
func (btr *nBT[TKey, TVal]) Size() int {
	return btr.size
}

func (btr *nBT[TKey, TVal]) LowestCommonAncestor(x, y TKey) (TKey, TVal) {
	if btr.size == 0 {
		return *new(TKey), *new(TVal)
	}
	n := findFstBetween(btr.root, x, y, true, true)
	return n.data.key, n.data.val
}

func newnBinaryTree[TKey constraints.Ordered, TVal any]() *nBT[TKey, TVal] {
	return &nBT[TKey, TVal]{size: 0}
}

func (btr *nBT[TKey, TVal]) Diameter() int {
	if btr.size == 0 {
		return 0
	}

	var diameter int
	btrTravel(btr.root, 0, &diameter)

	return diameter
}

func btrTravel[TKey constraints.Ordered, TVal any](n nBTNode[TKey, TVal], depth int, diameter *int) int {
	if n == nil {
		return depth
	}

	depth++

	ll := btrTravel(n.next, depth, diameter)
	rl := btrTravel(n.prev, depth, diameter)

	if ll+rl-2*depth > *diameter {
		*diameter = ll + rl - 2*depth
	}

	if ll > rl {
		return ll
	}
	return rl
}

// Create a binary tree. method should be either "n/node" nor "a/array"
func NewBinaryTree[TKey constraints.Ordered, TVal any](method string) BinaryTree[TKey, TVal] {
	switch method {
	case "n", "node":
		return newnBinaryTree[TKey, TVal]()
	case "a", "array":
		return nil
	default:
		return nil
	}
}
