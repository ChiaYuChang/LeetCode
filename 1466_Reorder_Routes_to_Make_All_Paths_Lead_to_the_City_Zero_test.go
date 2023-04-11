package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinReorder(t *testing.T) {
	type testCase struct {
		n   int
		con [][]int
		ans int
	}

	tcs := []testCase{
		{
			n:   6,
			con: [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}},
			ans: 3,
		},
		{
			n:   5,
			con: [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}},
			ans: 2,
		},
		{
			n:   5,
			con: [][]int{{4, 3}, {2, 3}, {1, 2}, {1, 0}},
			ans: 2,
		},
	}

	q := leetcode.Q1466{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.MinReorder(tc.n, tc.con))
			},
		)
	}

}
