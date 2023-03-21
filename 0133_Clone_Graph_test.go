package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestNumFriendRequest(t *testing.T) {
	type testCase struct {
		age []int
		ans int
	}

	tcs := []testCase{
		{
			age: []int{16, 16},
			ans: 2,
		},
		{
			age: []int{16, 17, 18},
			ans: 2,
		},
		{
			age: []int{20, 30, 100, 110, 120},
			ans: 3,
		},
		{
			age: []int{
				63, 52, 49, 26, 74, 59, 36, 19, 83, 30, 40,
				50, 110, 24, 69, 65, 87, 48, 34, 101, 112, 23,
				48, 78, 51, 29, 93, 43, 19, 106, 84, 80, 54,
				116, 4, 93, 44, 108, 20, 11, 77, 108, 11, 20,
				24, 75, 69, 114, 115, 27},
			ans: 480,
		},
	}

	q := leetcode.Q0823{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.NumFriendRequests(tc.age))
			},
		)
	}
}
