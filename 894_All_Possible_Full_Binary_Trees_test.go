package leetcode_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"
// 	"gitlab.com/gjerry134679/leetcode"
// )

// func requireBTEqual(t *testing.T, expected, actual *leetcode.Q894TreeNode) {
// 	if expected == nil {
// 		require.Nil(t, actual)
// 		return
// 	}
// 	require.NotNil(t, actual)
// 	t.Logf("Val: %d\n", expected.Val)
// 	require.Equal(t, expected.Val, actual.Val)
// 	requireBTEqual(t, expected.Left, actual.Left)
// 	requireBTEqual(t, expected.Right, actual.Right)
// }

// func TestCopyTreeNode(t *testing.T) {
// 	n1 := &leetcode.Q894TreeNode{Val: 1}

// 	curr := n1
// 	curr.Left = &leetcode.Q894TreeNode{Val: 2}
// 	curr.Right = &leetcode.Q894TreeNode{Val: 3}

// 	curr = n1.Left
// 	curr.Left = &leetcode.Q894TreeNode{Val: 4}
// 	curr.Right = &leetcode.Q894TreeNode{Val: 5}
// 	curr = curr.Left
// 	curr.Left = &leetcode.Q894TreeNode{Val: 8}

// 	curr = n1.Right
// 	curr.Left = &leetcode.Q894TreeNode{Val: 6}
// 	curr.Right = &leetcode.Q894TreeNode{Val: 7}

// 	n2 := n1.CopyTreeNode()

// 	requireBTEqual(t, n1, n2)
// }
