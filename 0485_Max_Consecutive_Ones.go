package leetcode

type Q0485 struct{}

func (q Q0485) FindMaxConsecutiveOnes(nums []int) int {
	mCO := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			nOne := 1
			j := i + 1
			for ; j < len(nums) && nums[j] == 1; j++ {
				nOne++
			}
			if nOne > mCO {
				mCO = nOne
			}
			i = j
		}
	}
	return mCO
}
