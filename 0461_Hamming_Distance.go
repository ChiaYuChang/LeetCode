package leetcode

type Q0461 struct{}

func (q Q0461) HammingDistance(x int, y int) int {
	// x should be greater than y
	if y > x {
		x, y = y, x
	}

	n := 0
	for y != 0 {
		if x%2 != y%2 {
			n++
		}
		x, y = x/2, y/2
	}

	for x != 0 {
		if x%2 == 1 {
			n++
		}
		x = x / 2
	}
	return n
}
