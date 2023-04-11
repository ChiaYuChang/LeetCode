package leetcode

type Q0746 struct{}

func (q Q0746) MinCostClimbingStairs(cost []int) int {
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	minCost := [4]int{}
	var i int
	for i = 0; i < len(cost)-1; i += 2 {
		minCost[1] = min(minCost[1], minCost[0]+cost[i])
		minCost[2] = min(minCost[0]+cost[i], minCost[1]+cost[i+1])
		minCost[3] = minCost[1] + cost[i+1]
		minCost[0], minCost[1] = minCost[2], minCost[3]
	}

	if i == len(cost)-1 {
		// len(cost) is odd
		minCost[1] = min(minCost[1], minCost[1]+cost[i])
		return minCost[1]
	}
	// len(cost) is even
	return minCost[0]
}
