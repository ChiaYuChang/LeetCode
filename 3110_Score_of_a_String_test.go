package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestScoreOfString(t *testing.T) {
	tcs := []struct {
		str string
		ans int
	}{
		{"hello", 13},
		{"zaz", 50},
	}

	q := leetcode.Q3110{}
	for i := range tcs {
		tc := tcs[i]
		require.Equal(t, tc.ans, q.ScoreOfString(tc.str))
	}
}
