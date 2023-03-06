package leetcode

type VersionCheckerAPI struct {
	IsBadVersion func(n int) bool
}

func (vc VersionCheckerAPI) FirstBadVersion(n int) int {
	lhs, rhs := 0, n
	for rhs-lhs > 1 {
		mid := (lhs + rhs) / 2
		if vc.IsBadVersion(mid) {
			rhs = mid
		} else {
			lhs = mid
		}
	}
	return rhs
}
