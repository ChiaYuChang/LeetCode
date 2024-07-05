package leetcode

import (
	"sort"
	"strings"
)

type Q0451 struct{}

func (q Q0451) FrequencySort(s string) string {
	count := make([][2]int, 62)
	for i := 0; i < 62; i++ {
		count[i][0] = i
	}

	for i := 0; i < len(s); i++ {
		count[q.Byte2Index(s[i])][1] += 1
	}

	sort.Slice(count, func(i, j int) bool {
		return count[i][1] >= count[j][1]
	})

	sb := strings.Builder{}
	for _, c := range count {
		sb.WriteString(strings.Repeat(string(q.Index2Char(c[0])), c[1]))
	}
	return sb.String()
}

func (q Q0451) Byte2Index(b byte) int {
	if 'A' <= b && b <= 'Z' {
		return int(b - 'A')
	} else if 'a' <= b && b <= 'z' {
		return int(b-'a') + 26
	} else {
		return int(b-'0') + 52
	}
}

func (q Q0451) Index2Char(i int) byte {
	if i < 26 {
		return byte(i + 'A')
	} else if i < 52 {
		return byte(i - 26 + 'a')
	} else {
		return byte(i - 52 + '0')
	}
}
