package leetcode

func AddToArrayForm(num []int, k int) []int {
	var ext []int

	var val, digit, carry int
	for i := len(num) - 1; k > 0 || carry != 0; i-- {
		digit = k % 10
		k /= 10

		if i >= 0 {
			val = num[i] + digit + carry
			num[i] = val % 10
		} else {
			val = digit + carry
			ext = append(ext, val%10)
		}
		carry = val / 10
	}

	if len(ext) > 0 {
		ans := make([]int, len(ext)+len(num))
		i := 0
		for ; i < len(ext); i++ {
			ans[i] = ext[len(ext)-i-1]
		}
		copy(ans[i:], num)
		return ans
	}
	return num
}
