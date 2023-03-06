package gdatastructs

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type OrderedStruct interface {
	Less(y OrderedStruct) bool
	String() string
}

type Ordered[T constraints.Ordered] struct {
	val T
}

func NewOrdered[T constraints.Ordered](val T) Ordered[T] {
	return Ordered[T]{val: val}
}

func (x Ordered[T]) Less(y OrderedStruct) bool {
	os, _ := y.(Ordered[T])
	return x.val < os.val
}

func (x Ordered[T]) String() string {
	return fmt.Sprintf("%v", x.val)
}

func (x Ordered[T]) Val() T {
	return x.val
}

type KeyData[TKey OrderedStruct, TVal any] struct {
	key TKey
	val TVal
}

func NewKeyData[TKey OrderedStruct, TVal any](key TKey, val TVal) *KeyData[TKey, TVal] {
	return &KeyData[TKey, TVal]{key, val}
}

func lessKeyData[TKey OrderedStruct, TVal any](d1, d2 *KeyData[TKey, TVal]) bool {
	d1IsNil := d1 == nil
	d2IsNil := d2 == nil

	// nil does not less than nil
	if d1IsNil && d2IsNil {
		return false
	}

	// anything is greater than nil
	if d1IsNil {
		return true
	}

	if d2IsNil {
		return false
	}

	// both d1 and d2 are not nil, compare
	return d1.key.Less(d2.key)
}

func (d1 *KeyData[TKey, TVal]) Less(d2 *KeyData[TKey, TVal]) bool {
	return lessKeyData(d1, d2)
}

type KeyNode[TKey OrderedStruct, TVal any] struct {
	left  *KeyNode[TKey, TVal]
	right *KeyNode[TKey, TVal]
	*KeyData[TKey, TVal]
}

func NewKeyNode[TKey OrderedStruct, TVal any](key TKey, val TVal) *KeyNode[TKey, TVal] {
	return &KeyNode[TKey, TVal]{
		KeyData: &KeyData[TKey, TVal]{
			key: key, val: val,
		},
	}
}

func lessKeyNode[TKey OrderedStruct, TVal any](x, y *KeyNode[TKey, TVal]) bool {
	return lessKeyData(x.KeyData, y.KeyData)
}

func (kn1 *KeyNode[TKey, TVal]) Less(kn2 *KeyNode[TKey, TVal]) bool {
	return lessKeyNode(kn1, kn2)
}

func (kn *KeyNode[TKey, TVal]) AppendToLeft(n *KeyNode[TKey, TVal]) *KeyNode[TKey, TVal] {
	kn.left = n
	return kn
}

func (kn *KeyNode[TKey, TVal]) AppendDataToLeft(key TKey, val TVal) *KeyNode[TKey, TVal] {
	return kn.AppendToLeft(NewKeyNode(key, val))
}

func (kn *KeyNode[TKey, TVal]) AppendToRight(n *KeyNode[TKey, TVal]) *KeyNode[TKey, TVal] {
	kn.right = n
	return kn
}

func (kn *KeyNode[TKey, TVal]) AppendDataToRight(key TKey, val TVal) *KeyNode[TKey, TVal] {
	return kn.AppendToRight(NewKeyNode(key, val))
}

func (kn *KeyNode[TKey, TVal]) Walk(c chan<- KeyData[TKey, TVal], rev bool) {
	go func() {
		c <- *kn.KeyData
		if rev {
			for kn.left != nil {
				kn = kn.left
				c <- *kn.KeyData
			}
		} else {
			for kn.right != nil {
				kn = kn.right
				c <- *kn.KeyData
			}
		}
		close(c)
	}()
}
