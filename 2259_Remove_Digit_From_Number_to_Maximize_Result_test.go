package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestQ2259(t *testing.T) {
	tcs := []struct {
		number string
		digit  byte
		answer string
	}{
		{"123", '3', "12"},
		{"1231", '1', "231"},
		{"551", '5', "51"},
		{"155", '5', "15"},
	}

	q := leetcode.Q2259{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case_%s", tc.number),
			func(t *testing.T) {
				require.Equal(t, tc.answer, q.RemoveDigit(tc.number, tc.digit))
			},
		)
	}
}
