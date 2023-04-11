package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestFindPoisonedDuration(t *testing.T) {
	type testCase struct {
		Name       string
		TimeSeries []int
		Duration   int
		Answer     int
	}

	tcs := []testCase{
		{
			Name:       "Intervals w/o intersection",
			TimeSeries: []int{1, 4},
			Duration:   2,
			Answer:     4,
		},
		{
			Name:       "Intervals w/ intersection",
			TimeSeries: []int{1, 2},
			Duration:   2,
			Answer:     3,
		},
		{
			Name:       "Intervals can be merged",
			TimeSeries: []int{1, 3, 5},
			Duration:   2,
			Answer:     6,
		},
	}

	q := leetcode.Q0495{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.Name),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.Answer, q.FindPoisonedDuration(tc.TimeSeries, tc.Duration))
			},
		)
	}
}
