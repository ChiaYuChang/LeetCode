package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestReformat(t *testing.T) {
	type testCase struct {
		s string
		a string
	}

	IsDigit := func(b byte) bool {
		return '0' <= b && b <= '9'
	}

	HasSameElements := func(s1, s2 string) bool {
		if len(s1) != len(s2) {
			return false
		}

		m := map[byte]int{}
		for i := 0; i < len(s1); i++ {
			m[s1[i]] += 1
		}

		for i := 0; i < len(s2); i++ {
			m[s2[i]] -= 1
		}

		for _, v := range m {
			if v != 0 {
				return false
			}
		}
		return true
	}

	tcs := []testCase{
		{s: "a0b1c2", a: "0a1b2c"},
		{s: "a0b1c2d", a: "a0b1c2d"},
		{s: "leetcode", a: ""},
		{s: "1229857369", a: ""},
	}

	q := leetcode.Q1417{}
	for i, tc := range tcs {
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.s),
			func(t *testing.T) {
				ans := q.Reformat(tc.s)
				if tc.a == "" {
					require.Empty(t, ans)
				} else {
					require.True(t, HasSameElements(tc.a, ans))
					f := IsDigit(ans[0])
					for i := 0; i < len(ans); i++ {
						if f {
							require.True(t, IsDigit(ans[i]))
						} else {
							require.False(t, IsDigit(ans[i]))
						}
						f = !f
					}
				}
			},
		)
	}
}
