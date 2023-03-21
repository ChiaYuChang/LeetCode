package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestHammingWeight(t *testing.T) {
	type testCase struct {
		name string
		num  uint32
		ans  int
	}

	tcs := []testCase{
		{
			name: "Case_1",
			num:  0b00000010100101000001111010011100,
			ans:  12,
		},
		{
			name: "Case_2",
			num:  0b11111111111111111111111111111101,
			ans:  31,
		},
		{
			name: "One",
			num:  0b00000000000000000000000000000001,
			ans:  1,
		},
		{
			name: "Zero",
			num:  0b00000000000000000000000000000000,
			ans:  0,
		},
	}

	q := leetcode.Q0191{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.HammingWeight(tc.num))
			},
		)
	}
}
