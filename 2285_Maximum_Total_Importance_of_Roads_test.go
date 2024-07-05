package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMaximumImportance(t *testing.T) {
	tcs := []struct {
		n     int
		roads [][]int
		ans   int64
	}{
		{
			n:     5,
			roads: [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}},
			ans:   43,
		},
		{
			n:     5,
			roads: [][]int{{0, 3}, {2, 4}, {1, 3}},
			ans:   20,
		},
		{
			n:     5,
			roads: [][]int{{0, 1}},
			ans:   9,
		},
	}

	q := leetcode.Q2285(true)
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case_%d", i+1),
			func(t *testing.T) {
				ans := q.MaximumImportance(tc.n, tc.roads)
				require.Equal(t, ans, tc.ans)
			},
		)
	}
}
