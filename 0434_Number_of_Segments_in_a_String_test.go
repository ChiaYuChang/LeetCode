package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestCountSegments(t *testing.T) {
	type testCase struct {
		s string
		n int
	}

	tcs := []testCase{
		{
			s: "Hello, my name is John",
			n: 5,
		},
		{
			s: "Hello",
			n: 1,
		},
		{
			s: "a, b, c",
			n: 3,
		},
	}

	q := leetcode.Q0434{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.s),
			func(t *testing.T) {
				require.Equal(t, tc.n, q.CountSegments(tc.s))
				require.Equal(t, tc.n, q.CountSegmentsPadding(tc.s))
			},
		)
	}
}
