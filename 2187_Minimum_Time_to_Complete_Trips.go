package leetcode

type Q2187 struct{}

func (q Q2187) MinimumTime(trips []int, totalTrips int) int64 {
	if totalTrips == 0 {
		return 0
	}

	// compute upper bounds
	lhs, mid, rhs := 0, 0, q.MaxTime(trips)*totalTrips
	nTrips := 0

	// binary search
	for rhs > lhs {
		mid = (lhs + rhs) / 2
		nTrips = q.CountTrips(trips, mid)
		if nTrips < totalTrips {
			// left boudn should always less than total trips
			lhs = mid + 1
		} else {
			// right bound greater or equal to total trips
			rhs = mid
		}
	}
	return int64(rhs)
}

func (q Q2187) CountTrips(trips []int, time int) (nTrips int) {
	for _, tp := range trips {
		nTrips += time / tp
	}
	return
}

func (q Q2187) MaxTime(nums []int) int {
	m := -1
	for _, n := range nums {
		if m < n {
			m = n
		}
	}
	return m
}
