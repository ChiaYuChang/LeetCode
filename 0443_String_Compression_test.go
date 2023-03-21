package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestStringCompression(t *testing.T) {
	type testCase struct {
		chars []byte
		ans   []byte
	}

	tcs := []testCase{
		{
			chars: []byte("aabbccc"),
			ans:   []byte("a2b2c3"),
		},
		{
			chars: []byte("a"),
			ans:   []byte("a"),
		},
		{
			chars: []byte("abbbbbbbbbbbb"),
			ans:   []byte("ab12"),
		},
		{
			chars: []byte("aaabbcdddd"),
			ans:   []byte("a3b2cd4"),
		},
		{
			chars: []byte{},
			ans:   []byte{},
		},
	}
	q := leetcode.Q0443{}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, string(tc.chars)),
			func(t *testing.T) {
				n := q.StringCompression(tc.chars)
				require.Equal(t, string(tc.ans), string(tc.chars[:n]))
				require.Equal(t, len(tc.ans), n)
			},
		)
	}
}
