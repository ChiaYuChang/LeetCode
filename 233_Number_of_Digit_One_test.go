package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestNumberOfDigitOne(t *testing.T) {
	require.Equal(t, 6, leetcode.NumberOfDigitOne(13))
	require.Equal(t, 0, leetcode.NumberOfDigitOne(0))
}
