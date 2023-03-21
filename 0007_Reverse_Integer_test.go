package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestReverseInteger(t *testing.T) {
	type testCase struct {
		name string
		x    int
		a    int
	}

	tcs := []testCase{
		{
			name: "Case_1_Pos_Int",
			x:    123,
			a:    321,
		},
		{
			name: "Case_2_Neg_Int",
			x:    -123,
			a:    -321,
		},
		{
			name: "Case_3_Ending_Zero",
			x:    120,
			a:    21,
		},
		{
			name: "Pos_Overflow",
			x:    1534236469,
			a:    0,
		},
		{
			name: "Neg_Overflow",
			x:    -1563847412,
			a:    0,
		},
		{
			name: "Zero",
			x:    0,
			a:    0,
		},
	}

	q := leetcode.Q0007{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.a, q.ReverseInt(tc.x))
			},
		)
	}
}
