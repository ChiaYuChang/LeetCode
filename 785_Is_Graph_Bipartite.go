package leetcode

import (
	"fmt"

	gds "gitlab.com/gjerry134679/leetcode/gDataStructs"
)

type TColor [2]bool

var (
	Unknown TColor = [2]bool{false, false}
	Red     TColor = [2]bool{true, true}
	Black   TColor = [2]bool{true, false}
)

func (c TColor) Flip() TColor {
	switch c {
	case Red:
		return Black
	case Black:
		return Red
	default:
		return Unknown
	}
}

func (c TColor) Equal(d TColor, omitUnknown bool) bool {
	if omitUnknown && (c == Unknown || d == Unknown) {
		return true
	}
	return c[1] == d[1]
}

func (c TColor) NotEqual(d TColor, omitUnknown bool) bool {
	if omitUnknown && (c == Unknown || d == Unknown) {
		return true
	}
	return c != d
}

func (c TColor) String() string {
	switch c {
	case Red:
		return "Red"
	case Black:
		return "Black"
	default:
		return "Unknown"
	}
}

type Graph struct {
	color   []TColor
	forward [][]int
	reverse [][]int
	noin    []int
	noout   []int
	n       []int
	orphant []int
}

type Set struct {
	len int
	set map[int]struct{}
}

// add one element to the set, return false if the element has
// already existed
func (s *Set) AddOne(i int) bool {
	_, ok := s.set[i]
	if ok {
		return false
	}
	s.set[i] = struct{}{}
	s.len++
	return true
}

// add elements to the set
func (s *Set) Add(i ...int) {
	for _, j := range i {
		s.AddOne(j)
	}
}

// remove given element from the set if exist
func (s *Set) RemoveOne(i int) {
	if s.IsEmpty() {
		return
	}

	_, ok := s.set[i]
	if !ok {
		return
	}
	delete(s.set, i)
	s.len--
}

// remove given elements from the set if exit
func (s *Set) Remove(i ...int) {
	for _, j := range i {
		s.RemoveOne(j)
	}
}

// return whether the set is empty
func (s *Set) IsEmpty() bool {
	return s.len == 0
}

func (s *Set) In(i int) bool {
	_, ok := s.set[i]
	return ok
}

// convert the set to an array
func (s *Set) ToList() []int {
	i := make([]int, 0, s.len)
	for k := range s.set {
		i = append(i, k)
	}
	return i
}

// create a new set struct
func NewSet() *Set {
	return &Set{len: 0, set: map[int]struct{}{}}
}

// create a new graph struct
func NewGraph(g [][]int) *Graph {
	c := make([]TColor, len(g))
	for i := range c {
		c[i] = Unknown
	}

	grph := &Graph{
		color:   c,
		forward: g,
		reverse: make([][]int, len(g)),
		noin:    nil,
		noout:   nil,
		n:       nil,
		orphant: nil,
	}

	noout := NewSet()
	n := NewSet()

	for i, ns := range grph.forward {
		for _, n := range ns {
			grph.reverse[n] = append(grph.reverse[n], i)
		}

		if len(ns) == 0 {
			noout.Add(i)
		} else {
			n.Add(i)
		}
	}

	for i, ns := range grph.reverse {
		if len(ns) == 0 {
			if noout.In(i) {
				noout.RemoveOne(i)
				grph.orphant = append(grph.orphant, i)
			} else {
				grph.noin = append(grph.noin, i)
				n.Remove(i)
			}
		}
	}
	grph.n = n.ToList()
	grph.noout = noout.ToList()
	return grph
}

func (grph *Graph) ReInitialize() *Graph {
	return NewGraph(grph.forward)
}

func (grph *Graph) pop(which string) (int, bool) {
	var i int
	switch which {
	case "noin":
		if len(grph.noin) == 0 {
			return 0, false
		}
		i = grph.noin[0]
		grph.noin = grph.noin[1:]
	case "noout":
		if len(grph.noout) == 0 {
			return 0, false
		}
		i = grph.noout[0]
		grph.noout = grph.noout[1:]
	case "n":
		if len(grph.n) == 0 {
			return 0, false
		}
		i = grph.n[0]
		grph.n = grph.n[1:]
	default:
		return 0, false
	}
	return i, true
}

