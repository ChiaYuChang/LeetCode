package leetcode_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindDuplicate(t *testing.T) {
	type testCase struct {
		name  string
		paths []string
		ans   [][]string
	}

	tcs := []testCase{
		{
			"General Case",
			[]string{
				"root/a 1.txt(abcd) 2.txt(efgh)",
				"root/c 3.txt(abcd)",
				"root/c/d 4.txt(efgh)",
				"root 4.txt(efgh)",
			},
			[][]string{
				{
					"root/a/2.txt",
					"root/c/d/4.txt",
					"root/4.txt",
				},
				{
					"root/a/1.txt",
					"root/c/3.txt",
				},
			},
		},
		{
			"No duplicated file",
			[]string{
				"root/a 1.txt(FB) 2.txt(a)",
				"root/c 3.txt(Ea)",
				"root/c/d 4.txt(b)",
				"root 4.txt(c)",
			},
			[][]string{},
		},
	}

	q := leetcode.Q0609{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d %s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				ds := q.FindDuplicate(tc.paths)

				ans := make([]string, len(tc.ans))
				for i := 0; i < len(tc.ans); i++ {
					sort.Sort(sort.StringSlice(tc.ans[i]))
					ans[i] = strings.Join(tc.ans[i], ", ")
				}

				rst := make([]string, len(ds))
				for i := 0; i < len(rst); i++ {
					sort.Sort(sort.StringSlice(ds[i]))
					rst[i] = strings.Join(ds[i], ", ")
				}
				require.ElementsMatch(t, ans, rst)
			},
		)
	}
}
