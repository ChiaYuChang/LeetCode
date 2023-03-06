package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLongestCommonPrefix(t *testing.T) {
	type testCase struct {
		name   string
		strs   []string
		prefix string
	}

	tcs := []testCase{
		{
			name:   "Case 1",
			strs:   []string{"flower", "flow", "flight"},
			prefix: "fl",
		},
		{
			name:   "Case 2",
			strs:   []string{"dog", "racecar", "car"},
			prefix: "",
		},
		{
			name:   "Case 3",
			strs:   []string{"", "flower", "flow", "flight"},
			prefix: "",
		},
		{
			name:   "Empty Strings Array",
			strs:   []string{},
			prefix: "",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.prefix, leetcode.LongestCommonPrefix(tc.strs))
			},
		)
	}
}
