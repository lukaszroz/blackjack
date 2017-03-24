package main

type Hand struct {
	cards []Card
	score Score
}

func NewHand() Hand {
	return Hand{[]Card{}, NewScore()}
}

func (h *Hand) AddCard(c Card) {
	h.score.AddCard(c)
	h.cards = append(h.cards, c)
}