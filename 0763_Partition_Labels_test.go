package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestPartitionLabels(t *testing.T) {
	type testCase struct {
		s   string
		asn []int
	}

	tcs := []testCase{
		{
			s:   "ababcbacadefegdehijhklij",
			asn: []int{9, 7, 8},
		},
		{
			s:   "eccbbbbdec",
			asn: []int{10},
		},
		{
			s:   "ababccdecdf",
			asn: []int{4, 6, 1},
		},
	}

	q := leetcode.Q0763{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.asn, q.PartitionLabels(tc.s))
			},
		)
	}
}
