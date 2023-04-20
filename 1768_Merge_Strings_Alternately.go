package leetcode

type Q1768 struct{}

func (q Q1768) MergeAlternately(word1 string, word2 string) string {
	ll, ls := len(word1), len(word2)
	hasSwapped := false
	if ls > ll {
		word1, word2 = word2, word1
		ll, ls = ls, ll
		hasSwapped = true
	}

	mWord := make([]rune, ll+ls)
	i, j := 0, 0
	for ; i < ls; i, j = i+1, j+2 {
		bl, bs := rune(word1[i]), rune(word2[i])
		if hasSwapped {
			mWord[j], mWord[j+1] = bs, bl
		} else {
			mWord[j], mWord[j+1] = bl, bs
		}
	}

	for ; i < ll; i, j = i+1, j+1 {
		mWord[j] = rune(word1[i])
	}

	return string(mWord)
}
