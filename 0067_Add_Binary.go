package leetcode

type Q0067 struct{}

func (q Q0067) AddBinary(a, b string) string {
	if a == "" || a == "0" {
		return b
	}

	if b == "" || b == "0" {
		return a
	}

	cmap := []bool{false, false, true, true}
	dmap := []byte{'0', '1', '0', '1'}

	alen := len(a)
	blen := len(b)

	l := alen
	if alen < blen {
		l = blen
	}

	var c bool = false
	ans := make([]byte, l+1, l+1)
	for i := 0; i < l; i++ {
		var k int8 = 0
		if c {
			k++
		}
		if alen-i-1 >= 0 && a[alen-i-1] == '1' {
			k++
		}
		if blen-i-1 >= 0 && b[blen-i-1] == '1' {
			k++
		}

		c = cmap[k]
		ans[l-i] = dmap[k]
	}

	if c {
		ans[0] = '1'
		return string(ans)
	}
	return string(ans[1:])
}
