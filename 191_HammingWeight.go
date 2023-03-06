package leetcode

func HammingWeight(num uint32) int {
	ones := 0
	for num > 0 {
		if num&1 == 1 {
			ones += 1
		}
		num = num >> 1
	}
	return ones
}
