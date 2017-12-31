package main

type stack struct {
	vals []uint16
	sp   int
}

func newStack() *stack {
	return &stack{[]uint16{}, -1}
}

func (s *stack) push(v uint16) {
	s.sp++
	if s.sp < len(s.vals)+1 {
		s.vals = append(s.vals, v)
	} else {
		s.vals[s.sp] = v
	}
}

func (s *stack) pop() (uint16, bool) {
	if s.sp < 0 {
		return 0, false
	}
	v := s.vals[s.sp]
	s.sp--
	return v, true
}
