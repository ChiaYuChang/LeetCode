package leetcode

type Q0319 struct{}

func (q Q0319) BulbSwitch(n int) int {
	NSquareNumber := 0
	for i := 1; i*i <= n; i++ {
		NSquareNumber++
	}
	return NSquareNumber
}
