package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestValidateStackSequences(t *testing.T) {
	type testCase struct {
		pushed []int
		popped []int
		answer bool
	}

	tcs := []testCase{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 5, 3, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 3, 5, 1, 2},
			false,
		},
		{
			[]int{0, 2, 1},
			[]int{0, 1, 2},
			true,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Cast %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.answer, leetcode.ValidateStackSequences(tc.pushed, tc.popped))
			},
		)
	}
}
