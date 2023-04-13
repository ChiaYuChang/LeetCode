package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestEscapeGhosts(t *testing.T) {
	type testCase struct {
		name      string
		ghosts    [][]int
		target    []int
		canEscape bool
	}

	tcs := []testCase{
		{
			name:      "OK",
			ghosts:    [][]int{{1, 0}, {0, 3}},
			target:    []int{0, 1},
			canEscape: true,
		},
		{
			name:      "Can not escape",
			ghosts:    [][]int{{1, 0}},
			target:    []int{1, 0},
			canEscape: false,
		},
		{
			name:      "Reach at the same time",
			ghosts:    [][]int{{2, 0}},
			target:    []int{1, 0},
			canEscape: false,
		},
		{
			name:      "Negative position I",
			ghosts:    [][]int{{1, 9}, {2, -5}, {3, 8}, {9, 8}, {-1, 3}},
			target:    []int{8, -10},
			canEscape: false,
		},
		{
			name:      "Negative position II",
			ghosts:    [][]int{{-1, 2}, {0, 1}, {-2, 3}, {0, 1}, {-5, 0}},
			target:    []int{-2, 0},
			canEscape: true,
		},
	}

	q := leetcode.Q0789{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.canEscape, q.EscapeGhosts(tc.ghosts, tc.target))
			},
		)
	}
}
