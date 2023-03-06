package leetcode

func RomanToInt(s string) int {
	if s == "" {
		return 0
	}

	r2i := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	ans := 0

	curr, prev := 0, 1001
	for i := 0; i < len(s); i++ {
		curr, _ = r2i[s[i]]
		if curr > prev {
			ans -= 2 * prev
		}
		ans += curr
		prev = curr
	}

	return ans
}
