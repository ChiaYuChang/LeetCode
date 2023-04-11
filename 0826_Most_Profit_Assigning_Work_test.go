package leetcode_test

import (
	"testing"

	"gitlab.com/gjerry134679/leetcode"
)

func TestMaxProfitAssignment(t *testing.T) {
	type testCase struct {
		difficulty []int
		profit     []int
		worker     []int
		maxProfit  int
	}

	tcs := []testCase{
		{
			difficulty: []int{2, 4, 6, 8, 10},
			profit:     []int{10, 20, 30, 40, 50},
			worker:     []int{4, 5, 6, 7},
			maxProfit:  100,
		},
		{
			difficulty: []int{85, 47, 57},
			profit:     []int{24, 66, 99},
			worker:     []int{40, 25, 25},
			maxProfit:  0,
		},
		{
			difficulty: []int{10, 30, 20},
			profit:     []int{10, 50, 30},
			worker:     []int{30, 32, 21, 9},
			maxProfit:  130,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Logf("want: %d, get: %d",
			tc.maxProfit,
			leetcode.MaxProfitAssignment(tc.difficulty, tc.profit, tc.worker),
		)
	}
}
