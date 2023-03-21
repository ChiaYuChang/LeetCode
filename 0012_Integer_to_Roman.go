package leetcode

import "strings"

// Q0012 and Q0013
type RomanIntConverter struct{}

func (q RomanIntConverter) IntToRoman(num int) string {
	if num == 0 {
		return ""
	}

	i2r := struct {
		is []int
		ss []string
	}{
		is: []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1},
		ss: []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"},
	}

	sb := strings.Builder{}
	for j, i := range i2r.is {
		if num == 0 {
			break
		}

		r := num / i
		if r > 0 {
			sb.WriteString(strings.Repeat(i2r.ss[j], num/i))
			num -= i * r
		}
	}
	return sb.String()
}
