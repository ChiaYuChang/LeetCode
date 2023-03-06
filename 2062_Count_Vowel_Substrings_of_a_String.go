package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

type Q2062 struct{}

type Q2062Char struct {
	char     byte
	position []int
}

func (q Q2062) CountVowelSubstrings(word string) int {
	set := NewQ2062Set(true)

	i := 0
	for _, sub := range q.ExtractAEIOUSubStrings(word, 5) {
		aeiou := NewAEIOU()
		for i := range sub {
			aeiou.Append(i, sub[i])
		}
		aeiou.Sort()

		for true {
			if s, e, ok := aeiou.NextAEIOUInterval(); ok {
				for j := e; j < len(sub); j++ {
					fmt.Printf("%3d: %s\n", i, sub[s:j+1])
					set.Append(sub[s : j+1])
					i++
				}
			} else {
				break
			}
		}
	}

	sub, i := make([]string, len(set.m)), 0
	for k := range set.m {
		sub[i] = k
		i++
	}

	sort.Sort(sort.StringSlice(sub))
	for i, s := range sub {
		fmt.Printf("%3d: %s\n", i, s)
	}

	return len(set.m)
}

func (q Q2062) IsVowel(b byte) bool {
	if b != 'a' && b != 'e' && b != 'i' && b != 'o' && b != 'u' {
		return false
	}
	return true
}

func (q Q2062) ExtractAEIOUSubStrings(word string, minlen int) []string {
	subwords := []string{}
	var i, j int
	for i = 0; i < len(word); i++ {
		if q.IsVowel(word[i]) {
			for j = i + 1; j < len(word); j++ {
				if !q.IsVowel(word[j]) {
					break
				}
			}
			if j-i >= 5 {
				subwords = append(subwords, word[i:j])
			}
			i = j - 1
		}
	}
	return subwords
}

type Q2062Set struct {
	m       map[string]struct{}
	v       []string
	onlylen bool
}

func NewQ2062Set(onlylen bool) *Q2062Set {
	return &Q2062Set{
		m:       map[string]struct{}{},
		v:       []string{},
		onlylen: onlylen,
	}
}

func (st *Q2062Set) Append(s string) bool {
	if _, ok := st.m[s]; ok {
		return false
	}

	st.m[s] = struct{}{}
	if !st.onlylen {
		st.v = append(st.v, s)
	}
	return true
}

type AEIOU struct {
	c []*Q2062Char
}

func NewAEIOU() *AEIOU {
	return &AEIOU{
		c: []*Q2062Char{
			{
				char:     'a',
				position: []int{},
			},
			{
				char:     'e',
				position: []int{},
			},
			{
				char:     'i',
				position: []int{},
			},
			{
				char:     'o',
				position: []int{},
			},
			{
				char:     'u',
				position: []int{},
			},
		},
	}
}

func (a *AEIOU) Append(i int, b byte) bool {
	switch b {
	case 'a':
		a.c[0].position = append(a.c[0].position, i)
	case 'e':
		a.c[1].position = append(a.c[1].position, i)
	case 'i':
		a.c[2].position = append(a.c[2].position, i)
	case 'o':
		a.c[3].position = append(a.c[3].position, i)
	case 'u':
		a.c[4].position = append(a.c[4].position, i)
	default:
		return false
	}
	return true
}

func (a *AEIOU) HasAll() bool {
	for i := 0; i < 5; i++ {
		if len(a.c[i].position) < 1 {
			return false
		}
	}
	return true
}

type SortByFstE []*Q2062Char

func (a SortByFstE) Len() int      { return len(a) }
func (a SortByFstE) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByFstE) Less(i, j int) bool {
	if len(a[i].position) == 0 {
		return true
	}

	if len(a[j].position) == 0 {
		return false
	}

	return a[i].position[0] < a[j].position[0]
}

func (a *AEIOU) Sort() {
	sort.Sort(SortByFstE(a.c))
}

func (a AEIOU) String() string {
	sb := strings.Builder{}
	sb.WriteString("AEIOU\n{\n")
	for i := range a.c {
		sb.WriteString(
			fmt.Sprintf(
				"\t%c: %d %v\n",
				a.c[i].char,
				len(a.c[i].position),
				a.c[i].position))
	}
	sb.WriteString("}")
	return sb.String()
}

func (a *AEIOU) NextAEIOUInterval() (start int, end int, ok bool) {
	if !a.HasAll() || len(a.c[0].position) == 0 {
		return 0, 0, false
	}

	start = a.c[0].position[0]
	end = a.c[4].position[0]

	a.c[0].position = a.c[0].position[1:]
	if len(a.c[0].position) > 0 {
		for i := 0; i < 4; i++ {
			if a.c[i].position[0] > a.c[i+1].position[0] {
				a.c[i], a.c[i+1] = a.c[i+1], a.c[i]
			}
		}
	}
	return start, end, true
}

func (a *AEIOU) AEIOUSubStrings(str string) []string {
	set := NewQ2062Set(false)

	for true {
		if s, e, ok := a.NextAEIOUInterval(); ok {
			for i := e; i < len(str); i++ {
				set.Append(str[s : i+1])
			}
		} else {
			break
		}
	}
	return set.v
}

func (a *AEIOU) NumAEIOUSubStrings(str string) int {
	set := NewQ2062Set(false)

	for true {
		if s, e, ok := a.NextAEIOUInterval(); ok {
			for i := e; i < len(str); i++ {
				set.Append(str[s : i+1])
			}
		} else {
			break
		}
	}
	return len(set.m)
}
