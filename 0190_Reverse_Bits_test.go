package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestReverseBits(t *testing.T) {
	type testCase struct {
		name string
		num  uint32
		ans  uint32
	}

	tcs := []testCase{
		{
			name: "Case_1",
			num:  0b00000010100101000001111010011100,
			ans:  0b00111001011110000010100101000000,
		},
		{
			name: "Case_2",
			num:  0b11111111111111111111111111111101,
			ans:  0b10111111111111111111111111111111,
		},
		{
			name: "One",
			num:  0b00000000000000000000000000000001,
			ans:  0b10000000000000000000000000000000,
		},
		{
			name: "Zero",
			num:  0b00000000000000000000000000000000,
			ans:  0b00000000000000000000000000000000,
		},
	}

	q := leetcode.Q0190{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.ReverseBits(tc.num))
			},
		)
	}
}
