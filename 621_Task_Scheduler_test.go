package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestLeastInterval(t *testing.T) {
	type testCase struct {
		name  string
		tasks []byte
		n     int
		ans   int
	}

	tcs := []testCase{
		{
			name:  "Case 1",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     2,
			ans:   8,
		},
		{
			name:  "Case 2",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     0,
			ans:   6,
		},
		{
			name:  "Case 3",
			tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
			n:     2,
			ans:   16,
		},
		{
			name:  "Case 4",
			tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
			n:     5,
			ans:   31,
		},
		{
			name:  "Case 5",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'D'},
			n:     3,
			ans:   12,
		},
		{
			name:  "Case 6",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'D'},
			n:     1,
			ans:   12,
		},
		{
			name:  "Case 7",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'},
			n:     2,
			ans:   12,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Log(tc.name)
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(t, tc.ans, leetcode.LeastInterval(tc.tasks, tc.n))
			},
		)
	}
}
