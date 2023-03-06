package leetcode

func DestCity(paths [][]string) string {
	// Name to Index
	n2i := map[string]int{}

	for _, e := range paths {
		n2i[e[0]] += 1 // out edge
		n2i[e[1]] -= 1 // in edge
	}

	var k string
	var v int
	for k, v = range n2i {
		if v < 0 {
			break
		}
	}
	return k
}
