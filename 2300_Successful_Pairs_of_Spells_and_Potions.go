package leetcode

import "sort"

type Q2300 struct{}

func (q Q2300) SuccessfulPairs(spells []int, potions []int, success int64) []int {
	for i, n := range potions {
		potions[i] = q.DividAndCeil(int(success), n)
	}
	sort.Sort(sort.IntSlice(potions))
	n := make([]int, len(spells))
	for i, spell := range spells {
		n[i] = q.BinarySearch(spell, potions)
	}
	return n
}

// divid and ceil
func (q Q2300) DividAndCeil(dividend, divisor int) int {
	quotient := dividend / divisor
	if dividend%divisor > 0 {
		quotient++
	}
	return quotient
}

// search for the right most element that less or equal to the target
func (q Q2300) BinarySearch(target int, nums []int) int {
	lhs, rhs := 0, len(nums)
	for lhs < rhs {
		mid := (lhs + rhs) / 2
		if nums[mid] > target {
			rhs = mid
		} else {
			lhs = mid + 1
		}
	}
	// should return rhs - 1 but 1 will be add back at SuccessfulPairs func
	return rhs
}
