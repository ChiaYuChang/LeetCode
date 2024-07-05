package leetcode

type Q1915 struct {
}

func (q Q1915) WonderfulSubstrings(word string) int64 {
	mask := map[uint16]int64{0: 1}

	nWd4Substr := int64(0)
	curMask := uint16(0)
	for i := 0; i < len(word); i++ {
		// add substring with all letter appears even times ending at i
		curMask = q.ExclusiveNot(curMask, 1<<(word[i]-'a'))
		nWd4Substr += mask[curMask]
		mask[curMask] = mask[curMask] + 1

		// find mask with only one bit different from curMask
		// if curMask equals to 0001, then
		// target masks are 0010, 0100, 1000, 0000
		for j := 0; j < len(q.letters()); j++ {
			// add substring with only one letter appears odd times ending at i
			nWd4Substr += mask[q.ExclusiveNot(curMask, 1<<(q.letters()[j]-'a'))]
		}
	}
	return nWd4Substr
}

// reture letters that allowed to appear in the word
func (q Q1915) letters() string {
	return "abcdefghij"
}

// return the distance between x and y
func (q Q1915) ExclusiveNot(x, y uint16) uint16 {
	//   | 0 | 1 |
	// --|---|---|
	// 0 | 0 | 1 |
	// --|---|---|
	// 1 | 1 | 0 |
	// --|---|---|
	return (x | y) & ^(x & y)
}
