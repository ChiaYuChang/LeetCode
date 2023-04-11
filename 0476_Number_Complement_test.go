package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindComplement(t *testing.T) {
	tcs := [][2]int{
		{5, 2},
		{1, 0},
		{2, 1},
	}

	q := leetcode.Q0476{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%d", i+1, tc[0]),
			func(t *testing.T) {
				require.Equal(t, tc[1], q.FindComplement(tc[0]))
			},
		)
	}
}
