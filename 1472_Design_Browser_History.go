package leetcode

import (
	"fmt"
	"strings"
)

type Q1472 struct{}

type Q1472BrowserHistory struct {
	curr int
	n    int
	urls []string
}

func (q Q1472) Constructor(homepage string) Q1472BrowserHistory {
	urls := make([]string, 1)
	urls[0] = homepage

	return Q1472BrowserHistory{
		curr: 0,
		n:    1,
		urls: urls,
	}
}

func (this *Q1472BrowserHistory) Visit(url string) {
	this.curr++
	if len(this.urls) > this.curr {
		this.urls[this.curr] = url
	} else {
		this.urls = append(this.urls, url)
	}
	this.n = this.curr + 1
}

func (this *Q1472BrowserHistory) Back(steps int) string {
	if steps < this.curr {
		this.curr -= steps
	} else {
		this.curr = 0
	}
	return this.urls[this.curr]
}

func (this *Q1472BrowserHistory) Forward(steps int) string {
	if this.curr+steps < this.n {
		this.curr += steps
	} else {
		this.curr = this.n - 1
	}
	return this.urls[this.curr]
}

func (this *Q1472BrowserHistory) String() string {
	sb := strings.Builder{}
	sb.WriteString("Browser History:\n")
	for i := 0; i < this.n; i++ {
		sb.WriteString(fmt.Sprintf("%2d: %s", i, this.urls[i]))
		sb.WriteRune('\n')
	}
	return sb.String()
}
