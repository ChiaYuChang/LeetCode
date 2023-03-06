package leetcode_test

import (
	"testing"

	"gitlab.com/gjerry134679/leetcode"
)

func TestMiniParserLex(t *testing.T) {
	expr := "[123,[456,[789]]]"

	p := leetcode.MiniExprParser{}
	tkns := p.Lext(expr)
	t.Log(tkns)
}
