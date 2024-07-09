package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ1701(t *testing.T) {
	var abs func(x float64) float64
	abs = func(x float64) float64 {
		if x < 0 {
			return -x
		}
		return x
	}
	var delta float64 = 0.0001

	tcs := []struct {
		customers [][]int
		answer    float64
	}{
		{
			customers: [][]int{
				{1, 2}, {2, 5}, {4, 3},
			},
			answer: 5.0,
		},
		{
			customers: [][]int{
				{5, 2}, {5, 4}, {10, 3}, {20, 1},
			},
			answer: 3.25,
		},
		{
			customers: [][]int{
				{30, 100}, {31, 20}, {45, 300}, {49, 1}, {49, 1}, {49, 1}, {50, 31},
			},
			answer: 323.85714,
		},
	}

	q := leetcode.Q1701{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Test %d", i+1),
			func(t *testing.T) {
				require.GreaterOrEqual(
					t, delta,
					abs(tc.answer-q.AverageWaitingTime(tc.customers)))
			},
		)
	}
}
