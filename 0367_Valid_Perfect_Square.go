package leetcode

type Q0367 struct{}

func (q Q0367) IsPerfectSquare(num int) bool {
	pre, cur := 0, 1
	for {
		if cur*cur < num {
			if pre == cur {
				return false
			}
			pre = cur
			cur = cur * 2
		} else if cur*cur > num {
			cur = (pre + cur) / 2
		} else {
			return true
		}
	}
}
