package leetcode

type Q1390 struct{}

func (q Q1390) SumFourDivisors(nums []int) int {
	sum := 0
	cache := map[int]int{}
	for _, num := range nums {
		if v, ok := cache[num]; ok {
			sum += v
			continue
		}

		if n, s := q.countDivisors(num, 4); n == 4 {
			sum += s
			cache[num] = s
		} else {
			cache[num] = 0
		}
	}
	return sum
}

func (q Q1390) countDivisors(num int, limit int) (n int, s int) {
	i := 1
	for ; i*i < num; i++ {
		if num%i == 0 {
			s += i + num/i
			n += 2
		}
		if n > limit {
			return -1, 0
		}
	}

	if i*i == num {
		s += i
		n += 1
	}
	return n, s
}
