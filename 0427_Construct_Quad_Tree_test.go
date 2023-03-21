package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func Test2DMatrix(t *testing.T) {
	q := leetcode.Q0427{}
	m := q.NewTwoDMatrix([][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{2, 2, 3, 3, 4, 4, 5, 5},
		{2, 2, 3, 3, 4, 4, 5, 5},
		{6, 6, 7, 7, 8, 8, 9, 9},
		{6, 6, 7, 7, 8, 8, 9, 9},
	})
	require.Equal(t, 8, m.NCol())
	require.Equal(t, 8, m.NRow())

	e, ok := m.Element(0, 0)
	require.True(t, ok)
	require.Equal(t, 1, e)

	e, ok = m.Element(0, 7)
	require.True(t, ok)
	require.Equal(t, 0, e)

	_, ok = m.Element(0, 8)
	require.False(t, ok)

	e, ok = m.Element(7, 0)
	require.True(t, ok)
	require.Equal(t, 6, e)

	e, ok = m.Element(7, 7)
	require.True(t, ok)
	require.Equal(t, 9, e)

	_, ok = m.Element(8, 0)
	require.False(t, ok)

	subm, _ := m.SubMatrix(4, 5, 2, 3)
	require.Equal(t, 2, subm.NCol())
	require.Equal(t, 2, subm.NRow())

	e, ok = subm.Element(0, 0)
	require.True(t, ok)
	require.Equal(t, 3, e)

	e, ok = subm.Element(1, 1)
	require.True(t, ok)
	require.Equal(t, 3, e)

	_, ok = subm.Element(2, 0)
	require.False(t, ok)

	_, ok = m.SubMatrix(4, 8, 2, 3)
	require.False(t, ok)
}
