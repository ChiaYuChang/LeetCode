package gdatastructs_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	gds "gitlab.com/gjerry134679/leetcode/gDataStructs"
)

func TestBinaryTree(t *testing.T) {
	tree := gds.NewBinaryTree[int, string]("node")
	require.NotNil(t, tree)
	require.Zero(t, tree.Size())

	type testCase struct {
		action string
		key    int
		val    string
		size   int
	}

	tcs := []testCase{
		{action: "append", key: 5, val: "five"},
		{action: "append", key: 10, val: "ten"},
		{action: "append", key: 7, val: "seven"},
		{action: "append", key: 2, val: "two"},
		{action: "append", key: 1, val: "one"},
		{action: "append", key: 3, val: "three"},
		{action: "append", key: 4, val: "four"},
		{action: "append", key: 6, val: "six"},
		{action: "append", key: 9, val: "typo"},
		{action: "udpate", key: 9, val: "nine"},
	}

	size := 1
	for i := range tcs {
		tc := tcs[i]
		tc.size = size
		if tc.action == "append" {
			size++
		}

		t.Run(
			fmt.Sprintf("%s %d: %s", tc.action, tc.key, tc.val),
			func(t *testing.T) {
				tree.Append(tc.key, tc.val)
				require.Equal(t, tc.size, tree.Size())
				v, ok := tree.Find(tc.key)
				require.True(t, ok)
				require.Equal(t, tc.val, v)
			},
		)
	}
}

func TestLowestCommonAncestorInBinaryTree(t *testing.T) {
	type testCase struct {
		name   string
		keys   []int
		vals   []string
		x      int
		y      int
		ansKey int
		ansVal string
	}

	tcs := []testCase{
		{
			name:   "Case_1_root",
			keys:   []int{6, 2, 8, 0, 4, 7, 9, 3, 5},
			vals:   []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
			x:      2,
			y:      8,
			ansKey: 6,
			ansVal: "a",
		},
		{
			name:   "Case_2_between",
			keys:   []int{6, 2, 8, 0, 4, 7, 9, 3, 5},
			vals:   []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
			x:      5,
			y:      0,
			ansKey: 2,
			ansVal: "b",
		},
		{
			name:   "Case_3_bounded",
			keys:   []int{6, 2, 8, 0, 4, 7, 9, 3, 5},
			vals:   []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
			x:      4,
			y:      2,
			ansKey: 2,
			ansVal: "b",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				t.Parallel()
				tree := gds.NewBinaryTree[int, string]("node")
				require.NotNil(t, tree)
				require.Zero(t, tree.Size())

				for i, k := range tc.keys {
					tree.Append(k, tc.vals[i])
				}

				k, v := tree.LowestCommonAncestor(2, 8)
				require.Equal(t, 6, k)
				require.Equal(t, "a", v)
			},
		)
	}

}
