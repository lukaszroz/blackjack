package main

import "encoding/json"

const BUST_LIMIT = 21

type Score struct {
	value      int
	deductions []int
}

func NewScore() Score {
	return Score{deductions: []int{}}
}

func (s *Score) AddCard(c Card) {
	s.value += c.HighValue
	if c.HighValue != c.LowValue {
		s.deductions = append(s.deductions, c.HighValue- c.LowValue)
	}
	for s.value > BUST_LIMIT && len(s.deductions) > 0 {
		last := len(s.deductions) - 1
		s.value -= s.deductions[last]
		s.deductions = s.deductions[:last]
	}
}

func (s Score) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}