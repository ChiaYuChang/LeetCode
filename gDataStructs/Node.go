package gdatastructs

import (
	"errors"
)

var ErrBranchIsNil = errors.New("the branch is not nil")

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]
	data *T
}

// Create a new node
func NewNode[T any](data T) *Node[T] {
	return &Node[T]{data: &data}
}

// Create a list node from array
func ArrayToNodes[T any](data []T, rev bool) (head *Node[T]) {
	if len(data) == 0 {
		return
	}

	cn := NewNode(data[0])
	head = cn
	for i := 1; i < len(data); i++ {
		if rev {
			cn.prev = NewNode(data[i])
			cn.prev.next = cn
			cn = cn.prev
		} else {
			cn.next = NewNode(data[i])
			cn.next.prev = cn
			cn = cn.next
		}
	}
	return
}

// Return next node
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Append data to the next node, return error if the next node is not empty.
// Set force to true to omit error
func (n *Node[T]) AppenDataToNext(data T, force bool) (*Node[T], error) {
	return n.AppenToNext(NewNode(data), force)
}

func (n *Node[T]) AppenToNext(node *Node[T], force bool) (*Node[T], error) {
	if force || n.HasNext() {
		return n, ErrBranchIsNil
	}
	n.next = node
	return n, nil
}

// Return previous node
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

// Append data to the previous node, return error if the next node is not empty.
// Set force to true to omit error
func (n *Node[T]) AppenDataToPrev(data T, force bool) (*Node[T], error) {
	return n.AppendToPrev(NewNode(data), force)
}

func (n *Node[T]) AppendToPrev(node *Node[T], force bool) (*Node[T], error) {
	if force || n.HasPrev() {
		return n, ErrBranchIsNil
	}
	n.prev = node
	return n, nil
}

// Return data of the node
func (n *Node[T]) Data() *T {
	return n.data
}

// Check whether the next brance is empty
func (n *Node[T]) HasNext() bool {
	return n.next != nil
}

// Check whether the previous brance is empty
func (n *Node[T]) HasPrev() bool {
	return n.prev != nil
}

func (n *Node[T]) Walk(c chan<- T, rev bool) {
	go func() {
		c <- *n.data
		if rev {
			for n.prev != nil {
				n = n.prev
				c <- *n.data
				// fmt.Println(*n.data)
			}
		} else {
			for n.next != nil {
				n = n.next
				c <- *n.data
				// fmt.Println(*n.data)
			}
		}
		close(c)
	}()
}

// func (n *Node[T]) String() string {
// fmt.Sprintf("{node: left: %, right: %v", n.)
// }
