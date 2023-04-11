package leetcode

import "sort"

type Q2410 struct{}

func (q Q2410) MatchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Sort(sort.IntSlice(players))
	sort.Sort(sort.IntSlice(trainers))

	nPair := 0
	for pp, pt := 0, 0; pp < len(players) && pt < len(trainers); pt++ {
		if players[pp] <= trainers[pt] {
			pp++
			nPair++
		}
	}
	return nPair
}
