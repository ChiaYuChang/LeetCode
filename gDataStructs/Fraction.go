package gdatastructs

import "fmt"

type Fraction struct {
	// []int{numerator denominator}
	frac         [2]int
	isNormalized bool
}

func NewFraction(numerator, denominator int) *Fraction {
	f := &Fraction{frac: [2]int{numerator, denominator}}
	f = f.normalize()
	return f
}

func (f *Fraction) normalize() *Fraction {
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

func (f *Fraction) ToFloat() float64 {
	if f.frac[1] == 0 {
		return 0
	}
	return float64(f.frac[0]) / float64(f.frac[1])
}

func (f1 *Fraction) Equal(f2 *Fraction) bool {
	if f1.isNormalized {
		f1.normalize()
	}

	if f2.isNormalized {
		f2.normalize()
	}
	return f1.frac == f2.frac
}

func (f1 *Fraction) Less(f2 *Fraction) bool {
	if f1.Equal(f2) {
		return false
	}
	return f1.ToFloat() < f2.ToFloat()
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d//%d", f.frac[0], f.frac[1])
}
