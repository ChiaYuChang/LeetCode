package leetcode

func Permute(nums []int, n int) [][]int {
	hu := make([]bool, len(nums))
	prfx := make([]int, n)
	ans := [][]int{}
	permute(nums, 0, n, &hu, &prfx, &ans)
	return ans
}

func permute(nums []int, l int, n int, hasUsed *[]bool, prefix *[]int, ans *[][]int) {
	if l < n {
		init := true
		// fmt.Printf("%s%d: %v\n", strings.Repeat("  ", l), l, *prefix)
		for i := 0; i < len(nums); i++ {
			if (*hasUsed)[i] {
				// skip element has already appended
				continue
			}

			var u []bool
			var p []int
			u = make([]bool, len(nums))
			copy(u, *hasUsed)
			if !init {
				p = make([]int, n)
				copy(p, (*prefix)[:l])
			} else {
				// keep useing current prefix int array
				p = *prefix
				init = false
			}

			u[i] = true
			p[l] = nums[i]
			permute(nums, l+1, n, &u, &p, ans)
		}
	} else {
		(*ans) = append((*ans), *prefix)
	}
}
