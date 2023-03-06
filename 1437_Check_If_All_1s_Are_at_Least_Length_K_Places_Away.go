package leetcode

func KLengthApart(nums []int, k int) bool {
	prevOne := 0
	for ; prevOne < len(nums); prevOne++ {
		if nums[prevOne] == 1 {
			break
		}
	}

	currOne := prevOne + 1
	for ; currOne < len(nums); currOne++ {
		if nums[currOne] == 1 {
			if currOne-prevOne-1 < k {
				return false
			}
			prevOne = currOne
		}
	}
	return true
}
