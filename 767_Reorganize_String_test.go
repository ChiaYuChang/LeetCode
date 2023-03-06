package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestReorganizeString(t *testing.T) {
	require.NotZero(t, leetcode.ReorganizeString("aab"))
	require.NotZero(t, leetcode.ReorganizeString("aabbcc"))
	require.NotZero(t, leetcode.ReorganizeString("aabbccddddddddeeeeeeeee"))
	require.NotZero(t, leetcode.ReorganizeString("vvvlo"))
	require.Zero(t, leetcode.ReorganizeString("aaab"))
}
