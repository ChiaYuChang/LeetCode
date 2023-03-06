package leetcode

import (
	"fmt"
)

type sSum struct {
	candidates []int
	max        int
	remain     int
}

func (s sSum) copy() sSum {
	c := make([]int, len(s.candidates), len(s.candidates)+1)
	copy(c, s.candidates)
	return sSum{c, s.max, s.remain}
}

func CombinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	cs := make([][]int, 0)

	// s := gdt.NewStackQueue[sSum]()
	// s.Enqueue(sSum{candidates: []int{}, max: 0, remain: target})
	s := make([]sSum, 1, 20)
	s[0] = sSum{candidates: []int{}, max: 0, remain: target}

	for i := 0; i < len(s); i++ {
		// sS, _ := s.Dequeue()
		sS := s[i]

		for _, c := range candidates {
			if c == 0 {
				return nil
			}
			if c >= sS.max {
				r := sS.remain - c
				if r > 0 {
					newsS := sS.copy()
					newsS.candidates = append(newsS.candidates, c)
					newsS.remain = sS.remain - c
					newsS.max = c
					// fmt.Printf("q (%d) candidates: %v\n", s.Length, newsS.candidates)
					// s.Enqueue(newsS)
					s = append(s, newsS)
				}
				if r == 0 {
					newsS := sS.copy()
					newsS.candidates = append(newsS.candidates, c)
					fmt.Printf(">> %v\n", newsS.candidates)
					cs = append(cs, newsS.candidates)
				}
			}
		}
	}
	return cs
}
