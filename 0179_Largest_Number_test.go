package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLargestNumber(t *testing.T) {
	type testCase struct {
		nums   []int
		answer string
	}

	tcs := []testCase{
		{
			nums:   []int{10, 2},
			answer: "210",
		},
		{
			nums:   []int{3, 30, 34, 5, 9},
			answer: "9534330",
		},
		{
			nums:   []int{34323, 3432},
			answer: "343234323",
		},
		{
			nums:   []int{8308, 8308, 830},
			answer: "83088308830",
		},
		{
			nums:   []int{0, 0},
			answer: "0",
		},
		{
			nums:   []int{3, 43, 48, 94, 85, 33, 64, 32, 63, 66},
			answer: "9485666463484333332",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.answer, leetcode.LargestNumber(tc.nums))
			},
		)
	}
}
