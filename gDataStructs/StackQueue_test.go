package gdatastructs_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	gdatastructs "gitlab.com/gjerry134679/leetcode/gDataStructs"
)

type tAction int8

type stackQueueAction[T any] struct {
	Action tAction
	Data   T
	Check  func(t *testing.T, ok bool, data any)
}

const (
	actCheckIsEmpty tAction = iota + 1
	actCheckLength
	actCheckHeadData
	actCheckTailData
	actPush
	actPop
	actEnqueue
	actDequeue
	actToArray
)

func (tA tAction) String() string {
	switch tA {
	case actCheckIsEmpty:
		return "IsEmpry"
	case actCheckLength:
		return "Length"
	case actPush:
		return "Push"
	case actPop:
		return "Pop"
	case actEnqueue:
		return "Enqueue"
	case actDequeue:
		return "Dequeue"
	case actToArray:
		return "ToArray"
	case actCheckHeadData:
		return "CheckHeadData"
	case actCheckTailData:
		return "CheckTailData"
	default:
		return "Unknown Action"
	}
}

func do[T any](t *testing.T, sq *gdatastructs.StackQueue[T], actions ...stackQueueAction[T]) {
	for i, action := range actions {
		t.Logf("action %3d %v", i, action.Action)
		switch action.Action {
		case actCheckIsEmpty:
			isEmpty := sq.IsEmpty()
			action.Check(t, true, isEmpty)
		case actCheckLength:
			sqLen := sq.Length
			action.Check(t, true, sqLen)
		case actCheckHeadData:
			if sq.Head != nil {
				hdata := *sq.Head.Data()
				action.Check(t, true, hdata)
			} else {
				action.Check(t, false, nil)
			}
		case actCheckTailData:
			if sq.Head != nil {
				tdata := *sq.Tail.Data()
				action.Check(t, true, tdata)
			} else {
				action.Check(t, false, nil)
			}
		case actPush:
			sq.Push(action.Data)
			if action.Check != nil {
				action.Check(t, true, *sq.Head.Data())
			}
		case actPop:
			d, ok := sq.Pop()
			if action.Check != nil {
				action.Check(t, ok, d)
			}
		case actEnqueue:
			sq.Enqueue(action.Data)
			if action.Check != nil {
				action.Check(t, true, *sq.Tail.Data())
			}
		case actDequeue:
			d, ok := sq.Dequeue()
			if action.Check != nil {
				action.Check(t, ok, d)
			}
		case actToArray:
			ds := sq.ToArray()
			action.Check(t, true, ds)
		default:
			t.Fatal("unknown stackqueue action")
		}
	}
}

func TestNewStackQueue(t *testing.T) {
	sq := gdatastructs.NewStackQueue[int]()
	require.NotNil(t, sq)
	require.True(t, sq.IsEmpty())
}

func TestEmptyStackQueue(t *testing.T) {
	sq1 := gdatastructs.NewStackQueue[int]()
	val, ok := sq1.Peek("head")
	require.False(t, ok)
	require.Zero(t, val)
	require.True(t, sq1.IsEmpty())

	val, ok = sq1.Pop()
	require.False(t, ok)
	require.Zero(t, val)

	val, ok = sq1.Dequeue()
	require.False(t, ok)
	require.Zero(t, val)

	array := sq1.ToArray()
	require.Len(t, array, 0)
	require.Nil(t, array)

	sq2 := gdatastructs.NewStackQueue[int]()
	sq2.Push(1)
	sq2.Pop()

	val, ok = sq2.Peek("head")
	require.False(t, ok)
	require.Zero(t, val)
	require.True(t, sq2.IsEmpty())

	val, ok = sq2.Pop()
	require.False(t, ok)
	require.Zero(t, val)

	val, ok = sq2.Dequeue()
	require.False(t, ok)
	require.Zero(t, val)

	array = sq2.ToArray()
	require.Len(t, array, 0)
	require.Nil(t, array)

	sq3 := gdatastructs.NewStackQueue[int]()
	sq3.Enqueue(1)
	sq3.Dequeue()
	val, ok = sq3.Peek("head")
	require.False(t, ok)
	require.Zero(t, val)
	require.True(t, sq3.IsEmpty())

	val, ok = sq3.Pop()
	require.False(t, ok)
	require.Zero(t, val)

	val, ok = sq3.Dequeue()
	require.False(t, ok)
	require.Zero(t, val)

	array = sq3.ToArray()
	require.Len(t, array, 0)
	require.Nil(t, array)
}

