package leetcode

// type Q894TreeNode struct {
// 	Val   int
// 	Left  *Q894TreeNode
// 	Right *Q894TreeNode
// }

// func allPossibleFBT(n int) []*Q894TreeNode {
// 	// nums := make([]int, n)
// 	// for i := nums {
// 	// 	nums[i] = i
// 	// }

// 	// for i := 1; i < n; i += 2 {
// 	// 	all

// 	// }
// 	return nil
// }

// type FBT struct {
// 	cache map[int][][]int
// }

// func (f *FBT) NewFBT(n int) *FBT {
// 	cache := map[int][][]int{}
// 	cache[1] = append(cache[1], []int{0})
// 	cache[3] = append(cache[3], []int{1, 0, 2})
// 	return &FBT{cache: cache}
// }

// func (f *FBT) FindFBT(nums *[]int) []int {
// 	if v, ok := f.cache[len(*nums)]; ok {
// 		n := make([]int, len(*nums))
// 		for i := range v {
// 			n[i] = nums[v[i]]
// 		}
// 		return n
// 	}

// 	for i := 1; i < len(nums); i += 2 {
// 		lhs := f.FindFBT(nums[:i])
// 		mid := nums[i]
// 		rhs := f.FindFBT(nums[(i + 1):])
// 		lhs = append(lhs, mid)
// 		lhs = append(lhs, rhs...)
// 	}
// }
