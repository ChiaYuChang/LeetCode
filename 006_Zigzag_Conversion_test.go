package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestZogzagConv(t *testing.T) {
	type testCase struct {
		name    string
		s       string
		numRows int
		ans     string
	}

	tcs := []testCase{
		{
			name:    "Case_1",
			s:       "PAYPALISHIRING",
			numRows: 3,
			ans:     "PAHNAPLSIIGYIR",
		},
		{
			name:    "Case_2",
			s:       "PAYPALISHIRING",
			numRows: 4,
			ans:     "PINALSIGYAHRPI",
		},
		{
			name:    "Case_3",
			s:       "A",
			numRows: 1,
			ans:     "A",
		},
		{
			name:    "Empty string",
			s:       "",
			numRows: 100,
			ans:     "",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				ans := leetcode.ZigzagConversion(tc.s, tc.numRows)
				require.Equal(t, tc.ans, ans)

			},
		)
	}
}
