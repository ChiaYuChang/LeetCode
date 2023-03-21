package leetcode

type Q0383 struct{}

func (q Q0383) CanConstruct(ransomNote string, magazine string) bool {
	rMagazine := make([]int, 26)

	for i := range magazine {
		rMagazine[i-'A'] += 1
	}

	for i := range ransomNote {
		j := ransomNote[i] - 'A'
		rMagazine[j] -= 1
		if rMagazine[j] < 0 {
			return false
		}
	}
	return true
}
