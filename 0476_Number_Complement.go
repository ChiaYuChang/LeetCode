package leetcode

type Q0476 struct{}

func (q Q0476) FindComplement(num int) int {
	base := 1
	m := 0
	for num > 0 {
		if num%2 == 0 {
			m += base
		}
		num = num >> 1
		base = base << 1
	}
	return m
}
