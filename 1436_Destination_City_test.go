package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestDestCity(t *testing.T) {

	type testCase struct {
		name  string
		paths [][]string
		dest  string
	}

	tcs := []testCase{
		{
			name: "Case 1",
			paths: [][]string{
				{"London", "New York"},
				{"New York", "Lima"},
				{"Lima", "Sao Paulo"},
			},
			dest: "Sao Paulo",
		},
		{
			name: "Case 2",
			paths: [][]string{
				{"B", "C"},
				{"D", "B"},
				{"C", "A"},
			},
			dest: "A",
		},
		{
			name: "Case 3",
			paths: [][]string{
				{"A", "Z"},
			},
			dest: "Z",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				require.Equal(t, tc.dest, leetcode.DestCity(tc.paths))
			},
		)
	}
}
