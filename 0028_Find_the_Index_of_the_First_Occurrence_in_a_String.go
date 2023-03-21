package leetcode

type Q0028 struct{}

func (q Q0028) StrStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}

	l := len(needle)
	if l == 0 {
		return 0
	}

	for i := 0; i < (len(haystack) - l + 1); i++ {
		if haystack[i:i+l] == needle {
			return i
		}
	}
	return -1
}
