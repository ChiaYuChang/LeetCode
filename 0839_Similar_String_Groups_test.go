package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestNumSimilarGroups(t *testing.T) {
	type testCase struct {
		strs []string
		ans  int
	}

	tcs := []testCase{
		{[]string{"tars", "rats", "arts", "star"}, 2},
		{[]string{"omv", "ovm"}, 1},
		{[]string{"abc", "abc"}, 1},
	}

	q := leetcode.Q0839{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.NumSimilarGroups(tc.strs))
			},
		)
	}
}
