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
