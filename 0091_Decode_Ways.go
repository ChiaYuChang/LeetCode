package leetcode

type Q0091 struct{}

type Q0091Token int

func (q Q0091) NewToken(s0, s1 byte) Q0091Token {
	return Q0091Token(int(s0-'0')*10 + int(s1-'0'))
}

func (n1 Q0091Token) Mergeable(n2 Q0091Token) bool {
	if n1 > 9 || n2 > 9 || n1*10+n2 > 26 {
		return false
	}
	return true
}

func (q Q0091) NumDecodings(s string) int {
	if !q.CheckFormat(s) {
		return 0
	}

	// split string into Num objects
	tokens := make([]Q0091Token, 1, len(s)+1)
	tokens[0] = 'z' + 1 // padding
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			tokens = append(tokens, q.NewToken(s[i-1], s[i]))
			i++
		} else {
			tokens = append(tokens, q.NewToken('0', s[i-1]))
		}
	}

	// append the last segments
	if s[len(s)-1] != '0' {
		tokens = append(tokens, q.NewToken('0', s[len(s)-1]))
	}

	nbranch := [3]int{1, 1, 1}
	for i := len(tokens) - 3; i >= 0; i-- {
		if tokens[i+1].Mergeable(tokens[i+2]) {
			nbranch[0] = nbranch[1] + nbranch[2]
		} else {
			nbranch[0] = nbranch[1]
		}
		nbranch[1], nbranch[2] = nbranch[0], nbranch[1]
	}
	return nbranch[0]
}

func (q Q0091) CheckFormat(s string) (ok bool) {
	if len(s) == 0 || s[0] == '0' {
		// should not start by 0
		return false
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			// only 2x, 1x is valid
			// 00 is not allowed
			if (s[i-1] > '2') || i+1 < len(s) && s[i+1] == '0' {
				return false
			}
			i++
		}
	}
	return true
}
