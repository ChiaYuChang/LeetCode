package leetcode

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	preffix := strings.Builder{}
	for i := 0; i < len(strs[0]); i++ {
		prfx := strs[0][i]
		for _, s := range strs[1:] {
			if getIdx(i, s) != prfx {
				return preffix.String()
			}
		}
		preffix.WriteByte(prfx)
	}
	return preffix.String()
}

func getIdx(i int, str string) byte {
	if i >= len(str) {
		return '#'
	}
	return str[i]
}
