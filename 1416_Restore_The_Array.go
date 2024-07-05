package leetcode

import (
	"fmt"
)

type Q1416 struct{}

const Q1416_MOD_NUM = 1_000_000_007

func (q Q1416) NumberOfArrays(s string, k int) int {
	if s[0] == '0' {
		return 0
	}

	ln := len(s)
	// ss := make([]string, 0, ln)
	nums := make([]int, 0, ln)
	dgts := make([]int, 0, ln)
	for i := ln - 1; i >= 0; i-- {
		if s[i] == '0' {
			j := i - 1
			d := 10
			for ; j >= 0; j-- {
				if s[j] != '0' {
					break
				}
				d *= 10
				if d > k {
					return 0
				}
			}
			num := int(s[j]-'0') * d
			if num > k {
				return 0
			}
			// ss = append(ss, s[j:i+1])
			nums = append(nums, num)
			dgts = append(dgts, d)
			i = j
		} else {
			// ss = append(ss, s[i:i+1])
			nums = append(nums, int(s[i]-'0'))
			dgts = append(dgts, 1)
		}
	}

	// fmt.Println(ss)
	fmt.Println("nums:", nums)
	fmt.Println("dgts:", dgts)

	// dp := make([]int, len(ss)+1)
	dp := make([]int, len(nums)+1)
	dp[0] = 1
	for i := 1; i <= len(nums); i++ {
		// ns := ""
		fmt.Println("i:", i)
		ns := 0
		for j := i - 1; j >= 0; j-- {
			ns += nums[j]
			// ns += ss[j]
			// n, _ := strconv.Atoi(ns)
			fmt.Println("\t ns:", ns)
			if ns > k {
				fmt.Println("\t break")
				break
			}
			ns *= 10
			if j-1 >= 0 {
				ns *= dgts[j-1]
			}
			dp[i] += dp[j]
			dp[i] %= Q1416_MOD_NUM
		}
		if dp[i] == 0 {
			return 0
		}
	}
	fmt.Println(dp)
	return dp[len(nums)]
	// return dp[len(ss)]
}
