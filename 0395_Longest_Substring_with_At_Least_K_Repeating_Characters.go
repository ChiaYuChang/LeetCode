package leetcode

type Q0395 struct{}

func (q Q0395) LongestSubstringRecursive(s string, k int) int {
	if len(s) < k {
		return 0
	}

	lk := [26][]int{}
	for i := 0; i < len(s); i++ {
		lk[s[i]-'a'] = append(lk[s[i]-'a'], i)
	}

	nb := 0
	bm := make([]bool, len(s))
	for _, i := range lk {
		if len(i) < k {
			for _, j := range i {
				bm[j] = true
				nb++
			}
		}
	}
	if nb == 0 {
		return len(s)
	}

	mssl := 0
	for i := 0; i < len(s); i++ {
		if !bm[i] {
			ssl, j := 1, i+1
			for ; j < len(s) && !bm[j]; j++ {
				ssl++
			}
			ssl = q.LongestSubstringRecursive(s[i:j], k)
			if ssl > mssl {
				mssl = ssl
			}
		}
	}
	return mssl
}

func (q Q0395) LongestSubstringSlidingWindows(s string, k int) int {
	if len(s) < k {
		return 0
	}

	maxLen := 0
	c := make(chan int)
	for i := 1; i <= 26; i++ {
		go q.longestSubstringWithTSymbol(s, i, k, c)
	}

	for i := 1; i <= 26; i++ {
		if l := <-c; l > maxLen {
			maxLen = l
		}
	}
	close(c)
	return maxLen
}

func (q Q0395) longestSubstringWithTSymbol(s string, T int, k int, c chan<- int) {
	begin, end := 0, 0
	sFreq := make([]int, 26)
	nSmbl := 0
	nFreqGtK := 0
	maxLen := 0

	for end < len(s) {
		if nSmbl <= T {
			if sFreq[s[end]-'a'] == 0 {
				// new symbol
				nSmbl++
			}
			sFreq[s[end]-'a']++
			if sFreq[s[end]-'a'] == k {
				nFreqGtK++
			}
			end++
		} else if nSmbl > T {
			if sFreq[s[begin]-'a'] == k {
				nFreqGtK--
			}
			sFreq[s[begin]-'a']--
			if sFreq[s[begin]-'a'] == 0 {
				// remove one symbol
				nSmbl--
			}
			begin++
		}

		if nSmbl == T && nFreqGtK == T && end-begin > maxLen {
			maxLen = end - begin
		}
	}
	c <- maxLen
}
