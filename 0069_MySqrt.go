package leetcode

type Q0069 struct{}

type Q0069IntPairs []int

func (ip Q0069IntPairs) reverse() {
	for i := 0; i < len(ip)/2; i++ {
		ip[i], ip[len(ip)-1-i] = ip[len(ip)-1-i], ip[i]
	}
}

func (ip Q0069IntPairs) sqrt() []int {
	value := make([]int, len(ip))
	i, j := 0, 0
	res := 0
	a := 0
	b := 0
	for _, i = range ip {
		res = res*100 + i
		for b = 0; b < 10; b++ {
			if (a+b+1)*(b+1) > res {
				break
			}
		}
		value[j] = b
		j++
		res -= (a + b) * b
		a = (a + b*2) * 10
	}
	return value
}

func (q Q0069) int2intPairs(x int) Q0069IntPairs {
	ip := []int{}
	for x > 0 {
		ip = append(ip, x%100)
		x /= 100
	}
	Q0069IntPairs(ip).reverse()
	return ip
}

func (q Q0069) MySqrt(x int) int {
	ip := q.int2intPairs(x)
	value := 0
	for _, v := range ip.sqrt() {
		value = value*10 + v
	}
	return value
}
