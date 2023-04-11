package leetcode

type Q0492 struct{}

func (q Q0492) ConstructRectangle(area int) []int {
	lw := []int{area, 1}
	for i := 2; i*i <= area; i++ {
		if area%i == 0 {
			lw[0], lw[1] = area/i, i
		}
	}
	return lw
}
