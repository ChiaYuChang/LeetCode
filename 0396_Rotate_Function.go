package leetcode

type Q0396 struct{}

func (q Q0396) MaxRotateFunction(nums []int) int {
	f, total := 0, 0
	for i, num := range nums {
		total += num
		f += i * num // f(0)
	}

	max := f
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		// total = a_0 + a_1 + ... a_{n-1}
		// f(i+1) + total = f(i) + n * a_{n-i}
		f = f + total - n*nums[i]
		if f > max {
			max = f
		}
	}
	return max
}
