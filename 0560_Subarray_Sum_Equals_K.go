package leetcode

import "fmt"

type Q0560 struct{}

func SubarraySum(nums []int, k int) int {
	nSumK := 0
	outMaxItr := 20
	for start, sum, i := 0, nums[0], 0; outMaxItr > 0 && i < len(nums); outMaxItr-- {
		maxItr := 10
		fmt.Printf("A [%d, %d] sum: %d\n", start, i, sum)
		for maxItr > 0 && sum <= k && i < len(nums) {
			i++
			sum += nums[i]
			fmt.Printf("\t-Extent [%d, %d] sum: %2d nSumK: %d", start, i, sum, nSumK)
			if sum == k {
				fmt.Println("...Hit!!")
				nSumK += 1
			} else {
				fmt.Println()
			}
			maxItr--
		}

		maxItr = 10
		fmt.Printf("B [%d, %d] sum: %d\n", start, i, sum)
		for maxItr > 0 && start < i && sum > k {
			sum -= nums[start]
			start++
			fmt.Printf("\t-Shrink [%d, %d] sum: %2d nSumK: %d", start, i, sum, nSumK)
			if sum == k {
				nSumK += 1
				fmt.Println("...Hit!!")
			} else {
				fmt.Println()
			}
			maxItr--
		}
	}
	return nSumK
}
