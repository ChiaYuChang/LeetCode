package leetcode

import (
	"math"
	"sort"
)

type Q0983 struct{}

func (q Q0983) MinCostTickets(days []int, costs []int) int {
	p := q.NewTicketCostPlanner(costs, []int{1, 7, 30})
	return p.MinCostTicket(days)
}

type Q0983Ticket struct {
	Cost int
	Type int
}

type Q0983TicketCostPlanner struct {
	tickets []Q0983Ticket
	maxTyp  int
	minCst  int
}

func (q Q0983) NewTicketCostPlanner(ticketCost, ticketType []int) Q0983TicketCostPlanner {
	l := len(ticketCost)
	p := Q0983TicketCostPlanner{
		tickets: make([]Q0983Ticket, l),
		minCst:  math.MaxInt,
		maxTyp:  math.MinInt,
	}

	for i := 0; i < l; i++ {
		p.tickets[i] = Q0983Ticket{Cost: ticketCost[i], Type: ticketType[i]}

		if p.minCst > ticketCost[i] {
			p.minCst = ticketCost[i]
		}
		if p.maxTyp < ticketType[i] {
			p.maxTyp = ticketType[i]
		}

	}
	return p
}

func (p Q0983TicketCostPlanner) MinCostTicket(days []int) int {
	if !sort.IsSorted(sort.IntSlice(days)) {
		sort.Sort(sort.IntSlice(days))
	}

	segs := p.Segments(days)

	var str, end int
	c := make(chan int)
	for i := 0; i < len(segs); i++ {
		str, end = end, segs[i]
		go func(c chan<- int, str, end int) {
			c <- p.MinCostTicketSegment(days[str:end])
		}(c, str, end)
	}

	minCost := 0
	for i := 0; i < len(segs); i++ {
		minCost += <-c
	}
	close(c)

	return minCost
}

func (p Q0983TicketCostPlanner) Segments(days []int) []int {
	segs := []int{}
	for i, d := range days {
		if i == 0 {
			continue
		}
		if d-days[i-1]-1 > p.maxTyp {
			segs = append(segs, i)
		}
	}
	segs = append(segs, len(days))
	return segs
}

func (p Q0983TicketCostPlanner) MinCostTicketSegment(seg []int) int {
	if len(seg) == 1 {
		return p.minCst
	}

	// fmt.Println("-", seg)
	minCost := make([]int, len(seg)+1)
	for i := 1; i < len(minCost); i++ {
		minCost[i] = math.MaxInt
	}

	for i, curr := range seg {
		prev := curr - 1
		for _, t := range p.tickets {
			// fmt.Printf("\t- %2d-day ticket\n", t.Type)
			for j := i; j < len(seg) && seg[j] <= prev+t.Type; j++ {
				// fmt.Printf("\t\t- [%d, %d] day: %3d, min(%d, %d)\n", prev+1, prev+t.Type, seg[j], minCost[j+1], minCost[i]+t.Cost)
				minCost[j+1] = p.Min(
					minCost[j+1],
					minCost[i]+t.Cost,
				)
			}
			// fmt.Printf("minCost: %v\n", minCost)
		}
	}
	return minCost[len(seg)]
}

func (p Q0983TicketCostPlanner) Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
