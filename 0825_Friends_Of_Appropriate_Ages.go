package leetcode

type Q0823 struct{}

const Q0823MaxAge = 120

func (q Q0823) NumFriendRequests(ages []int) (nreq int) {
	cnt := make([]int, Q0823MaxAge+1) // counter
	cum := make([]int, Q0823MaxAge+1) // accumuated sum
	for _, a := range ages {
		cnt[a] += 1
	}

	for i := 1; i < len(cnt); i++ {
		cum[i] += cum[i-1] + cnt[i]
	}

	var lb, ub, n int
	for x := 1; x < Q0823MaxAge+1; x++ {
		// (lb, ub]
		if x%2 == 0 {
			// x is an even number
			ub = x
			lb = x/2 + 7
		} else {
			// x is an odd number
			n = (x - 1) / 2
			ub = x
			lb = n + 7
		}

		if x < 100 {
			ub = q.min(99, ub)
		}

		if cnt[x] > 0 && ub > lb {
			nreq += cnt[x] * (cum[ub] - cum[lb])
			if ub >= x && x > lb {
				// remove self loop
				nreq -= cnt[x]
			}
		}
	}

	return nreq
}

func (q Q0823) min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