func TestStackQueuePush(t *testing.T) {
	sq := gdatastructs.NewStackQueue[int]()
	n := 10
	acts := make([]stackQueueAction[int], n+2)
	for i := 0; i < n; i++ {
		expData := i
		acts[i] = stackQueueAction[int]{
			Action: actPush,
			Data:   i,
			Check: func(t *testing.T, ok bool, data any) {
				getData, ok := data.(int)
				require.True(t, ok)
				require.Equal(t, expData, getData)
			},
		}
	}

	acts[n] = stackQueueAction[int]{
		actCheckLength,
		0,
		func(t *testing.T, ok bool, data any) {
			l, ok := data.(int)
			require.True(t, ok)
			require.Equal(t, n, l)
		},
	}

	acts[n+1] = stackQueueAction[int]{
		actToArray,
		0,
		func(t *testing.T, ok bool, data any) {
			a, ok := data.([]int)
			require.True(t, ok)
			require.Equal(t, n, len(a))
		},
	}
	do(t, sq, acts...)

	a := sq.ToArray()
	t.Logf("a: %v", a)

	require.Equal(t, n-1, *sq.Head.Data())
	require.Equal(t, 0, *sq.Tail.Data())
	require.False(t, sq.Head.HasPrev())
	require.False(t, sq.Tail.HasNext())

	c := make(chan int)
	sq.Walk(c, false)
	i := 0
	for j := range c {
		require.Equal(t, n-i-1, j)
		i++
	}
	t.Log(sq)
}

func TestStackQueuePop(t *testing.T) {
	sq := gdatastructs.NewStackQueue[int]()
	n := 10

	for i := 0; i < n; i++ {
		sq.Push(i)
	}
	require.Equal(t, n, sq.Length)

	actions := make([]stackQueueAction[int], 2*n)
	j := 0
	// pop all elements in the sq
	for i := n - 1; i >= 0; i-- {
		expData := i
		actions[j] = stackQueueAction[int]{
			Action: actPop,
			Data:   0,
			Check: func(t *testing.T, ok bool, data any) {
				require.True(t, ok)
				getData, ok := data.(int)
				require.True(t, ok)
				require.Equal(t, expData, getData)
			},
		}
		j += 1
		expLen := i
		actions[j] = stackQueueAction[int]{
			Action: actCheckLength,
			Data:   0,
			Check: func(t *testing.T, ok bool, data any) {
				getLen, ok := data.(int)
				require.True(t, ok)
				require.Equal(t, expLen, getLen)
			},
		}
		j += 1
	}
	actions = append(actions,
		// length of sq should be zero
		stackQueueAction[int]{
			Action: actCheckIsEmpty,
			Data:   0,
			Check: func(t *testing.T, ok bool, data any) {
				getIsEmpty, ok := data.(bool)
				require.True(t, ok)
				require.True(t, getIsEmpty)
			},
		},
		// pop should return zero value of type T when sq is empty
		stackQueueAction[int]{
			Action: actPop,
			Data:   0,
			Check: func(t *testing.T, ok bool, data any) {
				require.False(t, ok)
				getData, ok := data.(int)
				require.True(t, ok)
				require.Zero(t, getData)
			},
		},
		// tail data should be empty when sq is empty
		stackQueueAction[int]{
			Action: actCheckTailData,
			Data:   0,
			Check: func(t *testing.T, ok bool, data any) {
				require.False(t, ok)
			},
		},
		// head data should be empty when sq is empty
		stackQueueAction[int]{
			Action: actCheckHeadData,
			Data:   0,
			Check: func(t *testing.T, ok bool, data any) {
				require.False(t, ok)
			},
		})

	do(t, sq, actions...)
}

func TestArray2StackQueue(t *testing.T) {
	n := 10
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Int()
	}

	sq := gdatastructs.ArrayToStackQueue(ints, false)

	require.NotNil(t, sq)
	require.Equal(t, sq.Length, n)

	for i := 0; sq.Length > 0; i++ {
		j, ok := sq.Pop()
		require.True(t, ok)
		require.Equal(t, ints[i], j)
	}
}

func getNewPushAction[T any](data T) stackQueueAction[T] {
	expData := data
	return stackQueueAction[T]{
		Action: actPush,
		Data:   data,
		Check: func(t *testing.T, ok bool, data any) {
			require.True(t, ok)
			getData, ok := data.(T)
			require.True(t, ok)
			require.Equal(t, expData, getData)
		},
	}
}

func getNewEnqueueAction[T any](data T) stackQueueAction[T] {
	expData := data
	return stackQueueAction[T]{
		Action: actEnqueue,
		Data:   data,
		Check: func(t *testing.T, ok bool, data any) {
			require.True(t, ok)
			getData, ok := data.(T)
			require.True(t, ok)
			require.Equal(t, expData, getData)
		},
	}
}

func getNewPopAction[T any](expData T) stackQueueAction[T] {
	return stackQueueAction[T]{
		Action: actPop,
		Data:   *new(T),
		Check: func(t *testing.T, ok bool, data any) {
			require.True(t, ok)
			getData, ok := data.(T)
			require.True(t, ok)
			require.Equal(t, expData, getData)
		},
	}
}

func getNewDequeueAction[T any](expData T) stackQueueAction[T] {
	a := getNewPopAction(expData)
	a.Action = actDequeue
	return a
}

