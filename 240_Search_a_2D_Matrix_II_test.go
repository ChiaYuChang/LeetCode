package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestSearch2DMatrix2(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	t.Run(
		"Fast method",
		func(t *testing.T) {
			t.Parallel()
			require.True(t, leetcode.SearchMatrixII(matrix, 5, "f"))
			require.True(t, leetcode.SearchMatrixII(matrix, 30, "f"))
			require.True(t, leetcode.SearchMatrixII(matrix, 1, "f"))
			require.True(t, leetcode.SearchMatrixII(matrix, 24, "f"))
			require.True(t, leetcode.SearchMatrixII(matrix, 15, "f"))
			require.False(t, leetcode.SearchMatrixII(matrix, 99, "f"))
			require.False(t, leetcode.SearchMatrixII(matrix, 0, "f"))
		},
	)

	t.Run(
		"Slow method",
		func(t *testing.T) {
			t.Parallel()
			require.True(t, leetcode.SearchMatrixII(matrix, 5, "s"))
			require.True(t, leetcode.SearchMatrixII(matrix, 30, "s"))
			require.True(t, leetcode.SearchMatrixII(matrix, 1, "s"))
			require.True(t, leetcode.SearchMatrixII(matrix, 24, "s"))
			require.True(t, leetcode.SearchMatrixII(matrix, 15, "s"))
			require.False(t, leetcode.SearchMatrixII(matrix, 99, "s"))
			require.False(t, leetcode.SearchMatrixII(matrix, 0, "s"))
		},
	)
}
