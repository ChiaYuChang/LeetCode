package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQueensAttacktheKing(t *testing.T) {
	type testCase struct {
		queens [][]int
		king   []int
		answer [][]int
	}

	tcs := []testCase{
		{
			queens: [][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}},
			king:   []int{0, 0},
			answer: [][]int{{0, 1}, {1, 0}, {3, 3}},
		},
		{
			queens: [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {3, 5}, {4, 4}, {4, 5}},
			king:   []int{3, 3},
			answer: [][]int{{2, 2}, {3, 4}, {4, 4}},
		},
		{
			queens: [][]int{{5, 6}, {7, 7}, {2, 1}, {0, 7}, {1, 6}, {5, 1}, {3, 7}, {0, 3}, {4, 0}, {1, 2}, {6, 3}, {5, 0}, {0, 4}, {2, 2}, {1, 1}, {6, 4}, {5, 4}, {0, 0}, {2, 6}, {4, 5}, {5, 2}, {1, 4}, {7, 5}, {2, 3}, {0, 5}, {4, 2}, {1, 0}, {2, 7}, {0, 1}, {4, 6}, {6, 1}, {0, 6}, {4, 3}, {1, 7}},
			king:   []int{3, 4},
			answer: [][]int{{2, 3}, {1, 4}, {1, 6}, {3, 7}, {4, 3}, {5, 4}, {4, 5}},
		},
	}

	q := leetcode.Q1222{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-Distance", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.ElementsMatch(t, tc.answer,
					q.QueensAttacktheKing(tc.queens, tc.king),
				)
			},
		)
		t.Run(
			fmt.Sprintf("Case %d-FirstHit", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.ElementsMatch(t, tc.answer,
					q.QueensAttacktheKingFirstHit(tc.queens, tc.king),
				)
			},
		)
	}
}