func getNewCheckLengthAction[T any](expLen int) stackQueueAction[T] {
	return stackQueueAction[T]{
		Action: actCheckLength,
		Data:   *new(T),
		Check: func(t *testing.T, ok bool, data any) {
			require.True(t, ok)
			getLen, ok := data.(int)
			require.True(t, ok)
			require.Equal(t, expLen, getLen)
		},
	}
}

// Test Push and Pop of an Stackqueue
func TestStack(t *testing.T) {
	nPush := 200
	nPop := 100
	require.Greater(t, nPush, nPop)

	nums := make([]int, nPush)
	for i := range nums {
		nums[i] = i
	}

	actions := make([]stackQueueAction[int], nPush+nPop)

	sqLen := 0
	nIdx := 0
	sIdx := 0
	simpleStack := make([]int, nPush+nPop)

	// t.Logf("ss(%d): %v", sqLen, simpleStack)
	// t.Logf("lenght of actions: %d", len(actions))

	for k := range actions {
		// 0 for push
		// 1 for pop
		var a tAction
		if sqLen > 0 {
			x := rand.Float64()
			if x < 0.5 {
				if nPush < 1 {
					a = actPop
					nPop--
				} else {
					a = actPush
					nPush--
				}
			} else if x < 0.95 {
				if nPop < 1 {
					a = actPush
					nPush--
				} else {
					a = actPop
					nPop--
				}
			} else {
				a = actCheckLength
			}
		} else {
			a = actPush
			nPush--
		}

		switch a {
		case actPush:
			// t.Logf("> %d Push   %d", k, nums[nIdx])
			actions[k] = getNewPushAction(nums[nIdx])
			simpleStack[sIdx] = nums[nIdx]
			nIdx++
			sIdx++
			sqLen++
		case actPop:
			// t.Logf("> %d Pop    %d", k, simpleStack[sIdx-1])
			actions[k] = getNewPopAction(simpleStack[sIdx-1])
			// simpleStack[sIdx-1] = -1
			sIdx--
			sqLen--
		case actCheckLength:
			// t.Logf("> %d Check  %d", k, sqLen)
			actions[k] = getNewCheckLengthAction[int](sqLen)
		default:
			t.Fatalf("not supported action %d", a)
		}
		// t.Logf("ss(%d): %v", sqLen, simpleStack)
	}

	stack := gdatastructs.NewStackQueue[int]()
	do(t, stack, actions...)
}

func TestQueue(t *testing.T) {
	nEnqueue := 10
	nDequeue := 6
	require.Greater(t, nEnqueue, nDequeue)

	nums := make([]int, nEnqueue)
	for i := range nums {
		nums[i] = i
	}

	actions := make([]stackQueueAction[int], nEnqueue+nDequeue)

	sqLen := 0
	nIdx := 0
	sIdx := 0
	simpleQueue := make([]int, nEnqueue+nDequeue)
	// for i := range simpleQueue {
	// 	simpleQueue[i] = -1
	// }

	var toQueueIdx func(i int) int = func(i int) int {
		return len(simpleQueue) - 1 - i
	}

	// t.Logf("ss(%d): %v", sqLen, simpleStack)
	// t.Logf("lenght of actions: %d", len(actions))
	for k := range actions {
		// 0 for push
		// 1 for pop
		var a tAction
		if sqLen > 0 {
			x := rand.Float64()
			if x < 0.5 {
				if nEnqueue < 1 {
					a = actDequeue
					nDequeue--
				} else {
					a = actEnqueue
					nEnqueue--
				}
			} else if x < 0.95 {
				if nDequeue < 1 {
					a = actEnqueue
					nEnqueue--
				} else {
					a = actDequeue
					nDequeue--
				}
			} else {
				a = actCheckLength
			}
		} else {
			a = actEnqueue
			nEnqueue--
		}

		// t.Logf("nEnqueue: %d nDequeue %d", nEnqueue, nDequeue)
		switch a {
		case actEnqueue:
			// t.Logf("> %d Enqueue %d", k, nums[nIdx])
			actions[k] = getNewEnqueueAction(nums[nIdx])
			simpleQueue[toQueueIdx(sIdx)] = nums[nIdx]
			nIdx++
			sIdx++
			sqLen++
		case actDequeue:
			// t.Logf("> %d Dequeue %d", k, simpleQueue[toQueueIdx(sIdx-1)])
			actions[k] = getNewDequeueAction(simpleQueue[toQueueIdx(sIdx-1)])
			simpleQueue[toQueueIdx(sIdx-1)] = -1
			sIdx--
			sqLen--
		case actCheckLength:
			// t.Logf("> %d Check   %d", k, sqLen)
			actions[k] = getNewCheckLengthAction[int](sqLen)
		default:
			t.Fatalf("not supported action %d", a)
		}
		// t.Logf("ss(%d): %v", sqLen, simpleQueue)
	}
	// queue := gdatastructs.NewStackQueue[int]()
	// do(t, queue, actions...)
}
