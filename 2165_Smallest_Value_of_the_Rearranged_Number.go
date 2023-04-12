package leetcode

type Q2165 struct{}

func (q Q2165) SmallestNumber(num int64) int64 {
	sign := num > 0
	if !sign {
		num = -num
	}

	count := [10]int64{}
	for num > 0 {
		count[num%10]++
		num /= 10
	}

	if sign {
		// num > 0
		var i int64 = 1
		for ; i < 10; i++ {
			if count[i] > 0 {
				num += i
				count[i]--
				break
			}
		}

		for count[0] > 0 {
			num *= 10
			count[0]--
		}

		for ; i < 10; i++ {
			for count[i] > 0 {
				num = num*10 + i
				count[i]--
			}
		}
	} else {
		// num <= 0
		var i int64 = 9
		for ; i >= 0; i-- {
			for count[i] > 0 {
				num = num*10 - i
				count[i]--
			}
		}
	}
	return num
}
