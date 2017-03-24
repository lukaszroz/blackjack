package main

const LIMIT = 21

type Score struct {
	value int
	deductions []int
}

func NewScore() Score {
	return Score{deductions: []int{}}
}

func (s *Score) AddCard(c Card) {
	s.value += c.highValue
	if c.highValue != c.lowValue {
		s.deductions = append(s.deductions, c.highValue - c.lowValue)
	}
	for s.value > LIMIT && len(s.deductions) > 0 {
		last := len(s.deductions) - 1
		s.value -= s.deductions[last]
		s.deductions = s.deductions[:last]
	}
}