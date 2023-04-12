package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinDistance(t *testing.T) {
	type testCase struct {
		word1, word2 string
		distance     int
	}

	tcs := []testCase{
		{"sea", "eat", 2},
		{"leetcode", "etco", 4},
	}

	q := leetcode.Q0583{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.distance, q.MinDistance(tc.word1, tc.word2))
			},
		)
	}
}
