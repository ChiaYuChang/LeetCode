package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestStrStr(t *testing.T) {
	type testCase struct {
		name             string
		haystack, needle string
		anser            int
	}

	tcs := []testCase{
		{
			name:     "General case",
			haystack: "sadbutsad",
			needle:   "sad",
			anser:    0,
		},
		{
			name:     "Needle not found",
			haystack: "leetcode",
			needle:   "leeto",
			anser:    -1,
		},
		{
			name:     "Haystack contains only one kind of char",
			haystack: "aaa",
			needle:   "a",
		},
		{
			name:   "Empty needle",
			needle: "",
			anser:  0,
		},
	}

	q := leetcode.Q0028{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.anser,
					q.StrStr(tc.haystack, tc.needle))
			},
		)
	}
}
