package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestTopKFreqElements(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	require.ElementsMatch(t, []int{1, 2}, leetcode.TopKFrequent(nums, k))
}
