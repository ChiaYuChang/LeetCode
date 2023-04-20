package leetcode

type Q1432 struct{}

func (q Q1432) MaxDiff(num int) int {
	if num < 10 {
		return 8
	}

	// first digit, first digit less than 1, first digit greater than 1
	var fstDgt, fstDgtLT9, fstDgtGT1 int
	for n := num; n > 0; {
		fstDgt = n % 10
		if fstDgt != 9 {
			fstDgtLT9 = fstDgt
		}

		if fstDgt != 1 && fstDgt > 0 {
			fstDgtGT1 = fstDgt
		}
		n = n / 10
	}

	x := fstDgt
	if fstDgt == 9 {
		x = fstDgtLT9
	}

	y := fstDgt
	if fstDgt == 1 {
		y = fstDgtGT1
	}

	mxDff := 0
	nDgt := 1
	if x == y {
		for num > 0 {
			if num%10 == fstDgt {
				mxDff += (9 - 1) * nDgt
			}
			nDgt *= 10
			num = num / 10
		}
	} else {
		var d int
		for num > 0 {
			d = num % 10
			if d == x {
				mxDff += (9 - d) * nDgt
			}

			if d == y {
				if y == fstDgt {
					mxDff += (d - 1) * nDgt
				} else {
					mxDff += (d - 0) * nDgt
				}
			}
			nDgt *= 10
			num = num / 10
		}
	}
	return mxDff
}
