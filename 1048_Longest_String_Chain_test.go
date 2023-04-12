package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLongestStrChain(t *testing.T) {
	type testCase struct {
		words  []string
		length int
	}

	tcs := []testCase{
		{[]string{"a", "b", "ba", "bca", "bda", "bdca"}, 4},
		{[]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}, 5},
		{[]string{"abcd", "dbqca"}, 1},
	}

	q := leetcode.Q1048{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.length, q.LongestStrChain(tc.words))
			},
		)
	}
}
