package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestKidsWithCandies(t *testing.T) {

	type testCase struct {
		candies      []int
		extraCandies int
		answer       []bool
	}

	tcs := []testCase{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			answer:       []bool{true, true, true, false, true},
		},
		{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			answer:       []bool{true, false, false, false, false},
		},
		{
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			answer:       []bool{true, false, true},
		},
	}

	q := leetcode.Q1431{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.answer, q.KidsWithCandies(tc.candies, tc.extraCandies))
			},
		)
	}

}
