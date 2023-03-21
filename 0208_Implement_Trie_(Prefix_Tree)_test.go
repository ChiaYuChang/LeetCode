package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestPrefixTree(t *testing.T) {
	type Action struct {
		action   byte
		argument string
		output   bool
	}

	type testCase []*Action

	tcs := []testCase{
		{
			{action: 'n'},
			{action: 'i', argument: "apple"},
			{action: 's', argument: "apple", output: true},
			{action: 's', argument: "app", output: false},
			{action: 'p', argument: "app", output: true},
			{action: 'i', argument: "app"},
			{action: 's', argument: "app", output: true},
		},
	}

	q := leetcode.Q0208{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				var trie leetcode.Q0208Trie
				for i, a := range tc {
					t.Log("step", i)
					switch a.action {
					case 'n':
						trie = q.NewQ0208Trie()
					case 'i':
						trie.Insert(a.argument)
					case 's':
						require.Equal(t, a.output, trie.Search(a.argument))
					case 'p':
						require.Equal(t, a.output, trie.StartsWith(a.argument))
					}
				}
			},
		)
	}
}
