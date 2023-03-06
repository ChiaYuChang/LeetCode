package gdatastructs

type LinkList[T any] struct {
	cn *Node[T]
	sq *StackQueue[T]
}

func NewLinkList[T any]() *LinkList[T] {
	return &LinkList[T]{
		cn: nil,
		sq: NewStackQueue[T](),
	}
}

func NewLinkListFromArray[T any](data []T, rev bool) *LinkList[T] {
	ll := NewLinkList[T]()
	ll.sq = ArrayToStackQueue(data, rev)
	ll.cn = ll.sq.Head
	return ll
}

func (ll *LinkList[T]) HasNext() bool {
	return ll.cn.HasNext()
}

func (ll *LinkList[T]) HasPrev() bool {
	return ll.cn.HasPrev()
}

func (ll *LinkList[T]) Append(data ...T) *LinkList[T] {
	if len(data) < 1 {
		return ll
	}

	if ll.sq.IsEmpty() {
		ll.sq = ArrayToStackQueue(data, false)
		ll.cn = ll.sq.Head
		return ll
	}

	for _, d := range data {
		ll.sq.Enqueue(d)
	}
	return ll
}

func (ll *LinkList[T]) Data() (T, bool) {
	if ll.sq.IsEmpty() {
		return *new(T), false
	}
	return *ll.cn.data, true
}

func (ll *LinkList[T]) Forward() *LinkList[T] {
	ll.cn = ll.cn.Next()
	return ll
}

func (ll *LinkList[T]) Backward() *LinkList[T] {
	ll.cn = ll.cn.Prev()
	return ll
}

func (ll *LinkList[T]) Walk(c chan<- T) {
	ll.sq.Head.Walk(c, false)
}

func (ll *LinkList[T]) Reverse() *LinkList[T] {
	if ll.sq == nil {
		return ll
	}

	var n0, n1, n2 *Node[T]
	n1 = ll.sq.Head
	if n1 == nil {
		return ll
	}
	n2 = n1.next
	for n1 != nil && n2 != nil {
		n1.next = n0
		n0 = n1
		n1 = n2
		n2 = n1.next
	}
	n1.next = n0
	ll.sq.Head = n1
	return ll
}
