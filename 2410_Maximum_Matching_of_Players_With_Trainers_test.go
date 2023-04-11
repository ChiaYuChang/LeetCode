package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMatchPlayersAndTrainers(t *testing.T) {
	type testCase struct {
		players  []int
		trainers []int
		answer   int
	}

	tcs := []testCase{
		{
			[]int{4, 7, 9},
			[]int{8, 2, 5, 8},
			2,
		},
		{
			[]int{1, 1, 1},
			[]int{10},
			1,
		},
	}

	q := leetcode.Q2410{}
	for i, tc := range tcs {
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.answer, q.MatchPlayersAndTrainers(tc.players, tc.trainers))
			},
		)
	}
}
