package leetcode

type Q0763 struct{}

func (q Q0763) PartitionLabels(s string) []int {
	// the last postion of each letter
	lastPos := [26]int{}
	for i := 0; i < len(s); i++ {
		lastPos[s[i]-'a'] = i + 1
	}

	segLen := []int{}
	var segStr, segEnd int // default 0, 0
	for segEnd < len(s) {
		segStr, segEnd = segEnd, lastPos[s[segEnd]-'a']
		for i := segStr; i < segEnd; i++ {
			if segEnd < lastPos[s[i]-'a'] {
				segEnd = lastPos[s[i]-'a']
			}
		}
		segLen = append(segLen, segEnd-segStr)
	}
	return segLen
}
