package leetcode

func ReverseBits(num uint32) uint32 {
	var mun uint32 = 0

	i := 32
	for num > 0 {
		mun = mun << 1
		if num&1 == 1 {
			mun |= 1
		}
		num = num >> 1
		i--
	}

	if i > 0 {
		mun = mun << i
	}

	return mun
}
