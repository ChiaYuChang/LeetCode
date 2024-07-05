package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestHeap(t *testing.T) {
	type ActionType int8
	const (
		New ActionType = iota
		Push
		Pop
	)

	type Action struct {
		Type ActionType
		Attr []int
		Ans  []int
	}

	type testCase struct {
		actions []Action
	}

	tcs := []testCase{
		{
			actions: []Action{
				{New, []int{2, 4, 3, 1}, []int{4, 1}},
				{Pop, []int{}, []int{3, 1}},
				{Pop, []int{}, []int{2, 2}},
				{Push, []int{5}, []int{3, 3}},
				{Push, []int{-1}, []int{4, -1}},
				{Pop, []int{}, []int{3, -1}},
				{Pop, []int{}, []int{2, 3}},
				{Pop, []int{}, []int{1, 4}},
				{Pop, []int{}, []int{0, 5}},
				{Push, []int{10}, []int{1, 10}},
				{Push, []int{3}, []int{2, 3}},
				{Push, []int{11}, []int{3, 3}},
			},
		},
	}

	q := leetcode.Q2336{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				var h *leetcode.Q2336MinHeap
				for i, action := range tc.actions {
					t.Logf("Step: %d\n", i)
					switch action.Type {
					case New:
						h = q.NewHeap(action.Attr)
						require.Equal(t, action.Ans[0], h.Len())
						require.Equal(t, action.Ans[1], h.Peek())
					case Push:
						for _, n := range action.Attr {
							h.Push(n)
						}
						require.Equal(t, action.Ans[0], h.Len())
						require.Equal(t, action.Ans[1], h.Peek())
					case Pop:
						n, ok := h.Pop()
						require.Equal(t, action.Ans[0], h.Len())
						require.Equal(t, action.Ans[1], n)
						require.True(t, ok)
						if action.Ans[0] == 0 {
							_, ok := h.Pop()
							require.False(t, ok)
						}
					}
					t.Logf("%v\n", h)
				}
			},
		)
	}
}

func TestSmallestInfiniteSet(t *testing.T) {
	type ActionType int8
	const (
		SmallestInfiniteSet ActionType = iota
		AddBack
		PopSmallest
	)

	type Action struct {
		Type ActionType
		Attr int
		Ans  int
	}

	type testCase struct {
		actions []Action
	}

	tcs := []testCase{
		{
			actions: []Action{
				{SmallestInfiniteSet, 0, 0},
				{AddBack, 2, 0},
				{PopSmallest, 0, 1},
				{PopSmallest, 0, 2},
				{PopSmallest, 0, 3},
				{AddBack, 1, 0},
				{PopSmallest, 0, 1},
				{PopSmallest, 0, 4},
				{PopSmallest, 0, 5},
			},
		},
	}

	q := leetcode.Q2336{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				var smallestInfiniteSet leetcode.Q2336SmallestInfiniteSet
				for i, action := range tc.actions {
					t.Logf("Step: %d\n", i)
					switch action.Type {
					case SmallestInfiniteSet:
						smallestInfiniteSet = q.Constructor()
					case AddBack:
						smallestInfiniteSet.AddBack(action.Attr)
					case PopSmallest:
						require.Equal(t, action.Ans, smallestInfiniteSet.PopSmallest())
					}
				}
			},
		)
	}
}
