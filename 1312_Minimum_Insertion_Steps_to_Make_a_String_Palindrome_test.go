package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinInsertions(t *testing.T) {
	type testCase struct {
		s string
		a int
	}

	tcs := []testCase{
		{
			s: "zzazz",
			a: 0,
		},
		{
			s: "aaba",
			a: 1,
		},
		{
			s: "mbadm",
			a: 2,
		},
		{
			s: "leetcode",
			a: 5,
		},
		{
			s: "abc",
			a: 2,
		},
		{
			s: "a",
			a: 0,
		},
	}

	q := leetcode.Q1312{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.a, q.MinInsertions(tc.s))
			},
		)
	}
}
