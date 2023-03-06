package leetcode

func SearchMatrixII(matrix [][]int, target int, method string) bool {
	switch method {
	case "fast", "f":
		// O(m+n)
		return fastSearchMatrixII(matrix, target)
	case "slow", "s":
		// O(mlog(n) or nlog(m))
		return slowSearchMatrixII(matrix, target)
	default:
		return false
	}
}

func fastSearchMatrixII(matrix [][]int, target int) bool {
	row, col := len(matrix)-1, 0

	for row > -1 && col < len(matrix[0]) {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else {
			return true
		}
	}
	return false
}

func slowSearchMatrixII(matrix [][]int, target int) bool {
	nrow, ncol := len(matrix), len(matrix[0])

	if nrow > ncol {
		for col := 0; col < ncol; col++ {
			if matrix[0][col] <= target && target <= matrix[nrow-1][col] {
				l, r := 0, nrow-1
				for r-l > 10 {
					m := (l + r + 1) / 2
					if matrix[m][col] > target {
						r = m
					} else if matrix[m][col] < target {
						l = m
					} else {
						return true
					}
				}
				for i := l; i <= r; i++ {
					if matrix[i][col] == target {
						return true
					}
				}
			}
		}
	} else {
		for row := 0; row < nrow; row++ {
			if matrix[row][0] <= target && target <= matrix[row][ncol-1] {
				l, r := 0, ncol-1
				for r-l > 10 {
					m := (l + r + 1) / 2
					if matrix[row][m] > target {
						r = m
					} else if matrix[row][m] < target {
						l = m
					} else {
						return true
					}
				}
				for i := l; i <= r; i++ {
					if matrix[row][i] == target {
						return true
					}
				}
			}
		}
	}
	return false
}
