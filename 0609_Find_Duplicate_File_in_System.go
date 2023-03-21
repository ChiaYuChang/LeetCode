package leetcode

import (
	"fmt"
	"strings"
)

type Q0609 struct{}

func (q Q0609) FindDuplicate(paths []string) [][]string {
	content := map[string][]string{}

	for _, path := range paths {
		ss := strings.Split(path, " ")
		p := ss[0]
		for i := 1; i < len(ss); i++ {
			cs := strings.Index(ss[i], "(")
			fn, ct := ss[i][0:cs], ss[i][cs+1:len(ss[i])-1]
			content[ct] = append(content[ct], fmt.Sprintf("%s/%s", p, fn))
		}
	}

	result := make([][]string, 0)
	for _, vals := range content {
		if len(vals) > 1 {
			result = append(result, vals)
		}
	}
	return result
}
