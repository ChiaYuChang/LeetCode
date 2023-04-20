package leetcode

import "fmt"

type Q2584 struct{}

func FindValidSplit(nums []int) int {
	// cache := map[int]int{}
	pIs := map[int][]int{}
	fsi := make([][]int, len(nums))

	// ftr := NewFactorizer()
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	for p := 2; p <= max; p++ {
		lmax := -1
		var ok bool
		for i, num := range nums {
			if nums[i], ok = HasPrime(num, p); ok {
				if _, ok = pIs[p]; ok {
					pIs[p][1] = i
				} else {
					pIs[p] = []int{i, i}
				}
				fsi[i] = append(fsi[i], p)
			}

			if lmax < 0 || lmax < nums[i] {
				lmax = nums[i]
			}

		}
		max = lmax
	}

	end := 0
	for _, p := range fsi[0] {
		if pIs[p][1] > end {
			end = pIs[p][1]
		}
	}

	for i := 1; i < len(nums)-1 && i <= end; i++ {
		for _, p := range fsi[i] {
			if pIs[p][1] > end {
				end = pIs[p][1]
			}
		}
	}

	if end == len(nums)-1 {
		return -1
	}
	return end
}

type Factorizer struct {
	primte map[int]bool
}

func NewFactorizer() *Factorizer {
	return &Factorizer{primte: make(map[int]bool)}
}

func (ftr *Factorizer) factors(n int) []int {
	if n == 1 {
		return []int{}
	}

	fmt.Print("n:", n, ", f:")

	if ftr.isPrime(n) {
		fmt.Printf("[%d]\n", n)
		return []int{n}
	}

	f := []int{}
	var ok bool
	for i := 2; i <= n && n > 1; i++ {
		if n, ok = HasPrime(n, i); ok {
			f = append(f, i)
		}
	}
	fmt.Println(f)
	return f
}

func (ftr *Factorizer) isPrime(n int) bool {
	if v, ok := ftr.primte[n]; ok {
		return v
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			ftr.primte[n] = false
			return false
		}
	}
	ftr.primte[n] = true
	return true
}

func HasPrime(num, prm int) (int, bool) {
	// cnt := 0
	ok := false
	pwr := 1

	p := prm
	for num%p == 0 {
		num /= p
		// cnt += pwr
		ok = true
		pwr *= 2
		p *= p
	}

	for p, pwr = p/prm, pwr-1; num > 1 && p > 1; p, pwr = p/prm, pwr-1 {
		for p < num && num%p == 0 {
			num /= p
			// cnt += pwr
		}
	}
	// return num, cnt
	return num, ok
}
