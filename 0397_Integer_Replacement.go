package leetcode

type Q0397 struct{}

func (q Q0397) IntegerReplacement(n int) int {
	c := q.NewCache()
	return c.MinStepToOne(n)
}

type Q0397Cache map[int]int

func (q Q0397) NewCache() Q0397Cache {
	return Q0397Cache(map[int]int{1: 0, 2: 1})
}

func (c Q0397Cache) MinStepToOne(n int) int {
	if minStep, ok := c[n]; ok {
		return minStep
	}

	minStep := 1
	if n%2 == 0 {
		minStep += c.MinStepToOne(n / 2)
	} else {
		pOne := c.MinStepToOne(n + 1)
		mOne := c.MinStepToOne(n - 1)
		if pOne > mOne {
			minStep += mOne
		} else {
			minStep += pOne
		}
	}
	c[n] = minStep
	return minStep
}
