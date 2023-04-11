package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

type WordDictionaryActionType int8

const (
	WDActNew WordDictionaryActionType = iota
	WDActAddWord
	WDActSearch
)

type WordDictionaryAction struct {
	Action   WordDictionaryActionType
	Argument string
}

func NewWordDictionaryAction(action, arg string) WordDictionaryAction {
	a := WordDictionaryAction{}
	switch action {
	case "WordDictionary":
		a.Action = WDActNew
	case "addWord":
		a.Action = WDActAddWord
	case "search":
		a.Action = WDActSearch
	}
	a.Argument = arg
	return a
}

func NewWordDictionaryActions(actions, args []string) []WordDictionaryAction {
	as := make([]WordDictionaryAction, len(actions))
	for i, a := range actions {
		as[i] = NewWordDictionaryAction(a, args[i])
	}
	return as
}

func TestWordDictionary(t *testing.T) {

	type testCase struct {
		actions []string
		args    []string
		ans     []bool
	}

	tcs := []testCase{
		{
			actions: []string{
				"WordDictionary", "addWord", "search", "addWord", "addWord",
				"addWord", "search", "search", "search", "search",
			},
			args: []string{
				"", "badd", "bad", "bad", "dad",
				"mad", "pad", "bad", ".ad", "b..",
			},
			ans: []bool{
				false, false, false, false, false,
				false, false, true, true, true,
			},
		},
	}

	q := leetcode.Q0211{}
	for i := range tcs {
		tc := tcs[i]
		as := NewWordDictionaryActions(tc.actions, tc.args)
		var wd leetcode.Q0211WordDictionary
		for i, a := range as {
			switch a.Action {
			case WDActNew:
				wd = q.Constructor()
			case WDActAddWord:
				wd.AddWord(a.Argument)
				// fmt.Println(wd)
			case WDActSearch:
				ok := wd.Search(a.Argument)
				require.Equal(t, tc.ans[i], ok)
			}
		}
	}
}
