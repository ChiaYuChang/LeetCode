package leetcode

type MyQueue struct {
	s1   []int
	s2   []int
	head int
}

func Constructor() MyQueue {
	return MyQueue{
		s1: make([]int, 0),
		s2: nil,
	}
}

func (this *MyQueue) Push(x int) {
	if this.Empty() {
		this.head = x
	}
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}

	var e int
	this.s2 = make([]int, 0, len(this.s1)-1)
	for true {
		e = this.s1[len(this.s1)-1]
		this.s1 = this.s1[:len(this.s1)-1]
		if this.Empty() {
			break
		}
		this.s2 = append(this.s2, e)
	}

	this.s1 = make([]int, 0, len(this.s2))
	this.head = 0
	for len(this.s2) > 0 {
		tmp := this.s2[len(this.s2)-1]
		if this.Empty() {
			this.head = tmp
		}
		this.s2 = this.s2[:len(this.s2)-1]
		this.s1 = append(this.s1, this.s2[len(this.s2)-1])
	}
	this.s2 = nil
	return e
}

func (this *MyQueue) Peek() int {
	return this.head
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0
}
