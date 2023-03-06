package gdatastructs

import (
	"fmt"
	"strings"
)

var defaultHeapInitCap = 100

type Heap[TKey OrderedStruct, TVal any] struct {
	data []*KeyData[TKey, TVal]
}

func NewHeap[TKey OrderedStruct, TVal any]() *Heap[TKey, TVal] {
	return newHeapWithCap[TKey, TVal](defaultHeapInitCap)
}

func newHeapWithCap[TKey OrderedStruct, TVal any](cap int) *Heap[TKey, TVal] {
	return &Heap[TKey, TVal]{
		data: make([]*KeyData[TKey, TVal], 1, cap+1),
	}
}

func (h *Heap[TKey, TVal]) Size() int {
	return len(h.data) - 1
}

func (h *Heap[TKey, TVal]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *Heap[TKey, TVal]) Top() (TKey, TVal) {
	d := h.data[1]
	return d.key, d.val
}

func (h *Heap[TKey, TVal]) Pop() (TKey, TVal) {
	k, v := h.Top()
	h.swap(1, h.Size())
	h.data = h.data[:h.Size()]
	h.sink(1)
	return k, v
}

func (h *Heap[TKey, TVal]) Insert(key TKey, val TVal) {
	h.data = append(h.data, NewKeyData(key, val))
	h.swim(h.Size())
}

func (h *Heap[TKey, TVal]) sink(i int) {
	for i*2 <= h.Size() {
		j := i * 2
		if j < h.Size() {
			// get child with greater val
			if h.less(j, j+1) {
				j = j + 1
			}

			// exchange data[i] and data[j] if data[i] < data[j]
			if !h.less(i, j) {
				h.swap(i, j)
			} else {
				break
			}
		}
		i = j
	}
}

func (h *Heap[TKey, TVal]) swap(i, j int) {
	fmt.Printf("swap: %d %d\n", i, j)
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[TKey, TVal]) less(i, j int) bool {
	diIsNil := h.data[i] == nil
	djIsNil := h.data[j] == nil

	// nil does not less than nil
	if diIsNil && djIsNil {
		return false
	}

	// everything is greater than nil
	if diIsNil {
		return true
	}

	if djIsNil {
		return false
	}

	// both d1 and d2 are not nil, compare
	return h.data[i].key.Less(h.data[j].key)
}

func (h *Heap[TKey, TVal]) swim(i int) {
	fmt.Printf("swim: %d\n", i)
	for i > 1 && h.less(i/2, i) {
		h.swap(i, i/2)
		i = i / 2
	}
}

func (h *Heap[TKey, TVal]) String() string {
	sb := strings.Builder{}
	for _, d := range h.data[1:] {
		sb.WriteString(fmt.Sprintf("%v: %v", d.key, d.val))
		sb.WriteString(" -> ")
	}
	return sb.String()
}
