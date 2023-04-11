package leetcode

type Q0971 struct{}

type Q0971TreeNode struct {
	Val   int
	Left  *Q0971TreeNode
	Right *Q0971TreeNode
}

func (q Q0971) FlipMatchVoyage(root *Q0971TreeNode, voyage []int) []int {
	i := 1
	flipwhich := []int{}
	if root.Val != voyage[0] || !q.travenAndFlip(root, &i, voyage, &flipwhich) {
		return []int{-1}
	}
	return flipwhich
}

func (q Q0971) travenAndFlip(node *Q0971TreeNode, i *int, voyage []int, flipWhich *[]int) bool {
	var ok, hasL, hasR bool = true, node.Left != nil, node.Right != nil

	if hasL {
		// fmt.Printf("-  L %d - v[%d] %d\n", node.Left.Val, *i, voyage[*i])
		if node.Left.Val == voyage[*i] {
			(*i)++
			ok = q.travenAndFlip(node.Left, i, voyage, flipWhich)
		} else {
			// fmt.Printf("-f R %d - v[%d] %d\n", node.Right.Val, *i, voyage[*i])
			if hasR && node.Right.Val == voyage[*i] {
				node.Left, node.Right = node.Right, node.Left
				(*i)++
				(*flipWhich) = append((*flipWhich), node.Val)
				ok = q.travenAndFlip(node.Left, i, voyage, flipWhich)
			} else {
				ok = false
			}
		}
	}

	if ok && hasR {
		// fmt.Printf("-  R %d - v[%d] %d\n", node.Right.Val, *i, voyage[*i])
		if node.Right.Val == voyage[*i] {
			(*i)++
			ok = q.travenAndFlip(node.Right, i, voyage, flipWhich)
		} else {
			ok = false
		}
	}
	return ok
}
