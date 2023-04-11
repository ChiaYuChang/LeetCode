package leetcode

type Q0495 struct{}

func (q Q0495) FindPoisonedDuration(timeSeries []int, duration int) int {
	pst, strAt, endAt := 0, timeSeries[0], timeSeries[0]+duration

	for i := 1; i < len(timeSeries); i++ {
		curr := timeSeries[i]
		if curr > endAt {
			pst, strAt = pst+endAt-strAt, curr
		}
		endAt = curr + duration
	}
	pst += endAt - strAt
	return pst
}
