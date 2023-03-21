package leetcode

const Q0434Sep byte = ' '

type Q0434 struct{}

type Q0434SafeString struct {
	s *string
}

func (ss Q0434SafeString) Get(i int) byte {
	if i < 0 || i >= ss.Len() {
		return ' '
	}
	return (*ss.s)[i]
}

func (ss Q0434SafeString) Len() int {
	return len(*ss.s)
}

func (q Q0434) CountSegments(s string) int {
	ss := Q0434SafeString{&s}
	n := 0
	for i := 0; i <= ss.Len(); i++ {
		if ss.Get(i) != Q0434Sep {
			for j := i; j <= ss.Len(); j++ {
				if ss.Get(j) == Q0434Sep {
					i = j
					break
				}
			}
			n += 1
		}
	}
	return n
}

func (q Q0434) CountSegmentsPadding(s string) int {
	var n, i, j int
	s += " "
	l := len(s)
	for i = 0; i < l; i++ {
		if s[i] != Q0434Sep {
			for j = i; j < l; j++ {
				if s[j] == Q0434Sep {
					i = j
					break
				}
			}
			n += 1
		}
	}
	return n
}
