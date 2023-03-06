package leetcode

type Q409 struct{}

func (q Q409) LongestPalindrome(s string) int {
	counter := make([]int, 26*2)
	for i := 0; i < len(s); i++ {
		if j := q.Byte2Int(s[i]); j >= 0 {
			counter[j] += 1
		}
	}

	maxPalindromeLen := 0
	hasOdd := false
	for i := 0; i < len(counter); i++ {
		if counter[i]%2 == 0 {
			maxPalindromeLen += counter[i]
		} else {
			hasOdd = true
			maxPalindromeLen += counter[i] - 1
		}
	}

	if hasOdd {
		maxPalindromeLen += 1
	}
	return maxPalindromeLen
}

func (q Q409) Byte2Int(b byte) int {
	if 'a' <= b && b <= 'z' {
		return int(b - 'a')
	}

	if 'A' <= b && b <= 'Z' {
		return int(b-'A') + 26
	}
	return -1
}
