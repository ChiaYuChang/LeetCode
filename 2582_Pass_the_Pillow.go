package leetcode

type Q2582 struct{}

func (q Q2582) PassThePillow(n int, time int) int {
	if time/(n-1)%2 == 0 {
		return time%(n-1) + 1
	}
	return n - (time % (n - 1))
}
