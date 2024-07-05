package leetcode

// import (
// 	"fmt"
// )

// type Q1977 struct{}

// const MOD_NUM = 1_000_000_007

// func NumberOfCombinations(num string) int {
// 	if num[0] == '0' {
// 		return 0
// 	}

// 	ln := len(num)
// 	nlst := make([]int, 0)
// 	dgts := make([]int, 0)
// 	for i := ln - 1; i >= 0; i-- {
// 		if num[i] == '0' {
// 			j := i - 1
// 			d := 10
// 			for ; j >= 0; j-- {
// 				if num[j] != '0' {
// 					break
// 				}
// 				d *= 10
// 			}
// 			n := int(num[j]-'0') * d
// 			nlst = append(nlst, n)
// 			dgts = append(dgts, d)
// 			i = j
// 		} else {
// 			nlst = append(nlst, int(num[i]-'0'))
// 			dgts = append(dgts, 1)
// 		}
// 	}

// 	ln = len(nlst)
// 	nlst = append(nlst, -1)
// 	dp := make([]int, ln+1)
// 	dp[0] = 1
// 	for i := 1; i <= ln; i++ {
// 		ns := 0
// 		fmt.Println("i:", i)
// 		for j := i - 1; j > 0; j-- {
// 			ns += nlst[j]
// 			if nlst[i] > ns {
// 				fmt.Printf("\t %d  > %d...break\n", nlst[i], ns)
// 				break
// 			}
// 			fmt.Printf("\t %d <= %d...OK\n", nlst[i], ns)

// 			dp[i] += dp[j-1]
// 			dp[i] %= MOD_NUM
// 			ns *= 10
// 			if j-1 >= 0 {
// 				ns *= dgts[j-1]
// 			}
// 			fmt.Println("\t -dp:", dp)
// 		}
// 		fmt.Println("-> dp:", dp)
// 	}
// 	fmt.Println("-> dp:", dp)
// 	return dp[ln]
// }
