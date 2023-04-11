package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestHammingDistance(t *testing.T) {
	type testCase struct {
		x, y int
		a    int
	}

	tcs := []testCase{
		{x: 1, y: 4, a: 2},
		{x: 3, y: 1, a: 1},
	}

	q := leetcode.Q0461{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%d/%d", i+1, tc.x, tc.y),
			func(t *testing.T) {
				require.Equal(t, tc.a, q.HammingDistance(tc.x, tc.y))
			},
		)
	}
}
