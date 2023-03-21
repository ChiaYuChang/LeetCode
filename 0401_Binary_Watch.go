package leetcode

import "fmt"

type Q0401 struct{}

const Q0401nHours, Q0401nMinutes = 4, 6

func (q Q0401) ReadBinaryWatch(turnedOn int) []string {
	a := []string{}
	p := Q0401{}

	var turnedOnHours, turnedOnMinutes int
	for turnedOnHours = 0; turnedOnHours <= turnedOn; turnedOnHours++ {
		turnedOnMinutes = turnedOn - turnedOnHours
		hs := p.Permute(turnedOnHours, Q0401nHours-turnedOnHours)
		ms := p.Permute(turnedOnMinutes, Q0401nMinutes-turnedOnMinutes)
		for _, h := range hs {
			if h < 12 {
				for _, m := range ms {
					if m < 60 {
						a = append(a, fmt.Sprintf("%d:%02d", h, m))
					}
				}
			}
		}
	}
	return a
}

func (q Q0401) Permute(nTrue, nFalse int) []int {
	c := []int{}
	q.permute(0, nTrue, nFalse, &c)
	return c
}

func (q Q0401) permute(prefix, nTrue, nFalse int, c *[]int) {
	if nTrue+nFalse == 0 {
		(*c) = append((*c), prefix)
		return
	}

	if nTrue > 0 {
		q.permute((prefix<<1)+1, nTrue-1, nFalse, c)
	}

	if nFalse > 0 {
		q.permute((prefix<<1)+0, nTrue, nFalse-1, c)
	}
	return
}
