package leetcode

type Q2427 struct{}

func (q Q2427) CommonFactors(a int, b int) int {
	if a > b {
		a, b = b, a
	}

	ncf := 1
	for i, ub := 2, a/2; i <= ub; i++ {
		n := 1
		for a%i == 0 && b%i == 0 {
			n, a, b = n+1, a/i, b/i
		}
		ncf *= n
	}
	return ncf
}
