package leetcode

type Q0225 struct{}

type Q0225Stack struct {
	q1   []int
	q2   []int
	flag bool
}

func (q Q0225) NewMyStack() Q0225Stack {
	return Q0225Stack{
		q1:   make([]int, 0),
		q2:   nil,
		flag: true,
	}
}

func (this *Q0225Stack) Push(x int) {
	if this.flag {
		this.q1 = append(this.q1, x)
	} else {
		this.q2 = append(this.q2, x)
	}
}

func (this *Q0225Stack) Pop() int {
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

func (this *Q0225Stack) Top() int {
	e := this.Pop()
	this.Push(e)
	return e
}

func (this *Q0225Stack) Empty() bool {
	if this.flag {
		return len(this.q1) == 0
	}
	return len(this.q2) == 0
}
