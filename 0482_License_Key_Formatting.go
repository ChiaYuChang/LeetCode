package leetcode

import (
	"bytes"
	"strings"
)

type Q0482 struct{}

func (q Q0482) LicenseKeyFormattingStringBuilder(s string, k int) string {
	nAlpNum := 0
	for i := 0; i < len(s); i++ {
		if q.isAlpNum(s[i]) {
			nAlpNum++
		}
	}
	if nAlpNum < 1 {
		return ""
	}

	sb := strings.Builder{}
	j := nAlpNum % k
	if j == 0 {
		j = k
	}
	for i := 0; i < len(s); i++ {
		if q.isAlpNum(s[i]) {
			// sb.WriteByte(toUpper(s[i]))
			sb.WriteByte(s[i])
			j--
			if j == 0 {
				sb.WriteByte('-')
				j = k
			}
		} else {
			continue
		}
	}
	fs := sb.String()
	return strings.ToUpper(fs[:len(fs)-1])
}

func (q Q0482) LicenseKeyFormattingSegments(s string, k int) string {
	nAlpNum := 0
	for i := 0; i < len(s); i++ {
		if q.isAlpNum(s[i]) {
			nAlpNum++
		}
	}
	if nAlpNum < 1 {
		return ""
	}

	nSeg := nAlpNum / k
	nRes := nAlpNum % k
	if nRes != 0 {
		nSeg++
	}
	Segs := make([][]byte, nSeg)
	Segs[0] = make([]byte, 0, q.notZero(nRes, k))
	for i := 1; i < nSeg; i++ {
		Segs[i] = make([]byte, 0, k)
	}

	iSeg := 0
	for i := 0; i < len(s); i++ {
		if !q.isAlpNum(s[i]) {
			continue
		}
		if len(Segs[iSeg]) == cap(Segs[iSeg]) {
			iSeg++
		}
		Segs[iSeg] = append(Segs[iSeg], q.toUpper(s[i]))
	}

	return string(bytes.Join(Segs, []byte{'-'})[:])
}

func (q Q0482) isAlpNum(b byte) bool {
	return q.isNumber(b) || q.isLowerLetter(b) || q.isUpperLetter(b)
}

func (q Q0482) isNumber(b byte) bool {
	return '0' <= b && b <= '9'
}

func (q Q0482) isLowerLetter(b byte) bool {
	return 'a' <= b && b <= 'z'
}

func (q Q0482) isUpperLetter(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

func (q Q0482) toUpper(b byte) byte {
	if q.isLowerLetter(b) {
		return b - 'a' + 'A'
	}
	return b
}

func (q Q0482) notZero(x, d int) int {
	if x == 0 {
		return d
	}
	return x
}
