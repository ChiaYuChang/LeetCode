package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestDecodeCiphertext(t *testing.T) {
	type testCase struct {
		EncodedText  string
		Rows         int
		OriginalText string
	}

	tcs := []testCase{
		{
			EncodedText:  "ch   ie   pr",
			Rows:         3,
			OriginalText: "cipher",
		},
		{
			EncodedText:  "iveo    eed   l te   olc",
			Rows:         4,
			OriginalText: "i love leetcode",
		},
		{
			EncodedText:  "coding",
			Rows:         1,
			OriginalText: "coding",
		},
		{
			EncodedText:  "",
			Rows:         5,
			OriginalText: "",
		},
		{
			EncodedText:  "wmihfwf bddhzaizuzhbuoovyyjstardqceaqzafdzihjbj ywly amkeemr jmvsfaavbpgiafgxzciwmrrtasthc hqfrtwoizoilw",
			Rows:         2,
			OriginalText: "wammikhefewmfr  bjdmdvhszfaaiazvubzphgbiuaofogvxyzycjiswtmarrrdtqacsetahqcz ahfqdfzrithwjobijz oyiwllwy",
		},
	}

	q := leetcode.Q2075{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				t.Parallel()
				require.Equal(t, tc.OriginalText, q.DecodeCiphertext(tc.EncodedText, tc.Rows))
			},
		)
	}
}
