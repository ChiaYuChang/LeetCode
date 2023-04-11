package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestIntegerReplacement(t *testing.T) {
	type testCase struct {
		n       int
		minStep int
	}

	tcs := []testCase{
		{8, 3}, {7, 4}, {4, 2}, {1924, 13},
	}

	q := leetcode.Q0397{}
	n := q.NewCache()
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-n=%d", i+1, tc.n),
			func(t *testing.T) {
				// using cache
				require.Equal(t, tc.minStep, n.MinStepToOne(tc.n))
				// not using cache
				require.Equal(t, tc.minStep, q.IntegerReplacement(tc.n))
			},
		)
	}
}
