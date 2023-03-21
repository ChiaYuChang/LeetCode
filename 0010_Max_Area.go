package leetcode

type Q0010 struct{}

func (q Q0010) MaxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	l, r := 0, len(height)-1
	calcArea := func(left, right, height int) int {
		if left > right {
			return right * height
		}
		return left * height
	}

	maxarea := calcArea(height[l], height[r], r-l)
	for l < r {
		// determine direction
		if height[l] < height[r] {
			// l -->
			i := l + 1
			for height[l] > height[i] && i < r {
				i++
			}

			if i > r {
				break
			}

			a := calcArea(height[i], height[r], r-i)
			if a > maxarea {
				maxarea = a
			}
			l = i
		} else {
			// <--- r
			i := r - 1
			for height[r] > height[i] && i > l {
				i--
			}

			if i < l {
				break
			}

			a := calcArea(height[l], height[i], i-l)
			if a > maxarea {
				maxarea = a
			}
			r = i
		}
	}
	return maxarea
}
