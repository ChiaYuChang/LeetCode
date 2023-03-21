package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestCanPlaceFlowers(t *testing.T) {
	type testCase struct {
		name      string
		flowerbed []int
		n         int
		ans       bool
	}

	tcs := []testCase{
		{
			name:      "True",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			ans:       true,
		},
		{
			name:      "False",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			ans:       false,
		},
		{
			name:      "Even zeros",
			flowerbed: []int{0, 0, 0, 0},
			n:         2,
			ans:       true,
		},
		{
			name:      "Odd zeros",
			flowerbed: []int{0, 0, 0, 0, 0},
			n:         3,
			ans:       true,
		},
		{
			name:      "One zero",
			flowerbed: []int{0},
			n:         1,
			ans:       true,
		},
	}

	q := leetcode.Q0605{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.name,
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.ans, q.CanPlaceFlowers(tc.flowerbed, tc.n))
			})
	}
}
