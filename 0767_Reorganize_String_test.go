package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestReorganizeString(t *testing.T) {
	q := leetcode.Q0767{}
	require.NotZero(t, q.ReorganizeString("aab"))
	require.NotZero(t, q.ReorganizeString("aabbcc"))
	require.NotZero(t, q.ReorganizeString("aabbccddddddddeeeeeeeee"))
	require.NotZero(t, q.ReorganizeString("vvvlo"))
	require.Zero(t, q.ReorganizeString("aaab"))
}
