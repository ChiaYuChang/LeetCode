package leetcode

type Q0258 struct{}

func (q Q0258) AddDigitsMath(num int) int {
	if num == 0 {
		return 0
	}

	if num%9 == 0 {
		return 9
	}

	return ((num * 10) % 9) % 9
}

func (q Q0258) AddDigits(num int) int {
	for num >= 10 {
		digitSum := 0
		for num > 0 {
			digitSum += num % 10
			num /= 10
		}
		num = digitSum
	}
	return num
}
