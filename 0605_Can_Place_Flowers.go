package leetcode

type Q0605 struct{}

type Q0605PaddingFlowerBed []int

func (pf Q0605PaddingFlowerBed) Get(i int) bool {
	if i < -1 || i > len(pf) {
		return true
	} else if i == -1 || i == len(pf) {
		return false
	} else {
		return pf[i] == 1
	}
}

func (q Q0605) CanPlaceFlowers(flowerbed []int, n int) bool {
	pf := Q0605PaddingFlowerBed(flowerbed)
	var i, j, k int
	for i = -2; i <= len(pf)+1; i++ {
		if pf.Get(i) {
			for j = i + 1; j <= len(pf)+1; j++ {
				if pf.Get(j) {
					k = j - i - 3
					if k > 0 {
						n -= k/2 + k%2
					}
					break
				}
			}
			i = j - 1
		}
	}
	return n <= 0
}
