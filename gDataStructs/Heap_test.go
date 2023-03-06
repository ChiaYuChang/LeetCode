package gdatastructs_test

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	gds "gitlab.com/gjerry134679/leetcode/gDataStructs"
	"golang.org/x/exp/constraints"
)

type KeyValPair[T constraints.Ordered] struct {
	Keys gds.Ordered[T]
	Vals string
}

func (kp *KeyValPair[T]) DeepCopy() *KeyValPair[T] {
	return &KeyValPair[T]{
		Keys: gds.NewOrdered(kp.Keys.Val()),
		Vals: kp.Vals,
	}
}

type SortBy[T constraints.Ordered] []*KeyValPair[T]

func (a SortBy[T]) Len() int           { return len(a) }
func (a SortBy[T]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy[T]) Less(i, j int) bool { return a[i].Keys.Less(a[j].Keys) }

func TestHeap(t *testing.T) {
	type testCase struct {
		name  string
		n     int
		pairs []*KeyValPair[int]
	}

	tcs := []testCase{{n: 2}, {n: 5}}
	for i := range tcs {
		p := make([]*KeyValPair[int], tcs[i].n)
		for i := range p {
			s := rand.Int()
			p[i] = &KeyValPair[int]{
				Keys: gds.NewOrdered(s),
				Vals: strconv.Itoa(s),
			}
		}
		tcs[i].pairs = p
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(
			tc.name,
			func(t *testing.T) {
				hp := gds.NewHeap[gds.Ordered[int], string]()
				require.NotNil(t, hp)

				var i int
				i = 0
				for _, p := range tc.pairs {
					np := p.DeepCopy()
					hp.Insert(np.Keys, np.Vals)
					i++
					require.Equal(t, i, hp.Size())
				}

				sort.Sort(SortBy[int](tc.pairs))
				t.Log("list")
				for _, p := range tc.pairs {
					t.Log(p.Vals)
					// _, v := hp.Pop()
					// require.Equal(t, v, p.Vals)
				}

				t.Log("heap")
				i = 0
				for hp.Size() > 0 && i < tc.n*2 {
					_, v := hp.Pop()
					t.Log(v)
					i++
				}
			},
		)
	}

	// t.Log(hp.String())
	// _, v := hp.Top()
	// require.Equal(t, "two", v)

	// hp.Insert(gds.NewOrdered(1), "one")
	// _, v = hp.Top()
	// t.Log(hp.String())
	// require.Equal(t, "two", v)

	// hp.Insert(gds.NewOrdered(3), "three (1)")
	// t.Log(hp.String())
	// _, v = hp.Top()
	// require.Equal(t, "three (1)", v)

	// hp.Insert(gds.NewOrdered(4), "four")
	// t.Log(hp.String())
	// _, v = hp.Top()
	// require.Equal(t, "four", v)

	// hp.Insert(gds.NewOrdered(3), "three (2)")
	// t.Log(hp.String())
	// _, v = hp.Top()
	// require.Equal(t, "four", v)

	// _, v = hp.Pop()
	// t.Log(hp.String())
	// require.Equal(t, "four", v)
}
