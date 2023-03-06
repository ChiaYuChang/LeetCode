package gdatastructs

import (
	"errors"
	"fmt"
)

type SparseArray[T any] struct {
	length int
	data   map[int]T
}

var ErrIndexOutOfRange = errors.New("index out of range")
var ErrNegativeIndex = errors.New("negative index")
var ErrZeroValue = errors.New("zero value")

func NewSparseArray[T any](length int) *SparseArray[T] {
	return &SparseArray[T]{
		length: length, data: map[int]T{},
	}
}

func (s *SparseArray[T]) SetVal(i int, val T) (*SparseArray[T], error) {
	if i < 0 {
		return s, fmt.Errorf("%w: %d", ErrNegativeIndex, i)
	}

	if i > s.length {
		return s, fmt.Errorf("%w: %d", ErrIndexOutOfRange, s.length)
	}

	s.data[i] = val
	return s, nil
}

func (s *SparseArray[T]) MustSetVal(i int, val T) *SparseArray[T] {
	ns, _ := s.SetVal(i, val)
	return ns
}

func (s *SparseArray[T]) GetVal(i int) (T, error) {
	if i < 0 {
		return *new(T), fmt.Errorf("%w: %d", ErrNegativeIndex, i)
	}

	if i > s.length {
		return *new(T), fmt.Errorf("%w: %d", ErrIndexOutOfRange, s.length)
	}

	val, ok := s.data[i]
	if !ok {
		return *new(T), ErrZeroValue
	}
	return val, nil
}

func (s *SparseArray[T]) MustGetVal(i int) T {
	val, _ := s.GetVal(i)
	return val
}

func (s *SparseArray[T]) ToDense() []T {
	da := make([]T, s.length)
	for i := range da {
		da[i] = s.MustGetVal(i)
	}
	return da
}
