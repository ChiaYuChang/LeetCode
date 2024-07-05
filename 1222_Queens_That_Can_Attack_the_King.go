package leetcode

type Q1222 struct{}

type Q1222NearestQueen map[[2]int]*Q1222WhichQueen

type Q1222WhichQueen struct {
	index    int
	distance int
}

func (nq Q1222NearestQueen) Append(which [2]int, index int, distance int) {
	if q, ok := nq[which]; ok {
		if q.distance > distance {
			nq[which].distance = distance
			nq[which].index = index
		}
	} else {
		nq[which] = &Q1222WhichQueen{index, distance}
	}
}

func (q Q1222) QueensAttacktheKing(queens [][]int, king []int) [][]int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	nearestQueen := Q1222NearestQueen{}
	for i, queen := range queens {
		x := queen[0] - king[0]
		y := queen[1] - king[1]
		if x == 0 {
			nearestQueen.Append([2]int{0, y / abs(y)}, i, abs(y))
		}
		if y == 0 {
			nearestQueen.Append([2]int{x / abs(x), 0}, i, abs(x))
		}
		if x == y || x == -y {
			nearestQueen.Append([2]int{x / abs(x), y / abs(y)}, i, min(abs(x), abs(y)))
		}
	}

	attackers := make([][]int, 0, len(nearestQueen))
	for _, qd := range nearestQueen {
		attackers = append(attackers, queens[qd.index])
	}
	return attackers
}

func (q Q1222) QueensAttacktheKingFirstHit(queens [][]int, king []int) [][]int {
	directions := [][]int{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	targets := map[int]int{}
	for i, position := range queens {
		targets[position[0]*8+position[1]] = i
	}

	attackers := [][]int{}
	for _, direction := range directions {
		if i := q.firstHit(king, direction, targets); i >= 0 {
			attackers = append(attackers, queens[i])
		}
	}
	return attackers
}

func (q Q1222) firstHit(ori, direction []int, targets map[int]int) int {
	x, y := ori[0]+direction[0], ori[1]+direction[1]
	for x >= 0 && x < 8 && y >= 0 && y < 8 {
		if i, ok := targets[x*8+y]; ok {
			return i
		}
		x += direction[0]
		y += direction[1]
	}
	return -1
}
