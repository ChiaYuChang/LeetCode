package leetcode

type Q2058 struct{}

type Q2058ListNode struct {
	Val  int
	Next *Q2058ListNode
}

func (q Q2058) NewListNodeFromList(vals []int) *Q2058ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &Q2058ListNode{}
	curr := head
	for i, v := range vals {
		curr.Val = v
		if i < len(vals)-1 {
			curr.Next = &Q2058ListNode{}
		}
		curr = curr.Next
	}
	return head
}

func (q Q2058) NodesBetweenCriticalPoints(head *Q2058ListNode) []int {
	var n1 *Q2058ListNode = head
	var n2 *Q2058ListNode
	var n3 *Q2058ListNode
	var i = 1
	var prevCP = -1
	var firstCP = -1
	var dist = []int{-1, -1}

	if n1.Next == nil {
		return dist
	}
	n2 = n1.Next

	if n2.Next == nil {
		return dist
	}
	n3 = n2.Next

	for n3 != nil {
		if q.isCriticalpoint(n1.Val, n2.Val, n3.Val) {
			if prevCP > 0 {
				if dist[0] < 0 {
					dist[0] = i - prevCP
				} else {
					dist[0] = min(dist[0], i-prevCP)
				}
			} else {
				firstCP = i
			}
			prevCP = i
		}
		i++
		n1, n2, n3 = n2, n3, n3.Next
	}

	if prevCP != firstCP {
		dist[1] = prevCP - firstCP
	}
	return dist
}

func (q Q2058) isCriticalpoint(n1, n2, n3 int) bool {
	if n1 > n2 && n2 < n3 || n1 < n2 && n2 > n3 {
		return true
	}
	return false
}
