package leetcode

type Q0459 struct{}

func (q Q0459) RepeatedSubstringPattern(s string) bool {
	l := len(s)

	for i := l / 2; i >= 1; i-- {
		if l%i != 0 {
			continue
		}

		f := true
		for j := i; j < l; j += i {
			if s[:i] != s[j:j+i] {
				f = false
				break
			}
		}
		if f {
			return true
		}
	}
	return false
}
