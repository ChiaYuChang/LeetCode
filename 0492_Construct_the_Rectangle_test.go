package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestConstructRectangle(t *testing.T) {
	type testCase struct {
		area int
		lw   []int
	}

	tcs := []testCase{
		{
			area: 4,
			lw:   []int{2, 2},
		},
		{
			area: 37,
			lw:   []int{37, 1},
		},
		{
			area: 122122,
			lw:   []int{427, 286},
		},
	}

	q := leetcode.Q0492{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-Area %d", i+1, tc.area),
			func(t *testing.T) {
				require.Equal(t, tc.lw, q.ConstructRectangle(tc.area))
			},
		)
	}
}
