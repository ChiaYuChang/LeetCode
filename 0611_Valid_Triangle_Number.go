package leetcode

import (
	"sort"
)

type Q0611 struct{}

func (q Q0611) TriangleNumber(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	l := len(nums)
	n := 0

	for i := l - 1; i >= 2; i-- {
		lhs, rhs := 0, i-1
		for rhs > lhs {
			if nums[i] < nums[lhs]+nums[rhs] {
				n += rhs - lhs
				rhs--
			} else {
				lhs++
			}
		}
	}
	return n
}

func (q Q0611) SlowTriangleNumber(nums []int) int {
	count := map[int]int{}
	for _, n := range nums {
		if n > 0 {
			count[n] += 1
		}
	}

	if len(count) < 1 {
		return 0
	}

	keys := make([]int, 0, len(count))
	for k := range count {
		keys = append(keys, k)
	}
	sort.Sort(sort.IntSlice(keys))

	n := 0
	for i := 0; i < len(keys); i++ {
		for j := i; j < len(keys); j++ {
			// [min, max]
			min := keys[j] - keys[i] + 1
			max := keys[j] + keys[i] - 1
			lb, ub, ok := q.ClosedIntervalSearch(keys, min, max)
			if ok {
				for k := lb; k <= ub; k++ {
					if k < j {
						continue
					}
					if i == j && i == k {
						n += q.Choose(count[keys[i]], 3)
					} else if i == j {
						n += q.Choose(count[keys[i]], 2) * count[keys[k]]
					} else if i == k {
						n += q.Choose(count[keys[i]], 2) * count[keys[j]]
					} else if j == k {
						n += q.Choose(count[keys[j]], 2) * count[keys[i]]
					} else {
						n += count[keys[i]] * count[keys[j]] * count[keys[k]]
					}
				}
			}
		}
	}
	return n
}

func (q Q0611) ClosedIntervalSearch(nums []int, min, max int) (from, to int, ok bool) {
	if max < nums[0] || min > nums[len(nums)-1] {
		return 0, 0, false
	}
	return q.SearchLeftMost(nums, min), q.SearchRightMost(nums, max), true
}

func (q Q0611) SearchLeftMost(nums []int, target int) int {
	lhs, rhs, mid := 0, len(nums)-1, 0
	for lhs != rhs {
		mid = (lhs + rhs) / 2
		if nums[mid] < target {
			lhs = mid + 1
		} else {
			rhs = mid
		}
	}

	if nums[lhs] >= target {
		return lhs
	}

	return -1
}

func (q Q0611) SearchRightMost(nums []int, target int) int {
	lhs, rhs, mid := 0, len(nums)-1, 0
	for lhs != rhs {
		mid = (lhs + rhs + 1) / 2
		// fmt.Printf("LHS: %d, RHS: %d, MID: %d (%d)\n", lhs, rhs, mid, nums[mid])
		if nums[mid] > target {
			rhs = mid - 1
		} else {
			lhs = mid
		}
	}

	if nums[rhs] <= target {
		return rhs
	}

	return -1
}

func (q Q0611) Choose(n, m int) int {
	if n < m {
		return 0
	}

	if n-m < m {
		m = n - m
	}

	nmrtr, dnmrtr := 1, 1
	for i, j := 1, n; i <= m; {
		dnmrtr *= i
		nmrtr *= j
		i++
		j--
	}
	return nmrtr / dnmrtr
}
