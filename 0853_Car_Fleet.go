package leetcode

import (
	"fmt"
	"sort"
)

type Q0853 struct{}

type Q0853data struct {
	pos int // position
	afr int // arrive after
	res *Q0853Fraction
}

func (d1 Q0853data) LessOfEqualAfr(d2 Q0853data) bool {
	if d1.afr != d2.afr {
		return d1.afr <= d2.afr
	}
	return d1.res.Equal(d2.res) || d1.res.Less(d2.res)
}

func (d Q0853data) String() string {
	return fmt.Sprintf("t0 Position: %3d Arrive After: %3d %v", d.pos, d.afr, d.res)
}

// Sort Interface
type SortQ0853DataByPos []Q0853data

func (a SortQ0853DataByPos) Len() int {
	return len(a)
}

func (a SortQ0853DataByPos) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a SortQ0853DataByPos) Less(i, j int) bool {
	if a[i].pos != a[j].pos {
		return a[i].pos < a[j].pos
	}
	return a[i].LessOfEqualAfr(a[j])
}

// Q0853Fraction
type Q0853Fraction struct {
	// []int{numerator denominator}
	frac         [2]int
	isNormalized bool
}

func (q Q0853) NewFraction(numerator, denominator int) *Q0853Fraction {
	f := &Q0853Fraction{frac: [2]int{numerator, denominator}}
	f = f.normalize()
	return f
}

func (f *Q0853Fraction) normalize() *Q0853Fraction {
	if f.frac[0] == 0 {
		f.frac[1] = 0
		f.isNormalized = true
		return f
	}

	for i := 2; i < f.frac[1]; i++ {
		if f.frac[0]%i == 0 && f.frac[1]%i == 0 {
			f.frac[0] /= i
			f.frac[1] /= i
		}
		if i*i > f.frac[1] {
			break
		}
	}
	f.isNormalized = true
	return f
}

func (f *Q0853Fraction) ToFloat() float64 {
	if f.frac[1] == 0 {
		return 0
	}
	return float64(f.frac[0]) / float64(f.frac[1])
}

func (f1 *Q0853Fraction) Equal(f2 *Q0853Fraction) bool {
	if f1.isNormalized {
		f1.normalize()
	}

	if f2.isNormalized {
		f2.normalize()
	}
	return f1.frac == f2.frac
}

func (f1 *Q0853Fraction) Less(f2 *Q0853Fraction) bool {
	if f1.Equal(f2) {
		return false
	}
	return f1.ToFloat() < f2.ToFloat()
}

func (f Q0853Fraction) String() string {
	return fmt.Sprintf("%d//%d", f.frac[0], f.frac[1])
}

func (q Q0853) CarFleet(target int, position []int, speed []int) int {
	l := len(position)
	ds := make([]Q0853data, l)
	for i := 0; i < l; i++ {
		ds[i] = Q0853data{
			pos: position[i],
			afr: (target - position[i]) / speed[i],
			res: q.NewFraction((target-position[i])%speed[i], speed[i]),
		}
	}
	sort.Sort(sort.Reverse(SortQ0853DataByPos(ds)))

	nfleet := 1
	for i := 1; i < l; i++ {
		if ds[i].LessOfEqualAfr(ds[i-1]) {
			ds[i].afr = ds[i-1].afr
			ds[i].res.frac = ds[i-1].res.frac
		} else {
			nfleet += 1
		}
	}

	return nfleet
}
