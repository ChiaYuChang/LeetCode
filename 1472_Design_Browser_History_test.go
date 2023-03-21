package leetcode_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gjerry134679/leetcode"
)

func TestBrowserHistory(t *testing.T) {
	type Action struct {
		Type      rune
		Arguement string
		Output    string
	}

	type testCase struct {
		action []*Action
	}

	tcs := []testCase{
		{
			[]*Action{
				{'n', "leetcode.com", ""},
				{'v', "google.com", ""},
				{'v', "facebook.com", ""},
				{'v', "youtube.com", ""},
				{'b', "1", "facebook.com"},
				{'b', "1", "google.com"},
				{'f', "1", "facebook.com"},
				{'v', "linkedin.com", ""},
				{'f', "2", "linkedin.com"},
				{'b', "2", "google.com"},
				{'b', "7", "leetcode.com"},
			},
		},
		{
			[]*Action{
				{'n', "homepage.com", ""},
				{'v', "01.com", ""},
				{'v', "02.com", ""},
				{'v', "03.com", ""},
				{'b', "1", "02.com"},
				{'b', "3", "homepage.com"},
				{'f', "1", "01.com"},
				{'v', "04.com", ""},
				{'v', "05.com", ""},
				{'v', "06.com", ""},
				{'v', "07.com", ""},
				{'v', "08.com", ""},
				{'v', "09.com", ""},
				{'v', "10.com", ""},
				{'b', "10", "homepage.com"},
				{'f', "10", "10.com"},
				{'b', "3", "07.com"},
			},
		},
	}

	q := leetcode.Q1472{}
	for i := range tcs {
		tc := tcs[i]
		t.Run(
			fmt.Sprintf("Case %d", i+1),
			func(t *testing.T) {
				var bh leetcode.Q1472BrowserHistory
				for _, a := range tc.action {
					t.Log(bh.String())
					switch a.Type {
					case 'n':
						bh = q.Constructor(a.Arguement)
					case 'v':
						bh.Visit(a.Arguement)
					case 'b':
						i, _ := strconv.Atoi(a.Arguement)
						require.Equal(t, a.Output, bh.Back(i))
					case 'f':
						i, _ := strconv.Atoi(a.Arguement)
						require.Equal(t, a.Output, bh.Forward(i))
					}
				}
			},
		)
	}
}
