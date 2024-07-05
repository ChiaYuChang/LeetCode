package leetcode

import (
	"fmt"
	"sort"
)

type Q1712 struct{}

const Q1712_MOD_NUM = 1_000_000_007

func WaysToSplit(num []int) int {
	if len(num) < 0 {
		panic("length of num should be greater than or equal to 3")
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	for i := 1; i < len(num); i++ {
		num[i] = num[i-1] + num[i]
	}

	end := len(num) - 1
	nWay := 0
	if num[end] == 0 {
		return (len(num) - 1) * (len(num) - 2) / 2
	}

	var lMin, lMax, rMin, rMax int
	for lMin = 0; lMin < len(num)-2; lMin++ {
		rMin = lMin + 1 + sort.SearchInts(num[lMin+1:], num[lMin]*2)
		lMax = min(sort.SearchInts(num, num[lMin]+1)-1, rMin-1) // next none zero element
		if rMin < len(num)-1 {
			fmt.Printf(
				"\t find: %d ((%d + %d)/2 + 1) in %v\n",
				(num[end]+num[lMin])/2+1, num[end], num[lMin], num[rMin:end])
			rMax = rMin + sort.SearchInts(num[rMin:end], (num[end]+num[lMin])/2+1)
			// fmt.Printf("\t- Left [%d, %d] Right: [%d, %d) (%d)\n", lMin, lMax, rMin, rMax, (rMax - rMin))
			// for right := rMin; right < rMax; right++ {
			// 	fmt.Printf("\t-     num: %v %v %v\n", cpNum[:lMin+1], cpNum[lMin+1:right+1], cpNum[right+1:])
			// 	fmt.Printf("\t- cum sum: %v %v %v\n", num[:lMin+1], num[lMin+1:right+1], num[right+1:])
			// }
			nWay += (lMax - lMin + 1) * (rMax - rMin)
		}
		lMin = lMax
	}
	return nWay
}
