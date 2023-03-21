package leetcode

type Q0001 struct{}

func (q Q0001) TwoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	m := make(map[int]int, len(nums))
	for i, num := range nums {
		_, ok := m[num]
		if ok {
			return []int{m[num], i}
		}
		m[target-num] = i
	}
	return nil
}
