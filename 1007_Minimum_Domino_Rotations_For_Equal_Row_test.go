package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinDominoRotations(t *testing.T) {
	type testCase struct {
		name          string
		tops, bottoms []int
		answer        int
	}

	tcs := []testCase{
		{
			name:    "OK",
			tops:    []int{2, 1, 2, 4, 2, 2},
			bottoms: []int{5, 2, 6, 2, 3, 2},
			answer:  2,
		}, {
			name:    "Not possible",
			tops:    []int{3, 5, 1, 2, 3},
			bottoms: []int{3, 6, 3, 3, 4},
			answer:  -1,
		},
		{
			name:    "Bottoms",
			tops:    []int{1, 2, 1, 1, 1, 2, 2, 2},
			bottoms: []int{2, 1, 2, 2, 2, 2, 2, 2},
			answer:  1,
		},
	}

	q := leetcode.Q1007{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.answer, q.MinDominoRotations(tc.tops, tc.bottoms))
			},
		)
	}
}
