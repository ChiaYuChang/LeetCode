package leetcode

type Q0454 struct{}

func (q Q0454) Counts(nums []int) ([]int, []int) {
	val2i := map[int]int{}
	cnts := []int{}

	ln := 0
	for _, val := range nums {
		if idx, ok := val2i[val]; ok {
			cnts[idx]++
		} else {
			cnts = append(cnts, 1)
			nums[ln] = val
			val2i[val] = ln
			ln++
		}
	}
	return nums[:ln], cnts
}

func (q Q0454) Cross(num1, num2, cnt1, cnt2 []int) map[int]int {
	s := map[int]int{}
	for i1, n1 := range num1 {
		for i2, n2 := range num2 {
			s[n1+n2] += cnt1[i1] * cnt2[i2]
		}
	}
	return s
}

func (q Q0454) FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var cnts1, cnts2, cnts3, cnts4 []int
	nums1, cnts1 = q.Counts(nums1)
	nums2, cnts2 = q.Counts(nums2)
	s1 := q.Cross(nums1, nums2, cnts1, cnts2)

	nums3, cnts3 = q.Counts(nums3)
	nums4, cnts4 = q.Counts(nums4)
	s2 := q.Cross(nums3, nums4, cnts3, cnts4)

	nqps := 0
	for sum := range s1 {
		nqps += s1[sum] * s2[-sum]
	}
	return nqps
}
