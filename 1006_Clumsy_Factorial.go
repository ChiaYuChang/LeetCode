package leetcode

type Q1006 struct{}

func (q Q1006) Clumsy(n int) int {
	switch n {
	case 4:
		return 7
	case 3:
		return 6
	case 2:
		return 2
	case 1:
		return 1
	case 0:
		return 0
	}

	cf := n*(n-1)/(n-2) + (n - 3)
	n -= 4
	for n > 3 {
		cf += (n - 3) - n*(n-1)/(n-2)
		n -= 4
	}
	return cf - q.Clumsy(n)
}
