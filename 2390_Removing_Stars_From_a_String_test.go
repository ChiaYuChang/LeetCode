package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestRemoveStarss(t *testing.T) {
	type testCase struct {
		str string
		ans string
	}

	tcs := []testCase{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
		{"e**", ""},
		{"******e**", ""},
	}

	q := leetcode.Q2390{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.RemoveStars(tc.str))
			},
		)
	}
}
