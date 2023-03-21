package leetcode

type Q0238 struct{}

func (q Q0238) ProductExceptSelf(nums []int) []int {
	all := 1
	haszero := false
	zeropos := 0
	flag := 0b00

	for i, n := range nums {
		if n == 0 {
			if haszero {
				// more than one zero in nums
				flag = 0b10
				break
			}
			haszero = true
			zeropos = i // the position of the first zero in nums
			flag = 0b01
		} else {
			all *= n
		}
	}

	prod := make([]int, len(nums))
	switch flag {
	case 0b00:
		for i := range nums {
			prod[i] = all / nums[i]
		}
	case 0b01:
		prod[zeropos] = all
	}
	return prod
}
