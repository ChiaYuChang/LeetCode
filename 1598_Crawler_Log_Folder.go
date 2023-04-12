package leetcode

type Q1598 struct{}

func (q Q1598) MinOperations(logs []string) int {
	if len(logs) == 0 {
		return 0
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	depth := 0
	for _, cmd := range logs {
		cmd = cmd[:len(cmd)-1]
		if cmd == ".." {
			depth = max(0, depth-1)
		} else if cmd == "." {
			continue
		} else {
			depth++
		}
	}
	return depth
}
