package leetcode

import "fmt"

type node struct {
	next *node
	prev *node
	val  int
	step int
}

type queue struct {
	Head   *node
	Tail   *node
	Length int
}

func NewQueue() *queue {
	return &queue{Length: 0}
}

func (q *queue) IsEmpty() bool {
	return q.Length == 0
}

// removes and returns the topmost element from the stackqueue
func (q *queue) Dequeue() (val int, step int, ok bool) {
	if q.IsEmpty() {
		return 0, -1, false
	}
	newHead := q.Head.next // might be nil
	oldHead := q.Head

	q.Head = newHead
	q.Length -= 1
	if q.Length == 0 {
		q.Tail = nil
	}

	val, step, ok = oldHead.val, oldHead.step, true
	return
}

func (q *queue) Enqueue(val, step int) *queue {
	newTail := &node{val: val, step: step}
	oldTail := q.Tail

	q.Tail = newTail
	if q.Length == 0 {
		q.Head = newTail
	} else {
		newTail.prev = oldTail
		oldTail.next = newTail
	}
	q.Length += 1
	return q
}

func brokenCalc(startValue int, target int) int {
	if startValue >= target {
		return startValue - target
	}

	hasVisited := map[int]bool{}
	q := NewQueue()
	q.Enqueue(target, 0)
	var v, d2, p1, step int // divided by two, plus one

	for !q.IsEmpty() {
		v, step, _ = q.Dequeue()
		step++
		// divided by 2
		if v%2 == 0 {
			d2 = v / 2
		} else {
			d2 = 0
		}
		// plus one
		p1 = v + 1

		fmt.Println("step:", step)
		fmt.Printf("d2: %d p1: %d\n", d2, p1)
		if d2 == startValue || p1 == startValue {
			return step
		}

		if d2 > 0 && !hasVisited[d2] {
			q.Enqueue(d2, step)
			hasVisited[d2] = true
		}

		if p1 <= target+1 && !hasVisited[p1] {
			q.Enqueue(p1, step)
			hasVisited[p1] = true
		}
	}
	return -1
}
