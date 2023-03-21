package leetcode

import "fmt"

type Q0134 struct{}

func (q Q0134) CanCompleteCircuit(gas []int, cost []int) int {
	var net, cur, str, dif int

	// find the start index of the last positive interval
	for i, g := range gas {
		dif = g - cost[i]
		net += dif
		cur += dif
		if cur < 0 {
			cur = 0
			str = i + 1
		}
	}

	if net < 0 {
		return -1
	} else {
		return str
	}
}

type Q0134Interval struct {
	start, end int // end is unnecessary
	sum        int
}

func (i Q0134Interval) String() string {
	return fmt.Sprintf("Interval {[%d, %d) %d}", i.start, i.end, i.sum)
}

func (q Q0134) SlowCanCompleteCircuit(gas []int, cost []int) int {
	j := 0
	intervals := []*Q0134Interval{{start: 0, sum: 0}}
	total := 0
	for i := 0; i < len(gas); i++ {
		net := gas[i] - cost[i]
		total += net
		if intervals[j].sum >= 0 {
			// positive intervals
			intervals[j].sum += net
		} else if intervals[j].sum < 0 && net < 0 {
			// negative intervals
			intervals[j].sum += net
		} else {
			intervals[j].end = i
			intervals = append(intervals, &Q0134Interval{start: i, end: i, sum: net})
			j++
		}
	}
	intervals[j].end = len(gas)

	if total < 0 {
		return -1
	} else {
		if intervals[j].sum > 0 {
			return intervals[j].start
		} else {
			return intervals[0].start
		}
	}
}
