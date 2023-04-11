package leetcode

type Q0712 struct{}

func MinimumDeleteSum(s1 string, s2 string) int {
	l1 := len(s1) + 1
	l2 := len(s2) + 1

	if l2 > l1 {
		l1, l2, s1, s2 = l2, l1, s2, s1
	}

	var curr, prev []int

	curr = make([]int, l2)
	for i := 1; i < l2; i++ {
		curr[i] = curr[i-1] - int(s2[i-1])
	}

	for i := 1; i < l1; i++ {
		curr, prev = make([]int, l2), curr
		curr[0] = prev[0] - int(s1[i-1])
		for j := 1; j < l2; j++ {
			if s1[i-1] == s2[j-1] {
				curr[j] = Max(prev[j]-int(s1[i-1]), curr[j-1]-int(s2[j-1]), prev[j-1])
			} else {
				curr[j] = Max(prev[j]-int(s1[i-1]), curr[j-1]-int(s2[j-1]))
			}
		}
	}

	return -curr[l2-1]
}

func Max(nums ...int) int {
	max := nums[0]
	for _, n := range nums {
		if max < n {
			max = n
		}
	}
	return max
}
