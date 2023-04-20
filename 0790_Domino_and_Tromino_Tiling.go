package leetcode

// const MOD = 1_000_000_007

// type Q0790 struct{}

// type triplet [3]int

// func (tri triplet) String() string {
// 	return fmt.Sprintf("(%2d, %2d, %2d)", tri[0], tri[1], tri[2])
// }

// func (tri triplet) Sum() int {
// 	return tri[0] + tri[1] + tri[2]
// }

// func NumTilings(n int) int {
// 	// n = a + 2b + 3c

// 	t := triplet{}
// 	n, t[2] = n%3, n/3
// 	t[0], t[1] = n%2, n/2

// 	triplets := []triplet{t}
// 	for j := 1; j <= triplets[0][2]; j++ {
// 		triplets = append(triplets, triplet{triplets[0][0] + j, triplets[0][1] + j, triplets[0][2] - j})
// 	}

// 	l := len(triplets)
// 	for i := 1; i < l; i++ {
// 		for j := 1; j <= triplets[i][1]; j++ {
// 			triplets = append(triplets, triplet{triplets[i][0] + 2*j, triplets[i][1] - j, triplets[i][2]})
// 		}
// 	}

// 	total := 0
// 	for _, tri := range triplets {
// 		n1 := int(math.Pow(float64(t[0]), float64(tri[2]+1))) % MOD
// 		n2 := int(math.Pow(float64(t[1]), float64(2*tri[2]+1))) % MOD
// 		if n1 == 0 {
// 			n1 = 1
// 		}

// 		if n2 == 0 {
// 			n2 = 1
// 		}
// 		fmt.Println(tri, n1, n2)
// 		total += (n1 * n2 * 1 << tri[2]) % MOD
// 	}

// 	return total
// }

// func Compose(n, m int) int {
// 	if m > n-m {
// 		m = n - m
// 	}

// 	c := 1
// 	for i := m; i > 0; i-- {
// 		c *= n
// 		c /= i
// 		n--
// 	}
// 	return c
// }
