package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestAEIOUInterval(t *testing.T) {
	type testCase struct {
		str    string
		hasall bool
		ans    []string
	}

	tcs := []testCase{
		{
			str:    "aeiouu",
			hasall: true,
			ans: []string{
				"aeiou",
			},
		},
		{
			str:    "uaieuoua",
			hasall: true,
			ans: []string{
				"uaieuo",
				"aieuo",
				"ieuoua",
			},
		},
		{
			str:    "aaaaeiouaaaa",
			hasall: true,
			ans: []string{
				"aaaaeiou",
				"aaaeiou",
				"aaeiou",
				"aeiou",
				"eioua",
			},
		},
		{
			str:    "aeio",
			hasall: false,
			ans:    []string{},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.str),
			func(t *testing.T) {
				aeiou := leetcode.NewAEIOU()
				for i := 0; i < len(tc.str); i++ {
					aeiou.Append(i, tc.str[i])
				}
				require.Equal(t, tc.hasall, aeiou.HasAll())
				aeiou.Sort()

				for i := 0; i < len(tc.ans); i++ {
					s, e, ok := aeiou.NextAEIOUInterval()
					require.True(t, ok)
					require.Equal(t, tc.ans[i], tc.str[s:e+1])
				}
				_, _, ok := aeiou.NextAEIOUInterval()
				require.False(t, ok)
			},
		)
	}
}

func TestAEIOUSubStrings(t *testing.T) {
	type testCase struct {
		str    string
		hasall bool
		ans    []string
	}

	tcs := []testCase{
		{
			str:    "aeiouu",
			hasall: true,
			ans: []string{
				"aeiou",
				"aeiouu",
			},
		},
		{
			str:    "uaieuoua",
			hasall: true,
			ans: []string{
				"uaieuo",
				"uaieuou",
				"uaieuoua",
				"aieuo",
				"aieuou",
				"aieuoua",
				"ieuoua",
			},
		},
		{
			str:    "aaaaeiouaaaa",
			hasall: true,
			ans: []string{
				"aaaaeiou",
				"aaaaeioua",
				"aaaaeiouaa",
				"aaaaeiouaaa",
				"aaaaeiouaaaa",
				"aaaeiou",
				"aaaeioua",
				"aaaeiouaa",
				"aaaeiouaaa",
				"aaaeiouaaaa",
				"aaeiou",
				"aaeioua",
				"aaeiouaa",
				"aaeiouaaa",
				"aaeiouaaaa",
				"aeiou",
				"aeioua",
				"aeiouaa",
				"aeiouaaa",
				"aeiouaaaa",
				"eioua",
				"eiouaa",
				"eiouaaa",
				"eiouaaaa",
			},
		},
		{
			str:    "aeio",
			hasall: false,
			ans:    []string{},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.str),
			func(t *testing.T) {
				aeiou := leetcode.NewAEIOU()
				for i := 0; i < len(tc.str); i++ {
					aeiou.Append(i, tc.str[i])
				}
				require.Equal(t, tc.hasall, aeiou.HasAll())
				aeiou.Sort()

				sub := aeiou.AEIOUSubStrings(tc.str)
				require.ElementsMatch(t, tc.ans, sub)
			},
		)
	}
}

func TestNumAEIOUSubStrings(t *testing.T) {
	type testCase struct {
		str    string
		hasall bool
		ans    int
	}

	tcs := []testCase{
		{
			str:    "aeiouu",
			hasall: true,
			ans:    2,
		},
		{
			str:    "uaieuoua",
			hasall: true,
			ans:    7,
		},
		{
			str:    "aaaaeiouaaaa",
			hasall: true,
			ans:    24,
		},
		{
			str:    "aeio",
			hasall: false,
			ans:    0,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.str),
			func(t *testing.T) {
				aeiou := leetcode.NewAEIOU()
				for i := 0; i < len(tc.str); i++ {
					aeiou.Append(i, tc.str[i])
				}
				require.Equal(t, tc.hasall, aeiou.HasAll())
				aeiou.Sort()

				require.Equal(t, tc.ans, aeiou.NumAEIOUSubStrings(tc.str))
			},
		)
	}
}

func TestExtractAEIOUSubStrings(t *testing.T) {
	type testCase struct {
		word string
		ans  []string
	}

	tcs := []testCase{
		{
			word: "aeiou",
			ans:  []string{"aeiou"},
		},
		{
			word: "unicornarihan",
			ans:  []string{},
		},
		{
			word: "cuaieuouac",
			ans:  []string{"uaieuoua"},
		},
		{
			word: "aaeiouccccaeiouu",
			ans:  []string{"aaeiou", "aeiouu"},
		},
	}

	q := leetcode.Q2062{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.word),
			func(t *testing.T) {
				require.ElementsMatch(
					t,
					tc.ans,
					q.ExtractAEIOUSubStrings(tc.word, 5),
				)
			},
		)
	}
}

func TestCountAEUIUSubStrings(t *testing.T) {

	type testCase struct {
		word string
		ans  int
	}

	tcs := []testCase{
		{
			word: "aeiouu",
			ans:  2,
		},
		{
			word: "unicornarihan",
			ans:  0,
		},
		{
			word: "cuaieuouac",
			ans:  7,
		},
		{
			word: "aaeiouccccaeiouuiagjeklordrweqwsdcuaieuouac",
			ans:  13,
		},
	}

	q := leetcode.Q2062{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d-%s", i+1, tc.word),
			func(t *testing.T) {
				require.Equal(t, tc.ans,
					q.CountVowelSubstrings(tc.word))
			},
		)
	}
}
