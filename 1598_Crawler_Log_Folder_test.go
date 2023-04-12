package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinOperations(t *testing.T) {
	type testCase struct {
		logs []string
		ans  int
	}

	tcs := []testCase{
		{
			logs: []string{"d1/", "d2/", "../", "d21/", "./"},
			ans:  2,
		},
		{
			logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			ans:  3,
		},
		{
			logs: []string{"d1/", "../", "../", "../"},
			ans:  0,
		},
		{
			logs: []string{"./", "../", "./"},
			ans:  0,
		},
	}

	q := leetcode.Q1598{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(t, tc.ans, q.MinOperations(tc.logs))
			},
		)
	}
}
