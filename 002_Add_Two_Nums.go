package leetcode

import (
	gd "gitlab.com/gjerry134679/leetcode/gDataStructs"
)

func AddTwoNumbers(
	l1 *gd.Node[int],
	l2 *gd.Node[int]) (*gd.Node[int], error) {

	var GetNext func(l *gd.Node[int]) (*gd.Node[int], int)
	GetNext = func(l *gd.Node[int]) (*gd.Node[int], int) {
		if l == nil {
			return nil, 0
		}
		return l.Next(), *l.Data()
	}

	var ans, l *gd.Node[int]

	carry := 0
	for true {
		var s, d, d1, d2 int
		var err error

		l1, d1 = GetNext(l1)
		l2, d2 = GetNext(l2)

		s = d1 + d2 + carry
		if s == 0 && l1 == nil && l2 == nil {
			break
		}

		carry = s / 10
		d = s % 10

		if l == nil {
			l = gd.NewNode(d)
			ans = l
		} else {
			l, err = l.AppenToNext(d, false)
			l = l.Next()
			if err != nil {
				return l, err
			}
		}
	}

	if ans == nil {
		return gd.NewNode(0), nil
	}
	return ans, nil
}
