package leetcode

type Q0006 struct{}

func (q Q0006) ZigzagConversion(s string, numRows int) string {
	if s == "" || numRows <= 1 {
		return s
	}

	ans := make([]byte, len(s))
	k := 0

	l, r := 2*(numRows-1), 0
	for n := 0; n < numRows && l >= 0; n++ {
		c := n
		if l == 0 || r == 0 {
			for ; c < len(s); k++ {
				ans[k] = s[c]
				c += l + r
			}
		} else {
			for f := true; c < len(s); k++ {
				ans[k] = s[c]
				if f {
					c += l
				} else {
					c += r
				}
				f = !f
			}
		}
		l -= 2
		r += 2
	}
	return string(ans)
}
