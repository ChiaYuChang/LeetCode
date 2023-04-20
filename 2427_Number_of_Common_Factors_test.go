package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestCommonFactors(t *testing.T) {
	type testCase struct {
		a, b int
		ans  int
	}

	tcs := []testCase{
		{12, 6, 4},
		{25, 30, 2},
	}

	q := leetcode.Q2427{}
	for i, tc := range tcs {
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.CommonFactors(tc.a, tc.b))
			},
		)
	}
}
