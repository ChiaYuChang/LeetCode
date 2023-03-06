package leetcode

func LexicalOrder(n int) []int {
	nums := []int{}
	for prefix := 1; prefix <= 9; prefix++ {
		// numbers should not start with 0
		_ = AppendNextDigit(prefix, &n, &nums)
	}
	return nums
}

func AppendNextDigit(prefix int, target *int, nums *[]int) bool {
	if prefix > *target {
		// stop incrementing signal
		return false
	}

	// Append prefix to result array if prefix is less than target
	(*nums) = append((*nums), prefix)
	if prefix*10 <= *target {
		// Is there any number start with the same prefix?
		for i := 0; i <= 9; i++ {
			if ok := AppendNextDigit(prefix*10+i, target, nums); !ok {
				break
			}
		}
	}

	// keep incrementing signal
	return true
}
