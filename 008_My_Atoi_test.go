package leetcode_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMyAtio(t *testing.T) {
	type testCase struct {
		name string
		s    string
		d    int
	}

	tcs := []testCase{
		{
			name: "Case_1_Pos_Num_w/o_Leading_Space",
			s:    "123",
			d:    123,
		},
		{
			name: "Case_2_Neg_Num_w_Leading_Space",
			s:    "   -124",
			d:    -124,
		},
		{
			name: "Case_3_Pos_Num_w_Words",
			s:    "123 other-words",
			d:    123,
		},
		{
			name: "Leading_Words",
			s:    "other-words 123",
			d:    0,
		},
		{
			name: "POS overflow",
			s:    "+2147483649",
			d:    math.MaxInt32,
		},
		{
			name: "Min overflow",
			s:    "-2147483649",
			d:    math.MinInt32,
		},
		{
			name: "No Digit",
			s:    "-avd",
			d:    0,
		},
		{
			name: "Empty",
			s:    "",
			d:    0,
		},
		{
			name: "Only Plus",
			s:    "+",
			d:    0,
		},
		{
			name: "Only Minus",
			s:    "-",
			d:    0,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.d, leetcode.MyAtoi(tc.s))
			},
		)
	}
}
