package leetcode

type Q1457 struct{}

type Q1457TreeNode struct {
	Val   int
	Left  *Q1457TreeNode
	Right *Q1457TreeNode
}

type Q1457PalindromicPath [9]int

func (p Q1457PalindromicPath) Copy() Q1457PalindromicPath {
	newP := [9]int{}
	copy(newP[:], p[:])
	return newP
}

func (p Q1457PalindromicPath) IsPalindrome() bool {
	oneOddVal := true
	for i := 0; i < 9; i++ {
		if p[i]%2 != 0 {
			if oneOddVal {
				oneOddVal = false
				continue
			}
			return false
		}
	}
	return true
}

func (q Q1457) PseudoPalindromicPaths(root *Q1457TreeNode) int {
	if root == nil {
		panic("tree should contain a least one node")
	}
	npp := 0
	q.Travel(root, [9]int{}, &npp)
	return npp
}

func (q Q1457) Travel(node *Q1457TreeNode, pp Q1457PalindromicPath, npp *int) {
	lIsNil, rIsNil := node.Left == nil, node.Right == nil

	pp[node.Val-1]++
	if lIsNil && rIsNil {
		if pp.IsPalindrome() {
			(*npp)++
		}
		return
	}

	if !lIsNil {
		q.Travel(node.Left, pp.Copy(), npp)
	}

	if !rIsNil {
		q.Travel(node.Right, pp.Copy(), npp)
	}
}
