package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMinCostTickets(t *testing.T) {
	type testCase struct {
		name   string
		days   []int
		costs  []int
		answer int
	}

	tcs := []testCase{
		{
			name:   "Empty days",
			days:   []int{},
			costs:  []int{2, 7, 15},
			answer: 0,
		},
		{
			name:   "Len(1) days",
			days:   []int{230},
			costs:  []int{2, 10, 15},
			answer: 2,
		},
		{
			name:   "Two segements",
			days:   []int{1, 3, 5, 7, 61, 62},
			costs:  []int{2, 7, 15},
			answer: 11,
		},
		{
			name:   "General case I",
			days:   []int{1, 4, 6, 7, 8, 20},
			costs:  []int{2, 7, 15},
			answer: 11,
		},
		{
			name:   "General case II",
			days:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
			costs:  []int{2, 7, 15},
			answer: 17,
		},
		{
			name: "Hard case",
			days: []int{
				6, 9, 10, 14, 15, 16, 17, 18, 20, 22, 23, 24, 29, 30,
				31, 33, 35, 37, 38, 40, 41, 46, 47, 51, 54, 57, 59, 65,
				70, 76, 77, 81, 85, 87, 90, 91, 93, 94, 95, 97, 98, 100,
				103, 104, 105, 106, 107, 111, 112, 113, 114, 116, 117,
				118, 120, 124, 128, 129, 135, 137, 139, 145, 146, 151,
				152, 153, 157, 165, 166, 173, 174, 179, 181, 182, 185,
				187, 188, 190, 191, 192, 195, 196, 204, 205, 206, 208,
				210, 214, 218, 219, 221, 225, 229, 231, 233, 235, 239,
				240, 245, 247, 249, 251, 252, 258, 261, 263, 268, 270,
				273, 274, 275, 276, 280, 283, 285, 286, 288, 289, 290,
				291, 292, 293, 296, 298, 299, 301, 303, 307, 313, 314,
				319, 323, 325, 327, 329, 334, 339, 340, 341, 342, 344,
				346, 349, 352, 354, 355, 356, 357, 358, 359, 363, 364},
			costs:  []int{21, 115, 345},
			answer: 3040,
		},
	}

	q := leetcode.Q0983{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.name),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.answer, q.MinCostTickets(tc.days, tc.costs))
			},
		)
	}
}
