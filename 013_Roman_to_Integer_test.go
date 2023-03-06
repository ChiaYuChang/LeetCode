package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestRomaToInt(t *testing.T) {
	type testCase struct {
		name string
		ans  int
		s    string
	}

	tcs := []testCase{
		{
			name: "Case_1",
			ans:  58,
			s:    "LVIII",
		},
		{
			name: "Case_2",
			ans:  1994,
			s:    "MCMXCIV",
		},
		{
			name: "3724",
			s:    "MMMDCCXXIV",
			ans:  3724,
		},
		{
			name: "Zero",
			s:    "",
			ans:  0,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(t, tc.ans, leetcode.RomanToInt(tc.s))
			},
		)
	}
}
