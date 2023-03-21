package leetcode_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

type intlist [][]int

func (a intlist) Len() int      { return len(a) }
func (a intlist) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a intlist) Less(i, j int) bool {
	if len(a[i]) == 0 && len(a[j]) == 0 {
		return true
	}
	if len(a[i]) < len(a[j]) {
		return true
	} else if len(a[i]) > len(a[j]) {
		return false
	}

	for k := 0; k < len(a[i]); k++ {
		if a[i][k] != a[j][k] {
			return a[i][k] < a[j][k]
		}
	}
	return true
}

func TestCombinationSum(t *testing.T) {
	type testCase struct {
		name       string
		candidates []int
		target     int
		answer     intlist
	}

	tcs := []testCase{
		{
			name:       "Case_1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			answer:     intlist{{2, 2, 3}, {7}},
		},
		{
			name:       "Case_2",
			candidates: []int{2, 3, 5},
			target:     8,
			answer:     intlist{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "Case_3",
			candidates: []int{2},
			target:     1,
			answer:     intlist{},
		},
		{
			name:       "No Candidates",
			candidates: []int{},
			target:     12,
			answer:     nil,
		},
	}

	q := leetcode.Q0039{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				sort.Sort(tc.answer)
				il := intlist(q.CombinationSum(tc.candidates, tc.target))
				sort.Sort(il)

				t.Logf("expected: %v\n", tc.answer)
				t.Logf("actual  : %v\n", il)

				if tc.answer.Len() == 0 {
					require.Empty(t, il)
				} else {
					require.Equal(t, tc.answer.Len(), il.Len())
				}
			},
		)
	}
}

func TestIntListSortBy(t *testing.T) {
	ils := intlist{
		{2, 2, 2, 2},
		{1, 2, 2, 3},
		{1, 1, 2, 2, 2},
		{2, 3, 3},
		{3, 5},
	}
	sort.Sort(ils)

	for i, il := range ils {
		t.Logf("%d: %v", i, il)
	}
}
