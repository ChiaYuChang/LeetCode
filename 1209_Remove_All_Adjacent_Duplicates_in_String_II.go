package leetcode

import (
	"fmt"
	"strings"
)

type Q1209 struct{}

type Q1209ByteSlice struct {
	bytes   []byte
	counter []int
	end     int
}

const Q1209AlwayNotMatchChar = byte('*') // a char that does not match any char in s

func NewByteSlice(s string, k int) Q1209ByteSlice {
	bs := Q1209ByteSlice{
		bytes:   make([]byte, len(s)+1),
		counter: make([]int, len(s)+1),
		end:     1,
	}
	bs.bytes[0] = Q1209AlwayNotMatchChar
	bs.counter[0] = k - 1
	return bs
}

func (bs Q1209ByteSlice) Prev() (byte, int) {
	return bs.bytes[bs.end-1], bs.counter[bs.end-1]
}

func (bs *Q1209ByteSlice) Append(b byte) {
	bs.bytes[bs.end] = b
	if b == bs.bytes[bs.end-1] {
		bs.counter[bs.end] = bs.counter[bs.end-1] + 1
	} else {
		bs.counter[bs.end] = 1
	}
	bs.end++
}

func (bs *Q1209ByteSlice) Rewind(k int) {
	bs.end = bs.end - k + 1
}

func (bs Q1209ByteSlice) String() string {
	sb := strings.Builder{}
	sb.WriteString("ByteSlice: [")
	for i := 1; i < bs.end; i++ {
		sb.WriteString(fmt.Sprintf("%c(%d)", bs.bytes[i], bs.counter[i]))
		if i != bs.end-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteByte(']')
	return sb.String()
}

func (bs Q1209ByteSlice) ToString() string {
	return string(bs.bytes[1:bs.end])
}

func (q Q1209) RemoveDuplicates(s string, k int) string {
	bs := NewByteSlice(s, k)

	prevChar, precCount := bs.Prev()
	for i := 0; i < len(s); i++ {
		// fmt.Println(bs)
		if s[i] != prevChar {
			bs.Append(s[i])
		} else {
			if precCount+1 < k {
				bs.Append(s[i])
			} else {
				bs.Rewind(k)
			}
		}
		prevChar, precCount = bs.Prev()
	}
	return bs.ToString()
}
