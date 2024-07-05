package leetcode

type Q1417 struct{}

func (q Q1417) Reformat(s string) string {
	nLong, nShrt := 0, 0
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			nLong++
		} else {
			nShrt++
		}
	}

	dLead := true
	if nLong < nShrt {
		nLong, nShrt = nShrt, nLong
		dLead = false
	}

	if nLong-nShrt > 1 {
		return ""
	}

	p := make([]byte, len(s))
	iDigit, iAlpha := 0, 1
	if !dLead {
		iDigit, iAlpha = 1, 0
	}

	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			p[iDigit] = s[i]
			iDigit += 2
		} else {
			p[iAlpha] = s[i]
			iAlpha += 2
		}
	}
	return string(p)
}
