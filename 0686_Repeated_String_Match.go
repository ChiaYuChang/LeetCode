package leetcode

import "strings"

type Q0686 struct{}

func (q Q0686) RepeatedStringMatch(a string, b string) int {
	if len(a) <= len(b) {
		n := 0

		// look for the starting site
		f := false
		for i := 0; i < len(a); i++ {
			if a[i] == b[0] {
				if a[i:] == b[:len(a)-i] {
					b = b[len(a)-i:]
					n = n + 1
					f = true
					break
				}
			}
		}

		// can't find the starting site
		if !f {
			return -1
		}

		// trim b step by step
		for len(b) > len(a) {
			if b[:len(a)] == a {
				n++
				b = b[len(a):]
			} else {
				// pattern not match
				return -1
			}
		}

		// compare the residues
		if len(b) > 0 {
			for i := 0; i < len(b); i++ {
				if b[i] != a[i] {
					return -1
				}
			}
			n++
		}
		return n
	} else {
		// a = "aa" b = "a"
		if strings.Contains(a, b) {
			return 1
		}

		// a = "aaaab" b = "aba"
		if strings.Contains(a+a, b) {
			return 2
		}
		return -1
	}
}
