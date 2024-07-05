package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestTicTacToe(t *testing.T) {
	type testCase struct {
		board   []string
		isValid bool
	}

	tcs := []testCase{
		{[]string{"O  ", "   ", "   "}, false},
		{[]string{"XOX", " X ", "   "}, false},
		{[]string{"XXX", "XOO", "XOO"}, true},
		{[]string{"XOX", "OOO", "XOX"}, false},
		{[]string{"XXX", "   ", "OOO"}, false},
		{[]string{"XOX", "O O", "XOX"}, true},
	}

	q := leetcode.Q0794{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.isValid,
					q.ValidTicTacToe(tc.board),
				)
			},
		)
	}
}
