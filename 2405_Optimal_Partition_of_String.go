package leetcode

type Q2405 struct{}

func (q Q2405) PartitionString(s string) int {
	ul := [26]bool{} // used letters
	nSeg := 1
	// finding cutting sites
	for i := 0; i < len(s); i++ {
		if ul[s[i]-'a'] {
			nSeg += 1
			ul = [26]bool{}
		}
		ul[s[i]-'a'] = true
	}
	return nSeg
}
