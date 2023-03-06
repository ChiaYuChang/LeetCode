package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMaxArea(t *testing.T) {
	type testCase struct {
		name   string
		height []int
		ans    int
	}

	tcs := []testCase{
		{
			name:   "Case_1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			ans:    49,
		},
		{
			name:   "Case_2",
			height: []int{1, 1},
			ans:    1,
		},
		{
			name:   "Case_3",
			height: []int{2, 3, 10, 5, 7, 8, 9},
			ans:    36,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(t, tc.ans, leetcode.MaxArea(tc.height))
			},
		)
	}
}
