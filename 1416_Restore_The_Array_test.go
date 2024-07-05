package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestNumberOfArrays(t *testing.T) {
	type testCase struct {
		s string
		k int
		a int
	}

	tcs := []testCase{
		{
			s: "1000",
			k: 10000,
			a: 1,
		},
		{
			s: "1000",
			k: 10,
			a: 0,
		},
		{
			s: "1317",
			k: 2000,
			a: 8,
		},
		{
			s: "1234567890",
			k: 90,
			a: 34,
		},
		{
			s: "600342244431311113256628376226052681059918526204",
			k: 703,
			a: 411743991,
		},
		{
			s: "01",
			k: 100,
			a: 0,
		},
		{
			s: "1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			k: 10000,
			a: 0,
		},
	}

	q := leetcode.Q1416{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.a, q.NumberOfArrays(tc.s, tc.k))
			},
		)

	}
}
