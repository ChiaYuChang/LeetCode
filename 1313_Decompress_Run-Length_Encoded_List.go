package leetcode

type Q1313 struct{}

func (q Q1313) DecompressRLElist(nums []int) []int {
	// pre allocate slice memory
	n := 0
	for i := 0; i < len(nums); i += 2 {
		n += nums[i]
	}
	d := make([]int, n)

	// fill in values
	n = 0
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			d[n] = nums[i+1]
			n++
		}
	}
	return d
}
