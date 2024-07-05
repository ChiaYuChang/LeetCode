package leetcode

type Q2215 struct{}

const Q2215MinNum int = -1000
const Q2215MaxNum int = 1000

func (q Q2215) FindDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make([]bool, Q2215MaxNum-Q2215MinNum+1)
	set2 := make([]bool, Q2215MaxNum-Q2215MinNum+1)

	for _, n := range nums1 {
		set1[n-Q2215MinNum] = true
	}

	for _, n := range nums2 {
		set2[n-Q2215MinNum] = true
	}

	mutuallyExclusiveSet := [][]int{
		make([]int, 0, len(nums1)),
		make([]int, 0, len(nums2)),
	}

	for i := Q2215MinNum; i <= Q2215MaxNum; i++ {
		if set1[i-Q2215MinNum] != set2[i-Q2215MinNum] {
			if set1[i-Q2215MinNum] {
				mutuallyExclusiveSet[0] = append(mutuallyExclusiveSet[0], i)
			} else {
				mutuallyExclusiveSet[1] = append(mutuallyExclusiveSet[1], i)
			}
		}
	}
	return mutuallyExclusiveSet
}
