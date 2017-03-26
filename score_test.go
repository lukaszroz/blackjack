package main

import (
	"testing"
)

func assert(s Score, value int, deductions []int, t *testing.T) {
	if value != s.value {
		t.Error("Expected ", value, " value got ", s.value)
		return
	}

	if len(deductions) != len(s.deductions) {
		t.Error("Expected ", deductions, " deductions, got ", s.deductions)
		return
	}

	for i := range deductions {
		if deductions[i] != s.deductions[i] {
			t.Error("Expected ", deductions, " deductions, got ", s.deductions)
			return
		}
	}
}

func ace() Card {
	return Card{"Ace", "", 11, 1}
}

func ten() Card {
	return Card{"Ten", "", 10, 10}
}

func TestTwo(t *testing.T) {
	s := NewScore()
	s.AddCard(Card{"Two", "", 2, 2})
	assert(s, 2, []int{}, t)
}
func TestAce(t *testing.T) {
	s := NewScore()
	s.AddCard(ace())
	assert(s, 11, []int{10}, t)
}

func TestAceWith21Score(t *testing.T) {
	s := NewScore()
	s.AddCard(ten())
	s.AddCard(ace())
	assert(s, 21, []int{10}, t)
}
func TestHighAceOver21(t *testing.T) {
	s := NewScore()
	s.AddCard(ten())
	s.AddCard(ten())
	s.AddCard(ace())
	assert(s, 21, []int{}, t)
}

func TestTwoAces(t *testing.T) {
	s := NewScore()
	s.AddCard(ace())
	s.AddCard(ace())
	assert(s, 12, []int{10}, t)
}
