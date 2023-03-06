package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestAddBinary(t *testing.T) {
	type testCase struct {
		name string
		x    string
		y    string
		ans  string
	}

	tcs := []testCase{
		{
			name: "Case_1",
			x:    "11",
			y:    "1",
			ans:  "100",
		},
		{
			name: "Case_2",
			x:    "1010",
			y:    "1011",
			ans:  "10101",
		},
		{
			name: "Case_3",
			x:    "1",
			y:    "111",
			ans:  "1000",
		},
		{
			name: "Case_4",
			x:    "1010101011101",
			y:    "110011011",
			ans:  "1011011111000",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(
					t, tc.ans,
					leetcode.AddBinary(
						tc.x,
						tc.y,
					))
			},
		)
	}
}
