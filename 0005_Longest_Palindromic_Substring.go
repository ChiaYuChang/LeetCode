package leetcode

import gds "gitlab.com/gjerry134679/leetcode/gDataStructs"

type Q0005 struct{}

func (q Q0005) LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var evenHalfLen, evenHalfPos int
	var oddHalfLen, oddHalfPos int

	// dynamic programming
	// preRow := make([]int, len(s))
	preRow := gds.NewSparseArray[int](len(s))
	for j := len(s) - 1; j >= 0; j-- {
		// curRow := make([]int, len(s))
		curRow := gds.NewSparseArray[int](len(s))

		for i := 0; i <= j; i++ {
			if s[i] == s[j] {
				if i > 0 {
					// curRow[i] = 1 + preRow[i-1]
					curRow.MustSetVal(i, 1+preRow.MustGetVal(i-1))
				} else {
					// curRow[i] = 1
					curRow.MustSetVal(i, 1)
				}
			}
		}

		// if j > 0 && curRow[j-1] > evenHalfLen {
		// evenHalfLen = curRow[j-1]
		if curRow.MustGetVal(j-1) > evenHalfLen {
			evenHalfLen = curRow.MustGetVal(j - 1)
			evenHalfPos = j - 1
		}

		// if curRow[j] > oddHalfLen {
		// 	oddHalfLen = curRow[j]
		if curRow.MustGetVal(j) > oddHalfLen {
			oddHalfLen = curRow.MustGetVal(j)
			oddHalfPos = j
		}

		preRow = curRow
	}

	if (oddHalfLen*2 - 1) > (evenHalfLen * 2) {
		return s[(oddHalfPos - oddHalfLen + 1):(oddHalfPos + oddHalfLen)]
	}
	return s[(evenHalfPos - evenHalfLen + 1):(evenHalfPos + evenHalfLen + 1)]
}

func (q Q0005) FastLongestPalindrome(s string) string {
	var ans string

	for i := 0; i < len(s); i++ {
		l, r := i, i

		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > len(ans) {
				ans = s[l : r+1]
			}
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > len(ans) {
				ans = s[l : r+1]
			}
			l--
			r++
		}
	}
	return ans
}
