package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ1438(t *testing.T) {
	q1438 := leetcode.Q1438{}

	tcs := []struct {
		nums   []int
		limit  int
		answer int
	}{
		{
			nums:   []int{8, 2, 4, 7},
			limit:  4,
			answer: 2,
		},
		{
			nums:   []int{10, 1, 2, 4, 7, 2},
			limit:  5,
			answer: 4,
		},
		{
			nums:   []int{4, 2, 2, 2, 4, 4, 2, 2},
			limit:  0,
			answer: 3,
		},
		{
			nums:   []int{2, 2, 2, 4, 4, 2, 5, 5, 5, 5, 5, 2},
			limit:  2,
			answer: 6,
		},
		{
			nums:   []int{-2, -2, -2, -4, -4, -2, -5, -5, -5, -5, -5, -2},
			limit:  2,
			answer: 6,
		},
		{
			nums: []int{
				24, 12, 71, 33, 5, 87, 10, 11, 3, 58,
				2, 97, 97, 36, 32, 35, 15, 80, 24, 45,
				38, 9, 22, 21, 33, 68, 22, 85, 35, 83,
				92, 38, 59, 90, 42, 64, 61, 15, 4, 40,
				50, 44, 54, 25, 34, 14, 33, 94, 66, 27,
				78, 56, 3, 29, 3, 51, 19, 5, 93, 21,
				58, 91, 65, 87, 55, 70, 29, 81, 89, 67,
				58, 29, 68, 84, 4, 51, 87, 74, 42, 85,
				81, 55, 8, 95, 39},
			limit:  87,
			answer: 25,
		},
		{
			nums: []int{
				24, 12, 71, 33, 5, 87, 10, 11, 3, 58,
				2, 97, 96, 36, 32, 35, 15, 80, 24, 45,
				38, 9, 22, 21, 33, 68, 22, 85, 35, 83,
				92, 38, 59, 90, 42, 64, 61, 15, 4, 40,
				50, 44, 54, 25, 34, 14, 33, 94, 66, 27,
				78, 56, 3, 29, 3, 51, 19, 5, 93, 21,
				58, 91, 65, 87, 55, 70, 29, 81, 89, 67,
				58, 29, 68, 84, 4, 51, 87, 74, 42, 85,
				81, 55, 8, 95, 39},
			limit:  87,
			answer: 26,
		},
		{
			nums: []int{
				18, 100, 79, 85, 88, 90, 11, 57, 31, 49,
				56, 51, 22, 42, 57, 17, 80, 63, 28, 16,
				56, 77, 69, 53, 16, 85, 38, 36, 32, 49,
				96, 72, 1, 25, 68, 57, 75, 3, 4, 81,
				78, 32, 34, 27, 23, 37, 19, 70, 26, 35,
				40, 75, 97, 33, 88, 58, 22, 70, 46, 63,
				54, 16, 99, 27, 74, 50, 27, 37, 14, 4,
				16, 73, 96, 2, 70, 38, 87, 98, 93, 84,
				18, 10, 65, 48, 74, 40, 56, 65, 87, 46,
				98, 68, 42, 1, 16, 57, 92,
			},
			limit:  55,
			answer: 10,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(fmt.Sprintf("TwoQueues_Case_%d", i+1), func(t *testing.T) {
			t.Parallel()
			ans := q1438.TwoQueuesLongestSubarray(tc.nums, tc.limit)
			require.Equal(t, tc.answer, ans)
		})
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(fmt.Sprintf("TwoDeques_Case_%d", i+1), func(t *testing.T) {
			t.Parallel()
			ans := q1438.TwoDequesLongestSubarray(tc.nums, tc.limit)
			require.Equal(t, tc.answer, ans)
		})
	}
}
