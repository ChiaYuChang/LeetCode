package leetcode

type Q0045 struct{}

func (q Q0045) JumpSlow(nums []int) int {
	l := len(nums)

	cur := make([]int, l)
	cur[0] = 1
	for i := 0; i < l && cur[l-1] == 0; i++ {
		for j := i; j < l; j++ {
			if cur[j] > 0 {
				for k := 1; j+k < l && k <= nums[j]; k++ {
					if cur[j+k] == 0 || cur[j]+1 < cur[j+k] {
						cur[j+k] = cur[j] + 1
					}
				}
			} else {
				allzero := true
				for k := j; k < l; k++ {
					if cur[k] != 0 {
						allzero = false
						break
					}
				}
				if allzero {
					// could not reach the end
					return -1
				}
			}
		}
	}
	return cur[l-1] - 1
}

func (q Q0045) JumpFast(nums []int) int {
	l := len(nums)

	currMax := 0
	nJump := 0

	var i, j, maxFromJ int
	for i = 0; currMax < l-1 && i < l; i++ {
		j, maxFromJ = i, currMax
		for ; j < l && j <= currMax; j++ {
			if j+nums[j] > maxFromJ {
				maxFromJ = j + nums[j]
			}
		}
		currMax = maxFromJ
		i = j - 1
		nJump++
	}
	return nJump
}
