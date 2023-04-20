package leetcode

import "math"

type Q1007 struct{}

func (q Q1007) MinDominoRotations(tops []int, bottoms []int) int {
	nRotation, minRotation := 0, math.MaxInt
	for i := 1; i <= 6; i++ {
		nRotation = q.minDominoRotationsI(i, tops, bottoms)
		if nRotation >= 0 && minRotation > nRotation {
			minRotation = nRotation
		}

		nRotation = q.minDominoRotationsI(i, bottoms, tops)
		if nRotation >= 0 && minRotation > nRotation {
			minRotation = nRotation
		}
	}

	if minRotation == math.MaxInt {
		return -1
	}
	return minRotation
}

func (q Q1007) minDominoRotationsI(i int, tops []int, bottoms []int) int {
	nRotation := 0
	for j, n := range tops {
		if n == i {
			continue
		}

		if bottoms[j] == i {
			nRotation += 1
		} else {
			return -1
		}
	}
	return nRotation
}
