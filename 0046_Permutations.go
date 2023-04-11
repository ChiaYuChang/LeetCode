package leetcode

type Q0046 struct{}

func (q Q0046) Permute(nums []int) [][]int {
	p := Q0046Permuter{nums: nums, len: len(nums)}

	hu := make([]bool, len(nums))
	prfx := make([]int, len(nums))
	ans := [][]int{}
	p.Permute(0, &hu, &prfx, &ans)
	return ans
}

type Q0046Permuter struct {
	nums []int
	len  int
}

func (this Q0046Permuter) Permute(l int, hasUsed *[]bool, prefix *[]int, ans *[][]int) {
	if l < this.len {
		init := true
		for i := 0; i < this.len; i++ {
			if (*hasUsed)[i] {
				// skip element has already appended
				continue
			}

			var u []bool
			var p []int
			u = make([]bool, this.len)
			copy(u, *hasUsed)
			if !init {
				p = make([]int, this.len)
				copy(p, (*prefix)[:l])
			} else {
				// keep useing current prefix int array
				p = *prefix
				init = false
			}

			u[i] = true
			p[l] = this.nums[i]
			this.Permute(l+1, &u, &p, ans)
		}
	} else {
		(*ans) = append((*ans), *prefix)
	}
}
