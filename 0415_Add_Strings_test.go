package leetcode_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestAddByte(t *testing.T) {
	type testCase struct {
		b, b1, b2 byte
		c, c0     bool
	}

	tcs := []testCase{
		{
			b1: '1',
			b2: '2',
			c0: false,
			b:  '3',
			c:  false,
		},
		{
			b1: '7',
			b2: '2',
			c0: false,
			b:  '9',
			c:  false,
		},
		{
			b1: '4',
			b2: '6',
			c0: false,
			b:  '0',
			c:  true,
		},
		{
			b1: '4',
			b2: '8',
			c0: true,
			b:  '3',
			c:  true,
		},
	}

	q := leetcode.Q0415{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d: %c+%c", i+1, tc.b1, tc.b2),
			func(t *testing.T) {
				b, c := q.AddByte(tc.b1, tc.b2, tc.c0)
				require.Equal(t, tc.b, b)
				require.Equal(t, tc.c, c)
			},
		)
	}
}

func TestAddString(t *testing.T) {
	var AddStrings func(s1, s2 string) string

	AddStrings = func(s1, s2 string) string {
		n1, _ := strconv.Atoi(s1)
		n2, _ := strconv.Atoi(s2)
		return strconv.Itoa(n1 + n2)
	}

	type testCase struct {
		num1, num2 string
	}

	tcs := []testCase{
		{
			num1: "11",
			num2: "123",
		},
		{
			num1: "456",
			num2: "77",
		},
		{
			num1: "0",
			num2: "0",
		},
		{
			num1: "10000000",
			num2: "1",
		},
	}

	q := leetcode.Q0415{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s+%s", i+1, tc.num1, tc.num2),
			func(t *testing.T) {
				require.Equal(t,
					AddStrings(tc.num1, tc.num2),
					q.AddStrings(tc.num1, tc.num2),
				)

				require.Equal(t,
					AddStrings(tc.num2, tc.num1),
					q.AddStrings(tc.num2, tc.num1),
				)
			},
		)
	}
}
