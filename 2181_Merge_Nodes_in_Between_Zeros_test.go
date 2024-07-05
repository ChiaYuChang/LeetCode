package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ1281(t *testing.T) {
	q := leetcode.Q2181{}

	type testCase struct {
		head   []int
		answer []int
	}

	tcs := []testCase{
		{
			head:   []int{0, 3, 1, 0, 4, 5, 2, 0},
			answer: []int{4, 11},
		},
		{
			head:   []int{0, 1, 0, 3, 0, 2, 2, 0},
			answer: []int{1, 3, 4},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				head := q.NewListNodeFromList(tc.head)
				result := q.MergeNodes(head).ToList()
				require.Equal(t, tc.answer, result)
			})
	}
}
