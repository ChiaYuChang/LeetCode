package leetcode

import "fmt"

type Q0946 struct{}

func ValidateStackSequences(pushed []int, popped []int) bool {
	var l int
	if l = len(pushed); l != len(popped) {
		return false
	}

	var i, p1, p2 int
	for j := 0; j < l; j++ {
		if pushed[j] == popped[0] {
			i = j
			break
		}
	}

	p1, p2, i = i-1, i+1, 1
	if p1 >= 0 && p2 < l && i < l {
		fmt.Printf("i: %d p1: %d p2: %d\n", i, p1, p2)
		if popped[i] == pushed[p2] {
			i, p2 = i+1, p2+1
		} else if popped[i] == pushed[p1] {
			i, p1 = i+1, p1-1
		} else {
			return false
		}
	}

	if p1 >= 0 {
		for ; i < l; i++ {
			fmt.Printf("i: %d p1: %d p2: %d\n", i, p1, p2)
			if popped[i] != pushed[p1] {
				return false
			}
			p1--
		}
	}

	if p2 < l {
		for ; i < l; i++ {
			fmt.Printf("i: %d p1: %d p2: %d\n", i, p1, p2)
			if popped[i] != pushed[p2] {
				return false
			}
			p2++
		}
	}
	return true
}
