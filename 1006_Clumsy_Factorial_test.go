package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestClumsyFactorial(t *testing.T) {
	type testCase struct {
		num int
		ans int
	}

	tcs := []testCase{
		{4, 7},
		{10, 12},
		{100, 101},
	}

	q := leetcode.Q1006{}
	for i, tc := range tcs {
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.Clumsy(tc.num))
			},
		)
	}
}
