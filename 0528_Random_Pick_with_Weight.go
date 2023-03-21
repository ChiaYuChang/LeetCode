package leetcode

// import (
// 	"math/rand"
// 	"time"
// )

// type Solution struct {
// 	cw  *[]int
// 	ln  int
// 	rng *rand.Rand
// }

// func Constructor(ws []int) Solution {
// 	for i := 1; i < len(ws); i++ {
// 		ws[i] += ws[i-1]
// 	}

// 	return Solution{
// 		cw:  &ws,
// 		ln:  len(ws),
// 		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
// 	}
// }

// func (this *Solution) PickIndex() int {
// 	var i, w int

// 	j := this.rng.Intn(this.ln)
// 	for i, w = range *this.cw {
// 		if w > j {
// 			break
// 		}
// 	}
// 	return i
// }
