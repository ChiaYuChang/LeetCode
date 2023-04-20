package leetcode

import "sort"

type Q1431 struct{}

func (q Q1431) KidsWithCandies(candies []int, extraCandies int) []bool {
	type candy [2]int

	cs := make([]candy, len(candies))
	for i := 0; i < len(candies); i++ {
		cs[i] = candy{i, candies[i]}
	}

	sort.Slice(cs, func(i, j int) bool {
		return cs[i][1] > cs[j][1]
	})

	ans := make([]bool, len(candies))
	ans[cs[0][0]] = true
	for i := 1; i < len(cs); i++ {
		if cs[0][1] <= cs[i][1]+extraCandies {
			ans[cs[i][0]] = true
		}
	}
	return ans
}
