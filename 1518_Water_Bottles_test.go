package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ1518(t *testing.T) {
	q := leetcode.Q1518{}

	tcs := []struct {
		numBottles  int
		numExchange int
		answer      int
	}{
		{9, 3, 13},
		{15, 4, 19},
		{84, 13, 90},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("numBottles:%d numExchange: %d", tc.numBottles, tc.numExchange),
			func(t *testing.T) {
				require.Equal(t, tc.answer, q.NumWaterBottles(tc.numBottles, tc.numExchange))
			},
		)
	}
}
