package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestDecodings(t *testing.T) {
	type testCase struct {
		s string
		n int
	}

	tcs := []testCase{
		{"12", 2},
		{"226", 3},
		{"06", 0},
		{"1132610137", 12},
		{"110010", 0},
		{"100", 0},
		{"2611055971756562", 4},
		{"1", 1},
	}

	q := leetcode.Q0091{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.s),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.n, q.NumDecodings(tc.s))
			},
		)
	}
}
