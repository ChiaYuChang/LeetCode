package leetcode

import "fmt"

const nHours, nMinutes = 4, 6

func ReadBinaryWatch(turnedOn int) []string {
	a := []string{}
	p := Permutator{}

	var turnedOnHours, turnedOnMinutes int
	for turnedOnHours = 0; turnedOnHours <= turnedOn; turnedOnHours++ {
		turnedOnMinutes = turnedOn - turnedOnHours
		hs := p.Permute(turnedOnHours, nHours-turnedOnHours)
		ms := p.Permute(turnedOnMinutes, nMinutes-turnedOnMinutes)
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

type Permutator struct{}

func (p Permutator) Permute(nTrue, nFalse int) []int {
	c := []int{}
	p.permute(0, nTrue, nFalse, &c)
	return c
}

func (p Permutator) permute(prefix, nTrue, nFalse int, c *[]int) {
	if nTrue+nFalse == 0 {
		(*c) = append((*c), prefix)
		return
	}

	if nTrue > 0 {
		p.permute((prefix<<1)+1, nTrue-1, nFalse, c)
	}

	if nFalse > 0 {
		p.permute((prefix<<1)+0, nTrue, nFalse-1, c)
	}
	return
}
