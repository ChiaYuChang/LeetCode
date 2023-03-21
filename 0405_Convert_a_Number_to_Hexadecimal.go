package leetcode

import "bytes"

type Q0405DecimalToHex struct{}

func (d2h Q0405DecimalToHex) Flip(n int) int {
	if n < 0 && n > 15 {
		panic("n should be within [0, 15]")
	}
	return 15 - n
}

func (d2h Q0405DecimalToHex) d2h(n int) byte {
	if n < 0 && n > 15 {
		panic("n should be within [0, 15]")
	}
	var DtoH = map[int]byte{
		0:  '0',
		1:  '1',
		2:  '2',
		3:  '3',
		4:  '4',
		5:  '5',
		6:  '6',
		7:  '7',
		8:  '8',
		9:  '9',
		10: 'a',
		11: 'b',
		12: 'c',
		13: 'd',
		14: 'e',
		15: 'f',
	}
	return DtoH[n]
}

func (d2h Q0405DecimalToHex) ToHex(num int) string {
	if num == 0 {
		return "0"
	}

	sign := num < 0
	if sign {
		num = -num
	}

	// 32/4 = 8
	hx := make([]int, 8)
	for i := 0; num > 0; i++ {
		hx[i] = num % 16
		num /= 16
	}

	if sign {
		// two's complement method
		for i := range hx {
			hx[i] = d2h.Flip(hx[i])
		}

		cry := 1
		for i := 0; i < len(hx); i++ {
			hx[i] += cry
			cry = hx[i] / 16
			hx[i] = hx[i] % 16
		}
	}

	// reverse result
	bs := make([]byte, len(hx))
	for i, h := range hx {
		bs[len(hx)-i-1] = d2h.d2h(h)
	}

	return string(bytes.TrimLeft(bs, "0"))
}
