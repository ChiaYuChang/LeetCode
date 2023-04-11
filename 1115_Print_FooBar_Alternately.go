package leetcode

import "strings"

type Q1115 struct{}

func (q Q1115) FooBar(n int) string {
	if n == 0 {
		return ""
	}

	c1 := make(chan string)
	c2 := make(chan string)
	defer close(c1)
	defer close(c2)

	go q.foo(n, c1)
	go q.bar(n, c2)

	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteString(<-c1)
		sb.WriteString(<-c2)
	}
	return sb.String()
}

func (q Q1115) foo(n int, c chan<- string) {
	for i := 0; i < n; i++ {
		c <- "foo"
	}
}

func (q Q1115) bar(n int, c chan<- string) {
	for i := 0; i < n; i++ {
		c <- "bar"
	}
}
