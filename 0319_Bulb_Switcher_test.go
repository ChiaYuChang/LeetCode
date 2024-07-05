package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestBulbSwitch(t *testing.T) {
	type testCase struct {
		num int
		ans int
	}

	tcs := []testCase{
		{3, 1},
		{0, 0},
		{1, 1},
	}

	q := leetcode.Q0319{}
	for i, tc := range tcs {
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.BulbSwitch(tc.num))
			},
		)
	}
}
