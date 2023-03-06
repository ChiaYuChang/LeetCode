package leetcode_test

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestMultByte(t *testing.T) {
	type testCase struct {
		b1, b2, b byte
		d1, d2    int
		ans       string
	}

	tcs := []testCase{
		{
			b1:  '1',
			b2:  '2',
			d1:  0,
			d2:  0,
			ans: "2",
		},
		{
			b1:  '1',
			b2:  '2',
			d1:  1,
			d2:  0,
			ans: "20",
		},
		{
			b1:  '1',
			b2:  '2',
			d1:  0,
			d2:  1,
			ans: "20",
		},
		{
			b1:  '1',
			b2:  '2',
			d1:  1,
			d2:  2,
			ans: "2000",
		},
	}
	q := leetcode.Q043{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				require.Equal(
					t, tc.ans,
					q.MultByte(tc.b1, tc.b2, tc.d1, tc.d2).String(),
				)
			},
		)
	}
}

func TestMultString(t *testing.T) {
	type testCase struct {
		num1, num2 string
	}

	var multString func(s1, s2 string) string = func(s1, s2 string) string {
		n1, _ := strconv.Atoi(s1)
		n2, _ := strconv.Atoi(s2)

		b1 := big.NewInt(int64(n1))
		b2 := big.NewInt(int64(n2))

		p := b1.Mul(b1, b2)
		return p.String()
	}

	tcs := []testCase{
		{
			num1: "12",
			num2: "0",
		},
		{
			num1: "123",
			num2: "424",
		},
		{
			num1: "23",
			num2: "125253",
		},
		{
			num1: "498828660196",
			num2: "840477629533",
		},
	}

	q := leetcode.Q043{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%sx%s", i+1, tc.num1, tc.num2),
			func(t *testing.T) {
				require.Equal(
					t, multString(tc.num1, tc.num2),
					q.MultString(tc.num1, tc.num2),
				)
			},
		)
	}
}
