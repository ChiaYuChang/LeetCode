package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLongestPalindromeFromString(t *testing.T) {
	type testCase struct {
		s   string
		ans int
	}

	tcs := []testCase{
		{
			s:   "abccccdd",
			ans: 7,
		},
		{
			s:   "a",
			ans: 1,
		},
		{
			s:   "abccccdd",
			ans: 7,
		},
		{
			s:   "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoherega",
			ans: 185,
		},
		{
			s:   "aabbcCcC",
			ans: 8,
		},
	}

	q409 := leetcode.Q409{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q409.LongestPalindrome(tc.s))
			},
		)
	}
}
