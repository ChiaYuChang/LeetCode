package leetcode

import (
	"strings"
)

type Q1021 struct{}

func RemoveOuterParentheses(s string) string {
	var lvl, lhs int

	sb := strings.Builder{}
	for rhs := 0; rhs < len(s); rhs++ {
		if s[rhs] == '(' {
			lvl++
		} else {
			lvl--
		}

		if lvl == 0 {
			sb.WriteString(s[lhs+1 : rhs])
			lhs = rhs + 1
		}
	}
	return sb.String()
}
