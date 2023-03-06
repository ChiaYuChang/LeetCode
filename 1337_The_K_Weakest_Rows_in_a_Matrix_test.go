package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFirstZeroIndex(t *testing.T) {
	type testCase struct {
		name string
		ary  []int
		ans  int
	}

	tcs := []testCase{
		{
			name: "General Case 1",
			ary:  []int{1, 0},
			ans:  1,
		},
		{
			name: "General Case 1",
			ary:  []int{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			ans:  4,
		},
		{
			name: "w/o zeros",
			ary:  []int{1, 1, 1, 1, 1},
			ans:  5,
		},
		{
			name: "w/o ones",
			ary:  []int{0},
			ans:  0,
		},
	}

	q1337 := leetcode.Q1337{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(t, tc.ans, q1337.FirstZeroIndex(tc.ary))
			},
		)
	}
}

func TestKWeakestRows(t *testing.T) {
	type testCase struct {
		name string
		mat  [][]int
		k    int
		ans  []int
	}

	tcs := []testCase{
		{
			name: "Case 1",
			mat: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			k:   3,
			ans: []int{2, 0, 3},
		},
		{
			name: "Case 2",
			mat: [][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			},
			k:   2,
			ans: []int{0, 2},
		},
	}

	q1337 := leetcode.Q1337{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.ElementsMatch(t, tc.ans, q1337.KWeakestRows(tc.mat, tc.k))
			},
		)
	}
}
