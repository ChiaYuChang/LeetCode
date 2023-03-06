package leetcode

const shift = 'a' - 'A'

func LetterCasePermutation(s string) []string {
	empty := make([]byte, len(s))
	ans := []*[]byte{&empty}

	for i := 0; i < len(s); i++ {
		c1, c2, ok := Substitute(s[i])
		if ok {
			for j := range ans {
				src := *ans[j]
				dst := make([]byte, len(src))
				copy(dst, src)
				src[i] = c1
				dst[i] = c2
				ans = append(ans, &dst)
			}
		} else {
			for j := range ans {
				(*ans[j])[i] = c1
			}
		}
	}
	return ListBytesToStrings(ans)
}

func Substitute(c byte) (byte, byte, bool) {
	if 'A' <= c && c <= 'Z' {
		return c, c + shift, true
	}
	if 'a' <= c && c <= 'z' {
		return c, c - shift, true
	}
	return c, c, false
}

func ListBytesToStrings(bs []*[]byte) []string {
	rst := make([]string, len(bs))
	for i := range bs {
		rst[i] = string((*bs[i]))
	}
	return rst
}
