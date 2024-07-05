package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFrequencySort(t *testing.T) {
	q := leetcode.Q0451{}

	tcs := []struct {
		str string
		ans string
	}{
		{"tree", "eetr"}, // "eert" and "eetr" are both correct
		{"cccaaa", "aaaccc"},
		{"Aabb", "bbaA"}, // "bbaA" and "bbAa" are both correct
	}

	for i := range tcs {
		tc := tcs[i]
		require.Equal(t, tc.ans, q.FrequencySort(tc.str))
	}
}
