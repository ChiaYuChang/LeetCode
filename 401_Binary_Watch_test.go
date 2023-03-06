package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestPermute(t *testing.T) {

	type testCase struct {
		nOne  int
		nZero int
		ans   []int
	}

	tcs := []testCase{
		{
			nOne:  0,
			nZero: 3,
			ans: []int{
				0b0000,
			},
		},
		{
			nOne:  1,
			nZero: 3,
			ans: []int{
				0b0001,
				0b0010,
				0b0100,
				0b1000,
			},
		},
		{
			nOne:  3,
			nZero: 2,
			ans: []int{
				0b11100,
				0b11010,
				0b11001,
				0b10110,
				0b10101,
				0b10011,
				0b01110,
				0b01101,
				0b01011,
				0b00111,
			},
		},
	}

	p := leetcode.Permutator{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("# of Ones: %d, # of Zeros: %d\n", tc.nOne, tc.nZero),
			func(t *testing.T) {
				require.ElementsMatch(t, tc.ans, p.Permute(tc.nOne, tc.nZero))
			},
		)
	}
}

func TestReadBinaryWatch(t *testing.T) {
	require.ElementsMatch(t,
		[]string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		leetcode.ReadBinaryWatch(1),
	)

	require.ElementsMatch(t,
		[]string{},
		leetcode.ReadBinaryWatch(9),
	)
}
