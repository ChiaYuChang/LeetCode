package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSuccessfulPairs(t *testing.T) {
	type testCase struct {
		spells  []int
		potion  []int
		success int64
		answer  []int
	}

	tcs := []testCase{
		{
			spells:  []int{5, 1, 3},
			potion:  []int{1, 2, 3, 4, 5},
			success: 7,
			answer:  []int{4, 0, 3},
		},
		{
			spells:  []int{3, 1, 2},
			potion:  []int{8, 5, 8},
			success: 16,
			answer:  []int{2, 0, 2},
		},
	}

	q := leetcode.Q2300{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.answer, q.SuccessfulPairs(tc.spells, tc.potion, tc.success))
			},
		)
	}
}
