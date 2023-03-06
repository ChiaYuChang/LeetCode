package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestDecompressRLElist(t *testing.T) {
	type testCase struct {
		nums []int
		ans  []int
	}

	tcs := []testCase{
		{
			nums: []int{1, 2, 3, 4},
			ans:  []int{2, 4, 4, 4},
		},
		{
			nums: []int{1, 1, 2, 3},
			ans:  []int{1, 3, 3},
		},
	}

	q := leetcode.Q1313{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(
					t, tc.ans,
					q.DecompressRLElist(tc.nums),
				)
			},
		)
	}
}
