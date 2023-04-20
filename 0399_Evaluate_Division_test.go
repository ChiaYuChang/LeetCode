package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestCalcEquation(t *testing.T) {
	type testCase struct {
		equations [][]string
		values    []float64
		queries   [][]string
		answers   []float64
	}

	tcs := []testCase{
		{
			equations: [][]string{
				{"a", "b"},
				{"b", "c"},
			},
			values: []float64{
				2.0,
				3.0,
			},
			queries: [][]string{
				{"a", "c"},
				{"b", "a"},
				{"a", "e"},
				{"a", "a"},
				{"x", "x"},
			},
			answers: []float64{
				6.0,
				0.5,
				-1.0,
				1.0,
				-1.0,
			},
		},
		{
			equations: [][]string{
				{"a", "b"},
				{"b", "c"},
				{"bc", "cd"},
			},
			values: []float64{
				1.5,
				2.5,
				5.0,
			},
			queries: [][]string{
				{"a", "c"},
				{"c", "b"},
				{"bc", "cd"},
				{"cd", "bc"},
			},
			answers: []float64{
				3.75000,
				0.40000,
				5.00000,
				0.20000,
			},
		},
		{
			equations: [][]string{
				{"a", "b"},
			},
			values: []float64{
				0.5,
			},
			queries: [][]string{
				{"a", "b"},
				{"b", "a"},
				{"a", "c"},
				{"x", "y"},
			},
			answers: []float64{
				0.50000,
				2.00000,
				-1.00000,
				-1.00000,
			},
		},
		{
			equations: [][]string{
				{"x1", "x2"},
				{"x2", "x3"},
				{"x3", "x4"},
				{"x4", "x5"},
			},
			values: []float64{
				3.0,
				4.0,
				5.0,
				6.0,
			},
			queries: [][]string{
				{"x1", "x5"},
				{"x5", "x2"},
				{"x2", "x4"},
				{"x2", "x2"},
				{"x2", "x9"},
				{"x9", "x9"},
			},
			answers: []float64{
				360.00000,
				0.00833,
				20.00000,
				1.00000,
				-1.00000,
				-1.00000,
			},
		},
	}

	delta := 0.000_1
	abs := func(v float64) float64 {
		if v < 0 {
			return -v
		}
		return v
	}

	q := leetcode.Q0399{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				answer := q.CalcEquation(
					tc.equations, tc.values, tc.queries)
				require.Equal(t, len(tc.answers), len(answer))
				for i := range answer {
					require.GreaterOrEqual(t, delta, abs(answer[i]-tc.answers[i]))
				}
			},
		)
	}
}
