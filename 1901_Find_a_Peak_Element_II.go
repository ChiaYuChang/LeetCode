package leetcode

func PeakIndexInMountainArray(ary []int) int {
	lhs, rhs := 0, len(ary)

	for rhs-lhs > 32 {
		mid := (lhs + rhs) / 2

		gthl := ary[mid-1] < ary[mid]
		lthr := ary[mid] < ary[mid+1]
		if gthl && lthr {
			// increasing
			lhs = mid
		} else if !(gthl || lthr) {
			// decreasing
			rhs = mid
		} else if !gthl && lthr {
			// local minmum
			rhs = mid
		} else {
			// local maximum
			return mid
		}
	}

	// find max
	maxpos, _ := findMaxPos(lhs, rhs, ary)
	return maxpos
}

func findMaxPos(lhs, rhs int, ary []int) (pos int, val int) {
	if lhs == rhs {
		return lhs, ary[lhs]
	}
	pos = lhs + 1
	val = 0
	for i := lhs + 1; i < rhs; i++ {
		if ary[i] > val {
			val = ary[i]
			pos = i
		}
	}
	return pos, val
}
