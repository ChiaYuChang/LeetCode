package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
	gd "gitlab.com/gjerry134679/leetcode/gDataStructs"
)

func TestAddTwoNums(t *testing.T) {

	type testCase struct {
		l1  *gd.Node[int]
		l2  *gd.Node[int]
		ans *gd.Node[int]
	}

	tcs := []testCase{
		{
			l1:  gd.ArrayToStackQueue([]int{2, 4, 3}, false).Head,
			l2:  gd.ArrayToStackQueue([]int{5, 6, 4}, false).Head,
			ans: gd.ArrayToStackQueue([]int{7, 0, 8}, false).Head,
		},
		{
			l1:  gd.ArrayToStackQueue([]int{0}, false).Head,
			l2:  gd.ArrayToStackQueue([]int{0}, false).Head,
			ans: gd.ArrayToStackQueue([]int{0}, false).Head,
		},
		{
			l1:  gd.ArrayToStackQueue([]int{9, 9, 9, 9, 9, 9, 9}, false).Head,
			l2:  gd.ArrayToStackQueue([]int{9, 9, 9, 9}, false).Head,
			ans: gd.ArrayToStackQueue([]int{8, 9, 9, 9, 0, 0, 0, 1}, false).Head,
		},
		{
			l1:  gd.ArrayToStackQueue([]int{1, 6, 0, 3, 3, 6, 7, 2, 0, 1}, false).Head,
			l2:  gd.ArrayToStackQueue([]int{6, 3, 0, 8, 9, 6, 6, 9, 6, 1}, false).Head,
			ans: gd.ArrayToStackQueue([]int{7, 9, 0, 1, 3, 3, 4, 2, 7, 2}, false).Head,
		},
	}

	for i, c := range tcs {
		t.Run(
			fmt.Sprintf("Case %d", i),
			func(t *testing.T) {
				expAns := c.ans
				getAns, err := leetcode.AddTwoNumbers(c.l1, c.l2)
				require.NoError(t, err)
				for expAns != nil {
					require.Equal(t, *expAns.Data(), *getAns.Data())
					expAns = expAns.Next()
					getAns = getAns.Next()
				}
			},
		)
	}
}
