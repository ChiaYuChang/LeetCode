package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMatrixBlockSum(t *testing.T) {
	type testCase struct {
		mat [][]int
		k   int
		ans [][]int
	}

	tcs := []testCase{
		{
			mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:   0,
			ans: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:   1,
			ans: [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}},
		},
		{
			mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:   2,
			ans: [][]int{{45, 45, 45}, {45, 45, 45}, {45, 45, 45}},
		},
	}

	q := leetcode.Q1314{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d k=%d", i+1, tc.k),
			func(t *testing.T) {
				require.Equal(
					t, tc.ans,
					q.MatrixBlockSum(tc.mat, tc.k),
				)
			},
		)
	}
}
