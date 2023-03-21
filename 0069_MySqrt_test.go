package leetcode_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMySqrt(t *testing.T) {
	tcs := []int{1, 3, 5, 10, 24, 100, 2147483647}

	q := leetcode.Q0069{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%d", i+1, tc),
			func(t *testing.T) {
				a := int(math.Floor(math.Sqrt(float64(tc))))
				require.Equal(t, a, q.MySqrt(tc))
			},
		)
	}
}
