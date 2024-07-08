package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ1823(t *testing.T) {
	q := leetcode.Q1823{}

	tcs := []struct {
		n      int
		k      int
		answer int
	}{
		{5, 2, 3},
		{6, 5, 1},
		{3, 1, 3},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("n=%d,k=%d", tc.n, tc.k),
			func(t *testing.T) {
				require.Equal(t, tc.answer, q.FindTheWinner(tc.n, tc.k))
			},
		)
	}
}
