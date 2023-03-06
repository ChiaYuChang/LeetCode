package leetcode

import (
	"bytes"
	"sort"
)

type q405 struct{}

func NewQ405() q405 {
	return q405{}
}

type q405Data struct {
	char  byte
	count int
}

type Sortq405DataByCount []*q405Data

func (a Sortq405DataByCount) Len() int      { return len(a) }
func (a Sortq405DataByCount) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Sortq405DataByCount) Less(i, j int) bool {
	if a[i].count == a[j].count {
		return a[i].char < a[j].char
	}
	return a[i].count < a[j].count
}

type Sortq405DataByChar []*q405Data

func (a Sortq405DataByChar) Len() int           { return len(a) }
func (a Sortq405DataByChar) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Sortq405DataByChar) Less(i, j int) bool { return a[i].char < a[j].char }

func RemoveKdigits(num string, k int) string {
	data := make([]*q405Data, 26)
	for i := range data {
		data[i] = &q405Data{char: byte('0' + i), count: 0}
	}

	for i := 0; i < len(num); i++ {
		data[num[i]-'0'].count += 1
	}

	sort.Sort(sort.Reverse(Sortq405DataByCount(data)))
	data = data[k:]
	sort.Sort(Sortq405DataByChar(data))

	bs := bytes.NewBufferString("")
	for _, d := range data {
		for i := 0; i < d.count; i++ {
			bs.WriteByte(d.char)
		}
	}

	return bs.String()
}
