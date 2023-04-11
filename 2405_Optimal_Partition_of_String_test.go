package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestPartitionString(t *testing.T) {
	type testCase struct {
		str string
		ans int
	}

	tcs := []testCase{
		{"abacaba", 4},
		{"ssssss", 6},
	}

	q := leetcode.Q2405{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.PartitionString(tc.str))
			},
		)
	}
}
