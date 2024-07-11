package leetcode

import (
	"strconv"
	"strings"
)

type Q1190 struct{}

type Q1190Stack struct {
	status int
	data   []int
}

func (q Q1190) NewStack(cap int) *Q1190Stack {
	return &Q1190Stack{
		status: -1,
		data:   make([]int, 0, cap),
	}
}

func (s *Q1190Stack) Push(n int) {
	s.status += 1
	if len(s.data) > s.status {
		s.data[s.status] = n
	} else {
		s.data = append(s.data, n)
	}
}

func (s Q1190Stack) Peek() (int, bool) {
	if s.status < 0 {
		return 0, false
	}
	return s.data[s.status], true
}

func (s *Q1190Stack) Pop() (int, bool) {
	if s.status < 0 {
		return 0, false
	}
	value := s.data[s.status]
	s.status -= 1
	return value, true
}

func (s Q1190Stack) String() string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	for i := s.status; i >= 0; i-- {
		sb.WriteString(strconv.Itoa(s.data[i]))
		if i > 0 {
			sb.WriteString(", ")
		}
	}
	sb.WriteRune(']')
	return sb.String()
}

func (q Q1190) ReverseParentheses(s string) string {
	jumpTo := map[int]int{}
	strLen := 0
	stack := q.NewStack(10)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			stack.Push(i)
		} else if s[i] == '(' {
			j, _ := stack.Pop()
			jumpTo[i], jumpTo[j] = j, i
		} else {
			strLen += 1
		}
	}

	builder := make([]byte, 0, strLen)
	for i, direction := 0, 1; i < len(s); i += direction {
		if j, ok := jumpTo[i]; !ok {
			builder = append(builder, s[i])
		} else {
			i = j
			if top, ok := stack.Peek(); ok && i == top {
				stack.Pop()
			} else {
				stack.Push(i)
			}
			direction *= -1
		}
	}
	return string(builder)
}
