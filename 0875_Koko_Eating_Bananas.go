package leetcode

type Q0875 struct{}

func (q Q0875) MinEatingSpeed(piles []int, h int) int {
	if h < len(piles) {
		return -1
	}

	var lhs, mid, rhs int
	lhs = 1
	rhs = q.max(piles)

	for rhs > lhs {
		mid = (rhs + lhs) / 2
		if q.canEatAll(piles, mid, h) {
			rhs = mid
		} else {
			lhs = mid + 1
		}
	}
	return lhs
}

func (q Q0875) canEatAll(piles []int, speed, limitedTime int) bool {
	consumedTime := 0
	for _, p := range piles {
		if p <= speed {
			consumedTime += 1
		} else {
			consumedTime += p / speed
			if p%speed > 0 {
				consumedTime++
			}
		}
	}
	return consumedTime <= limitedTime
}

func (q Q0875) max(nums []int) int {
	m := -1
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}
