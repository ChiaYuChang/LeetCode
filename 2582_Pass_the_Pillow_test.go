package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ2582(t *testing.T) {
	tcs := []struct {
		n    int
		time int
		ans  int
	}{
		{4, 8, 3},
		{3, 2, 3},
		{300, 523, 76},
		{9, 4, 5},
		{4, 8, 3},
	}

	q := leetcode.Q2582{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("n=%d,time=%d", tc.n, tc.time),
			func(t *testing.T) {
				require.Equal(
					t, tc.ans,
					q.PassThePillow(tc.n, tc.time),
				)
			},
		)
	}
}
