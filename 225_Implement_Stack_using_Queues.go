package leetcode

type MyStack struct {
	q1   []int
	q2   []int
	flag bool
}

func NewMyStack() MyStack {
	return MyStack{
		q1:   make([]int, 0),
		q2:   nil,
		flag: true,
	}
}

func (this *MyStack) Push(x int) {
	if this.flag {
		this.q1 = append(this.q1, x)
	} else {
		this.q2 = append(this.q2, x)
	}
}

func (this *MyStack) Pop() int {
	lastElement := 0
	if this.flag {
		this.q2 = make([]int, 0, len(this.q1))
		for true {
			lastElement = this.q1[0]
			this.q1 = this.q1[1:]
			if this.Empty() {
				break
			}
			this.q2 = append(this.q2, lastElement)
		}
		this.q1 = nil
	} else {
		this.q1 = make([]int, 0, len(this.q2))
		for true {
			lastElement = this.q2[0]
			this.q2 = this.q2[1:]
			if this.Empty() {
				break
			}
			this.q1 = append(this.q1, lastElement)
		}
		this.q2 = nil
	}
	this.flag = !this.flag
	return lastElement
}

func (this *MyStack) Top() int {
	e := this.Pop()
	this.Push(e)
	return e
}

func (this *MyStack) Empty() bool {
	if this.flag {
		return len(this.q1) == 0
	}
	return len(this.q2) == 0
}
