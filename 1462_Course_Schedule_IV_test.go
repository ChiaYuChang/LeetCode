package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestBooleanMatrix(t *testing.T) {
	var m, mc leetcode.BooleanMatrix
	var ok bool

	m = leetcode.NewBooleanMatrix(3, 3)
	m.SetDiagonal(true)
	m.Set(1, 2, true)
	m.Set(1, 0, true)
	m.Set(2, 0, true)

	mc, ok = m.ToConverge(100)
	require.True(t, ok)
	t.Log(mc)

	m = leetcode.NewBooleanMatrix(2, 2)
	m.SetDiagonal(true)
	m.Set(1, 0, true)

	mc, ok = m.ToConverge(100)
	require.True(t, ok)
	t.Log(mc)

	m = leetcode.NewBooleanMatrix(5, 5)
	m.SetDiagonal(true)
	m.Set(1, 0, true)
	m.Set(1, 2, true)
	m.Set(3, 4, true)
	m.Set(4, 2, true)
	m.Set(0, 3, true)
	t.Log(m)
}

func TestCheckIfPrerequisite(t *testing.T) {
	type testCase struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
		answer        []bool
	}

	tcs := []testCase{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			queries:       [][]int{{0, 1}, {1, 0}},
			answer:        []bool{false, true},
		},
		{
			numCourses:    2,
			prerequisites: [][]int{},
			queries:       [][]int{{1, 0}, {0, 1}},
			answer:        []bool{false, false},
		},
		{
			numCourses:    3,
			prerequisites: [][]int{{1, 2}, {1, 0}, {2, 0}},
			queries:       [][]int{{1, 0}, {1, 2}},
			answer:        []bool{true, true},
		},
		{
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {2, 0}},
			queries:       [][]int{{0, 1}, {2, 0}},
			answer:        []bool{false, true},
		},
		{
			numCourses:    5,
			prerequisites: [][]int{{1, 0}, {1, 2}, {3, 4}, {4, 2}, {0, 3}},
			queries:       [][]int{{1, 3}, {0, 4}, {3, 2}, {2, 3}},
			answer:        []bool{true, true, true, false},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{2, 3}, {2, 1}, {0, 3}, {0, 1}},
			queries:       [][]int{{0, 1}, {0, 3}, {2, 3}, {3, 0}, {2, 0}, {0, 2}},
			answer:        []bool{true, true, true, false, false, false},
		},
		{
			numCourses: 7,
			prerequisites: [][]int{
				{2, 3}, {2, 1}, {2, 0}, {3, 4}, {3, 6},
				{5, 1}, {5, 0}, {1, 4}, {1, 0}, {4, 0},
				{0, 6}},
			queries: [][]int{
				{3, 0}, {6, 4}, {5, 6}, {2, 6}, {2, 3},
				{5, 6}, {4, 0}, {2, 6}, {3, 5}, {5, 3},
				{1, 6}, {1, 0}, {3, 5}, {6, 5}, {2, 3},
				{3, 0}, {3, 4}, {3, 4}, {2, 5}, {0, 3},
				{4, 0}, {6, 4}, {5, 0}, {6, 5}, {5, 6},
				{6, 5}, {1, 0}, {3, 4}, {1, 5}, {1, 4},
				{3, 6}, {0, 1}, {1, 2}, {5, 1}, {5, 3},
				{5, 3}, {3, 4}, {5, 4}, {5, 4}, {5, 3}},
			answer: []bool{
				true, false, true, true, true,
				true, true, true, false, false,
				true, true, false, false, true,
				true, true, true, false, false,
				true, false, true, false, true,
				false, true, true, false, true,
				true, false, false, true, false,
				false, true, true, true, false},
		},
	}

	q := leetcode.Q1462{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				ans := q.CheckIfPrerequisite(tc.numCourses, tc.prerequisites, tc.queries)
				require.Equal(t, len(tc.answer), len(ans))
				for i := 0; i < len(ans); i++ {
					require.Equal(t, tc.answer[i], ans[i])
				}
			},
		)
	}
}
