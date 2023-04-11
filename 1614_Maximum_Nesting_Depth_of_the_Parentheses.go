package leetcode

type Q1614 struct{}

func (q Q1614) MaxDepth(s string) int {
	mD := 0
	for lhs := 0; lhs < len(s); lhs++ {
		if s[lhs] == '(' {
			rhs := lhs + 1
			for lvl := 1; rhs < len(s); rhs++ {
				if s[rhs] == ')' {
					lvl--
				} else if s[rhs] == '(' {
					lvl++
				}

				if lvl == 0 {
					break
				}
			}
			if d := 1 + q.MaxDepth(s[lhs+1:rhs]); d > mD {
				mD = d
			}
			lhs = rhs
		}
	}
	return mD
}
