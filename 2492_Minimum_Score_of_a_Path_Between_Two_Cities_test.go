package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinScore(t *testing.T) {
	type testCase struct {
		n     int
		roads [][]int
		ans   int
	}

	tcs := []testCase{
		{
			n: 4,
			roads: [][]int{
				{1, 2, 9},
				{2, 3, 6},
				{2, 4, 5},
				{1, 4, 7},
			},
			ans: 5,
		},
		{
			n: 4,
			roads: [][]int{
				{1, 2, 2},
				{1, 3, 4},
				{3, 4, 7},
			},
			ans: 2,
		},
		{
			n: 20,
			roads: [][]int{
				{18, 20, 9207},
				{14, 12, 1024},
				{11, 9, 3056},
				{8, 19, 416},
				{3, 18, 5898},
				{17, 3, 6779},
				{13, 15, 3539},
				{15, 11, 1451},
				{19, 2, 3805},
				{9, 8, 2238},
				{1, 16, 618},
				{16, 14, 55},
				{17, 7, 6903},
				{12, 13, 1559},
				{2, 17, 3693},
			},
			ans: 55,
		},
	}

	q := leetcode.Q2492{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(
					t, tc.ans,
					q.MinScore(tc.n, tc.roads),
				)
			},
		)
	}
}