func (grph *Graph) Pop() (int, bool) {
	i, ok := grph.pop("noin")
	if ok {
		return i, true
	}

	i, ok = grph.pop("n")
	if ok {
		return i, true
	}
	return 0, false
}

// colorize the ith node with given color, return false if the
// the node has been colorized with different color.
func (grph *Graph) Colorize(i int, c TColor) bool {
	if grph.HasVisited(i) {
		return grph.color[i] == c
	}
	grph.color[i] = c
	return true
}

// return the color of ith node
func (grph *Graph) Color(i int) TColor {
	return grph.color[i]
}

// check the given nodes are the same color
func (grpc *Graph) IsHomogenious(omitUnknown bool, i ...int) (bool, TColor) {
	if len(i) == 0 {
		return true, Unknown
	}

	c := grpc.Color(i[0])
	for j := 1; j < len(i); j++ {
		if !c.Equal(grpc.Color(j), omitUnknown) {
			return false, Unknown
		}
	}
	return true, c
}

// return whether the node has been access before
func (grph *Graph) HasVisited(i int) bool {
	return grph.color[i][0]
}

// return neighbors of givn nodes
func (grph *Graph) Neighbors(rev bool, i ...int) []int {
	s := NewSet()
	for _, j := range i {
		if rev {
			s.Add(grph.reverse[j]...)
		} else {
			s.Add(grph.forward[j]...)
		}
	}
	return s.ToList()
}

// return not visited neighbors of given nodes
func (grph *Graph) NotVisitedNeighbors(rev bool, i ...int) ([]int, bool) {
	s := NewSet()
	c := grph.Color(i[0])
	for _, j := range i {
		if rev {
			for _, n := range grph.reverse[j] {
				if !grph.HasVisited(n) {
					s.Add(n)
				} else {
					if c.Equal(grph.Color(n), true) {
						return nil, false
					}
				}
			}
		} else {
			for _, n := range grph.forward[j] {
				if !grph.HasVisited(n) {
					s.Add(n)
				} else {
					if c.Equal(grph.Color(n), true) {
						return nil, false
					}
				}
			}
		}
	}
	return s.ToList(), true
}

func GarphPartition(g [][]int) [][]int {
	cc := gds.NewConnectedComponent(len(g))
	for i, neighbors := range g {
		for _, n := range neighbors {
			cc.Connect(i, n)
		}
	}
	return cc.Groups()
}

func IsBipartite(graph [][]int) bool {
	grph := NewGraph(graph)

	color := Red
	for true {

	}
	seed, ok := grph.Pop()
	if !ok {
		return true
	}

	for ok {
		if grph.Color(seed) != Unknown {
			color = grph.Color(seed)
		}
		_ = grph.Colorize(seed, color)
		neighbors := grph.Neighbors(false, seed)
		for len(neighbors) > 0 {
			var ok bool
			color = color.Flip()
			for _, n := range neighbors {
				ok = grph.Colorize(n, color)
				if !ok {
					return false
				}
			}
			neighbors, ok = grph.NotVisitedNeighbors(false, neighbors...)
			if !ok {
				return false
			}
		}
		seed, ok = grph.Pop()
	}

	for _, noin := range grph.noin {
		if grph.Color(noin) == Unknown {
			ok, c := grph.IsHomogenious(true, grph.Neighbors(false, noin)...)
			if !ok {
				return false
			}
			grph.Colorize(noin, c)
		}
	}

	for _, noout := range grph.noout {
		if grph.Color(noout) == Unknown {
			ok, c := grph.IsHomogenious(true, grph.Neighbors(true, noout)...)
			if !ok {
				return false
			}
			grph.Colorize(noout, c)
		}
	}

	for i, n := range grph.color {
		fmt.Printf("%2d %v\n", i, n)
	}

	return true
}
