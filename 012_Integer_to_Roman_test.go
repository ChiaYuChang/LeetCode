package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestIntToRoman(t *testing.T) {
	type testCase struct {
		name string
		num  int
		ans  string
	}

	tcs := []testCase{
		{
			name: "Case_1",
			num:  58,
			ans:  "LVIII",
		},
		{
			name: "Case_2",
			num:  1994,
			ans:  "MCMXCIV",
		},
		{
			name: "3724",
			num:  3724,
			ans:  "MMMDCCXXIV",
		},
		{
			name: "Zero",
			num:  0,
			ans:  "",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(t, tc.ans, leetcode.IntToRoman(tc.num))
			},
		)
	}
}
