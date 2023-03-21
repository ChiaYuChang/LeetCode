package leetcode

type Q0191 struct{}

func (q Q0191) HammingWeight(num uint32) int {
	ones := 0
	for num > 0 {
		if num&1 == 1 {
			ones += 1
		}
		num = num >> 1
	}
	return ones
}
