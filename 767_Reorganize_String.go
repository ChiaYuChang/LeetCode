package leetcode

type nchar struct {
	char  byte
	count int
}

func ReorganizeString(s string) string {
	cs := make([]*nchar, 26)
	for i := 0; i < 26; i++ {
		cs[i] = &nchar{char: byte('a' + i), count: 0}
	}

	for i := 0; i < len(s); i++ {
		cs[s[i]-'a'].count += 1
	}

	ans := make([]byte, len(s))
	for i := range ans {
		ans[i] = '.'
	}

	// find mode char
	whichmax := 0
	max := 0
	for i := 0; i < 26; i++ {
		if cs[i].count > max {
			whichmax = i
			max = cs[i].count
		}
	}
	// always fill the slide first with mode
	cs[0], cs[whichmax] = cs[whichmax], cs[0]

	cchar := 0
	for i := 0; i < 2; i++ {
		for j := i; j < len(s) && cchar < 26; j += 2 {
			ans[j] = cs[cchar].char
			if j > 0 && ans[j] == ans[j-1] {
				return ""
			}
			cs[cchar].count--
			for cchar < 26 && cs[cchar].count < 1 {
				cchar++
			}
		}
	}
	return string(ans)
}
