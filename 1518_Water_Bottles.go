package leetcode

type Q1518 struct{}

func (q Q1518) NumWaterBottles(numBottles int, numExchange int) int {
	total, empty := numBottles, numBottles
	for empty >= numExchange {
		q := empty / numExchange
		r := empty % numExchange
		total += q
		empty = q + r
	}
	return total
}
