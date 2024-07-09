package leetcode

type Q1701 struct{}

func (q Q1701) AverageWaitingTime(customers [][]int) float64 {
	prcssEndAt := 0
	avgWT := 0.0
	nCstmr := float64(len(customers))
	for _, cstmr := range customers {
		if prcssEndAt < cstmr[0] {
			prcssEndAt = cstmr[0] + cstmr[1]
			avgWT += float64(cstmr[1]) / nCstmr
		} else {
			prcssEndAt += cstmr[1]
			avgWT += float64(prcssEndAt-cstmr[0]) / nCstmr
		}
	}
	return avgWT
}
