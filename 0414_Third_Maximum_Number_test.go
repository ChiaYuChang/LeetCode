package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestTriplet(t *testing.T) {
	triplet := leetcode.NewQ414Triplet()

	triplet.Append(0)
	t.Log(triplet)
	require.Equal(t, 0, triplet.Top3())

	triplet.Append(3)
	t.Log(triplet)
	require.Equal(t, 3, triplet.Top3())

	triplet.Append(-1)
	t.Log(triplet)
	require.Equal(t, -1, triplet.Top3())

	triplet.Append(2)
	t.Log(triplet)
	require.Equal(t, 0, triplet.Top3())

	triplet.Append(4)
	t.Log(triplet)
	require.Equal(t, 2, triplet.Top3())
}

func TestThridMax(t *testing.T) {
	type testCase struct {
		nums []int
		ans  int
	}

	tcs := []testCase{
		{
			nums: []int{3, 1, 2},
			ans:  1,
		},
		{
			nums: []int{1, 2},
			ans:  2,
		},
		{
			nums: []int{1, 3, 2, 2},
			ans:  1,
		},
	}

	q := leetcode.Q0414{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Test Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.ThirdMax(tc.nums))
			})
	}
}
