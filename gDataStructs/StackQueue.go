package gdatastructs

import (
	"fmt"
	"strings"
)

// type SQ[T any] interface {
// 	IsEmpry() bool
// 	Peek() (T, bool)
// 	Push(data T) SQ[T]
// 	Pop(data T) (T, bool)
// 	Enqueue(data T) SQ[T]
// 	Dequeue(data T) (T, bool)
// 	String() string
// 	ToArray() []T
// }

type StackQueue[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func NewStackQueue[T any]() *StackQueue[T] {
	return &StackQueue[T]{Length: 0}
}

func ArrayToStackQueue[T any](data []T, rev bool) *StackQueue[T] {
	ll := NewStackQueue[T]()
	if rev {
		for _, d := range data {
			ll.Push(d)
		}
	} else {
		for _, d := range data {
			ll.Enqueue(d)
		}
	}
	return ll
}

func StackQueueToArray[T any](ll StackQueue[T]) []T {
	if ll.Length < 0 {
		return nil
	}

	a := make([]T, ll.Length)
	i := 0
	currNode := ll.Head
	for ; currNode != nil; i += 1 {
		a[i] = *currNode.data
		currNode = currNode.Next()
	}
	return a
}

func (sq *StackQueue[T]) IsEmpty() bool {
	return sq.Length == 0
}

func (sq *StackQueue[T]) Peek(which string) (T, bool) {
	if sq.Length < 1 {
		return *new(T), false
	}
	switch which {
	case "head", "h":
		return *sq.Head.data, true
	case "tail", "t":
		return *sq.Tail.data, true
	default:
		return *new(T), false
	}
}

// adds an element to the top of the stackqueue
func (sq *StackQueue[T]) Push(data T) *StackQueue[T] {
	return sq.PushNode(NewNode(data))
}

func (sq *StackQueue[T]) PushNode(newHead *Node[T]) *StackQueue[T] {
	oldHead := sq.Head
	sq.Head = newHead
	if sq.Length == 0 {
		sq.Tail = newHead
	} else {
		newHead.next = oldHead
		oldHead.prev = newHead
	}
	sq.Length += 1
	return sq
}

// removes and returns the topmost element from the stackqueue
func (sq *StackQueue[T]) Pop() (T, bool) {
	node, ok := sq.PopNode()
	if !ok {
		return *new(T), false
	}
	return *node.data, true
}

func (sq *StackQueue[T]) PopNode() (*Node[T], bool) {
	if sq.IsEmpty() {
		return nil, false
	}
	newHead := sq.Head.next // might be nil
	oldHead := sq.Head

	sq.Head = newHead
	sq.Length -= 1
	if sq.Length == 0 {
		sq.Tail = nil
	}
	return oldHead, true
}

// adds an element to the tail of the stackqueue
func (sq *StackQueue[T]) Enqueue(data T) *StackQueue[T] {
	return sq.EnqueueNode(NewNode(data))
}

func (sq *StackQueue[T]) EnqueueNode(newTail *Node[T]) *StackQueue[T] {
	oldTail := sq.Tail

	sq.Tail = newTail
	if sq.Length == 0 {
		sq.Head = newTail
	} else {
		newTail.prev = oldTail
		oldTail.next = newTail
	}
	sq.Length += 1
	return sq
}

// an alias of pop
func (sq *StackQueue[T]) Dequeue() (T, bool) {
	return sq.Pop()
}

// an alias of popNode
func (sq *StackQueue[T]) DequeueNode() (*Node[T], bool) {
	return sq.PopNode()
}

// create an array from stackqueue
func (sq *StackQueue[T]) ToArray() []T {
	if sq.Length < 1 {
		return nil
	}
	return StackQueueToArray(*sq)
}

func (sq *StackQueue[T]) Walk(c chan<- T, rev bool) {
	if rev {
		sq.Tail.Walk(c, rev)
	} else {
		sq.Head.Walk(c, rev)
	}
}

func (sq *StackQueue[T]) String() string {
	if sq.Length < 1 {
		return "StackQueue(0)"
	}

	sb := strings.Builder{}
	c := make(chan T)
	sq.Walk(c, false)
	sb.WriteString(fmt.Sprintf("StackQueue(%d): ", sq.Length))
	sb.WriteString("(head) ->")
	for d := range c {
		sb.WriteString(fmt.Sprintf(" %v ", d))
		sb.WriteString("->")
	}
	sb.WriteString(" (tail)")
	return sb.String()
}
