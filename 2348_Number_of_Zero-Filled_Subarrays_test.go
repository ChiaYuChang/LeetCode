package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindZeroSegLen(t *testing.T) {
	type testCase struct {
		nums []int
		sl   []int
	}

	tcs := []testCase{
		{
			nums: []int{0, 0, 1, 0, 0, 1},
			sl:   []int{2, 2},
		},
		{
			nums: []int{0, 0, 1, 0, 0, 0},
			sl:   []int{2, 3},
		},
		{
			nums: []int{0, 0, 0, 0, 1},
			sl:   []int{4},
		},
		{
			nums: []int{0, 0, 0, 0, 0},
			sl:   []int{5},
		},
		{
			nums: []int{1, 0, 1, 0, 0, 1},
			sl:   []int{1, 2},
		},
		{
			nums: []int{1, 0, 1, 0, 0, 0},
			sl:   []int{1, 3},
		},
		{
			nums: []int{1, 0, 0, 0, 0, 1},
			sl:   []int{4},
		},
		{
			nums: []int{1, 0, 0, 0, 0, 0},
			sl:   []int{5},
		},
		{
			nums: []int{1, 1, 1, 1, 1, 0},
			sl:   []int{1},
		},
		{
			nums: []int{1, 1, 1, 1, 1, 1},
			sl:   []int{},
		},
	}

	q := leetcode.Q2848{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				sl := q.FindZeroSegLen(tc.nums)
				t.Log(sl)
				require.Equal(t, len(tc.sl), len(sl))
				for j := range tc.sl {
					require.Equal(t, tc.sl[j], sl[j])
				}
			},
		)
	}
}

func TestZeroFilledSubArray(t *testing.T) {
	type testCase struct {
		name string
		nums []int
		ans  int64
	}

	tcs := []testCase{
		{
			name: "not end with zero",
			nums: []int{1, 3, 0, 0, 2, 0, 0, 4},
			ans:  6,
		},
		{
			name: "end with zero",
			nums: []int{0, 0, 0, 2, 0, 0},
			ans:  9,
		},
		{
			name: "three zeros",
			nums: []int{0, 0, 0},
			ans:  6,
		},
		{
			name: "no zeros",
			nums: []int{2, 10, 2019},
			ans:  0,
		},
		{
			name: "empty array",
			nums: []int{},
			ans:  0,
		},
	}

	q := leetcode.Q2848{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.ZeroFilledSubarray(tc.nums))
			},
		)
	}
}
