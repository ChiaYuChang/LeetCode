package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestPossibleBipartition(t *testing.T) {
	type testCase struct {
		n        int
		dislikes [][]int
		answer   bool
	}

	tcs := []testCase{
		{
			n:        4,
			dislikes: [][]int{{1, 2}, {1, 3}, {2, 4}},
			answer:   true,
		},
		{
			n:        3,
			dislikes: [][]int{{1, 2}, {1, 3}, {2, 3}},
			answer:   false,
		},
	}

	q := leetcode.Q0886{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.answer, q.PossibleBipartition(tc.n, tc.dislikes))
			},
		)
	}
}
