package leetcode

import (
	"fmt"
	"strconv"
)

type NumberNode struct {
	N       int
	IsOne   bool // is Nums == []int{1} ?
	IsMax   bool
	PreOnes int
}

func (n NumberNode) String() string {
	return fmt.Sprintf(
		"NumberNode{IsOne: %v. IsMax: %v, PreOnes %d, N%3d}",
		n.IsOne, n.IsMax, n.PreOnes, n.N)
}

func NewNumberNode(n int, isMax, isOne bool, preOnes int) *NumberNode {
	if isOne {
		preOnes += 1
	}
	newnode := &NumberNode{
		N:       n,
		IsOne:   isOne,
		IsMax:   isMax,
		PreOnes: preOnes,
	}
	fmt.Println(newnode)
	return newnode
}

func NumberOfDigitOne(num int) int {
	if num < 1 {
		return 0
	}

	b2i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := strconv.Itoa(num)

	var qo, qe []*NumberNode
	f := true
	qe = append(qe, NewNumberNode(1, true, false, 0))

	for i := 0; i < len(s); i++ {
		d := b2i[s[i]-'0']
		fmt.Println("d:", d)
		if f {
			qo = []*NumberNode{}
			AppendNumberNode(&qe, &qo, d)
			qe = nil
		} else {
			qe = []*NumberNode{}
			AppendNumberNode(&qo, &qe, d)
			qo = nil
		}
		f = !f
	}

	num1 := 0
	if !f {
		for _, n := range qo {
			num1 += n.PreOnes * n.N
		}
	} else {
		for _, n := range qe {
			num1 += n.PreOnes * n.N
		}
	}
	return num1
}

func AppendNumberNode(src, dst *[]*NumberNode, digit int) {
	for _, n := range *src {
		fmt.Println("---------------------------------------------------------------")
		if n.IsMax {
			// next digit should be less than digit
			switch digit {
			case 0:
				(*dst) = append((*dst), NewNumberNode(n.N, true, false, n.PreOnes)) // 0
			case 1:
				(*dst) = append((*dst), NewNumberNode(n.N, true, true, n.PreOnes))   // 1
				(*dst) = append((*dst), NewNumberNode(n.N, false, false, n.PreOnes)) // 0
			default: // 2...9
				(*dst) = append((*dst), NewNumberNode(n.N, true, false, n.PreOnes))            // d
				(*dst) = append((*dst), NewNumberNode(n.N, false, true, n.PreOnes))            // 1
				(*dst) = append((*dst), NewNumberNode(n.N*(digit-1), false, false, n.PreOnes)) // id d == 2, 0; eles 0, 2...(d-1)
			}
		} else {
			// next digit w/o constrain
			(*dst) = append((*dst), NewNumberNode(n.N, false, true, n.PreOnes))    // 1
			(*dst) = append((*dst), NewNumberNode(n.N*9, false, false, n.PreOnes)) // 0, 2...9
		}
	}
	fmt.Println("---------------------------------------------------------------")
}

type pascal struct {
	rows        [][]int
	powerOfNine []int
	nrow        int
}

func (p *pascal) Next() (*pascal, []int) {
	l := p.rows[p.nrow]
	nl := make([]int, len(l)+2, len(l)+2)
	nl[0] = 1
	nl[len(l)-1] = 1
	for i := 1; i < len(l); i++ {
		nl[i] = l[i-1] + l[i]
	}
	p.rows = append(p.rows, nl)
	p.powerOfNine = append(p.powerOfNine, p.powerOfNine[p.nrow]*9)
	p.nrow++
	return p, nl
}

func (p *pascal) DotPowerOfNine(nrow int) int {
	sum := 0
	coefficients := p.GetN(nrow)
	for i := 0; i < nrow; i++ {
		sum += p.powerOfNine[i] * coefficients[i]
	}
	return sum
}

func (p *pascal) GetN(n int) []int {
	for n > p.nrow {
		p.Next()
	}
	return p.rows[n]
}

func newPascal() *pascal {
	return &pascal{
		rows:        [][]int{{1}},
		powerOfNine: []int{1},
		nrow:        0,
	}
}
