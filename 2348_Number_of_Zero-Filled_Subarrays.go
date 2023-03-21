package leetcode

type Q2848 struct{}

type Q2848PaddingIntArray []int

func (pa Q2848PaddingIntArray) Len() int64 {
	return int64(len(pa))
}

func (pa Q2848PaddingIntArray) IsZeroAt(i int64) bool {
	if i >= pa.Len() {
		return false
	}
	return pa[i] == 0
}

func (q Q2848) ZeroFilledSubarray(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}

	pnums := Q2848PaddingIntArray(nums)
	var i, j, n int64
	for i = 0; i <= pnums.Len(); i++ {
		if pnums.IsZeroAt(i) {
			for j = i + 1; j <= pnums.Len(); j++ {
				if !pnums.IsZeroAt(j) {
					n += (j - i) * (j - i + 1) / 2
					break
				}
			}
			i = j
		}
	}
	return int64(n)
}

func (q Q2848) FindZeroSegLen(nums []int) []int {
	endWithZero := nums[len(nums)-1] == 0
	if endWithZero {
		nums[len(nums)-1] = 1
	}

	segs := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			var j int
			for j = i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					segs = append(segs, j-i)
					break
				}
			}
			i = j
		}
	}

	if endWithZero {
		if len(segs) > 0 {
			nums[len(nums)-1] = 0
			segs[len(segs)-1] += 1
		} else {
			segs = append(segs, 1)
		}
	}
	return segs
}
