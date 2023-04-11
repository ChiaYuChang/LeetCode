package leetcode

type Q0633 struct{}

func (q Q0633) JudgeSquareSum(c int) bool {
	var isPerfectSquare func(num int) (bool, int) = func(num int) (bool, int) {
		pre, cur := 0, 1
		for {
			if cur*cur < num {
				if pre == cur {
					return false, pre
				}
				pre = cur
				cur = cur * 2
			} else if cur*cur > num {
				cur = (pre + cur) / 2
			} else {
				return true, pre
			}
		}
	}

	lhs := 0
	ok, rhs := isPerfectSquare(c)

	if ok {
		return true
	}

	rhs++
	for lhs <= rhs {
		tmp := lhs*lhs + rhs*rhs
		if tmp > c {
			rhs--
		} else if tmp < c {
			lhs++
		} else {
			return true
		}
	}
	return false
}
