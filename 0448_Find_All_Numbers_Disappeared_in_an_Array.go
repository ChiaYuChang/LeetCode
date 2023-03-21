package leetcode

type Q0448 struct{}

func (q Q0448) FindDisappearedNumbers(nums []int) []int {
	missNums := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}

	for i, num := range nums {
		if i != num-1 {
			missNums = append(missNums, i+1)
		}
	}
	return missNums
}

func (q Q0448) FindDisappearedNumbersSlow(nums []int) []int {
	n := len(nums)

	bitAry := make([]bool, n+1)
	nMissNums := n

	for _, num := range nums {
		if !bitAry[num] {
			nMissNums--
		}
		bitAry[num] = true
	}

	missNums := make([]int, 0, nMissNums)
	for i := 1; i <= n; i++ {
		if !bitAry[i] {
			missNums = append(missNums, i)
		}
	}
	return missNums
}
