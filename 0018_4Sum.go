package leetcode

import (
	"fmt"
	"strings"
)

type Q0018 struct{}

type Q0018Data struct {
	Value int
	Index int
}

func (d Q0018Data) String() string {
	return fmt.Sprintf("%2d(%d)", d.Value, d.Index)
}

type Q0018Couple [2]Q0018Data

func (c Q0018Couple) String() string {
	return fmt.Sprintf("(%v, %v)", c[0], c[1])
}

type Q0018CoupleSet map[Q0018Couple]struct{}

func (cs Q0018CoupleSet) String() string {
	sb := strings.Builder{}
	sb.WriteString("Couple Set: {")
	i := 0
	for k := range cs {
		sb.WriteString(k.String())
		if i < len(cs)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteRune('}')
	return sb.String()
}

func (q Q0018) MergeCouple(x, y Q0018Couple) ([4]int, bool) {
	z := [4]int{}

	for _, dx := range x {
		for _, dy := range y {
			if dx.Index == dy.Index {
				return [4]int{}, false
			}
		}
	}

	xi, yi := 0, 0
	for i := 0; i < 4; i++ {
		if xi < 2 && yi < 2 {
			if x[xi].Value <= y[yi].Value {
				z[i] = x[xi].Value
				xi++
			} else {
				z[i] = y[yi].Value
				yi++
			}
		} else if xi < 2 {
			z[i] = x[xi].Value
			xi++
		} else {
			// yi < 2
			z[i] = y[yi].Value
			yi++
		}
	}
	return z, true
}

func (q Q0018) NewCouple(xi, yi, xval, yval int) Q0018Couple {
	if xval > yval {
		return Q0018Couple{
			Q0018Data{Value: yval, Index: yi},
			Q0018Data{Value: xval, Index: xi},
		}
	}
	return Q0018Couple{
		Q0018Data{Value: xval, Index: xi},
		Q0018Data{Value: yval, Index: yi},
	}
}

func (cs Q0018CoupleSet) Append(c Q0018Couple) {
	cs[c] = struct{}{}
}

func (q Q0018) FourSum(nums []int, target int) [][]int {
	counter := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if counter[nums[i]] < 4 {
			counter[nums[i]]++
		}
	}

	l := 0
	for n, c := range counter {
		for i := 0; i < c; i++ {
			nums[l] = n
			l++
		}
	}
	nums = nums[:l]

	m := map[int]Q0018CoupleSet{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sm := nums[i] + nums[j]
			if m[sm] == nil {
				m[sm] = Q0018CoupleSet{}
			}
			m[sm].Append(q.NewCouple(i, j, nums[i], nums[j]))
		}
	}

	// for k, v := range m {
	// 	fmt.Printf("%3d: %v\n", k, v)
	// }

	fs := map[[4]int]struct{}{}
	for k, s1 := range m {
		s2 := m[target-k]
		for e1 := range s1 {
			for e2 := range s2 {
				// fmt.Print("Merge", e1, e2)
				if qp, ok := q.MergeCouple(e1, e2); ok {
					// fmt.Println("->", qp)
					fs[qp] = struct{}{}
				} else {
					// fmt.Println("-> Failed")
				}
			}
		}
	}

	qps := make([][]int, len(fs))
	i := 0
	for k := range fs {
		qp := make([]int, 4)
		copy(qp, k[:])
		qps[i] = qp
		i++
	}
	return qps
}
