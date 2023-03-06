package gdatastructs_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	gds "gitlab.com/gjerry134679/leetcode/gDataStructs"
)

func TestReverseLinkedList(t *testing.T) {
	var genRdnInts func(int) []int = func(n int) []int {
		ints := make([]int, n)
		for i := range ints {
			ints[i] = rand.Int()
		}
		return ints
	}

	tcs := []int{1, 10, 100, 500, 1000}

	for i := range tcs {
		n := tcs[i]
		t.Run(
			fmt.Sprintf("Case_%d", i+1),
			func(t *testing.T) {
				t.Parallel()
				var i int
				a := genRdnInts(n)
				ll := gds.NewLinkListFromArray(a, false)

				c1 := make(chan int)
				ll.Walk(c1)
				i = 0
				for n := range c1 {
					require.Equal(t, a[i], n)
					i++
				}

				ll.Reverse()
				c2 := make(chan int)
				ll.Walk(c2)

				i = len(a) - 1
				for n := range c2 {
					require.Equal(t, a[i], n)
					i--
				}
			},
		)
	}
}
