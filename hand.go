package main

type Hand struct {
	Cards []Card
	Score Score
}

func NewHand() Hand {
	return Hand{[]Card{}, NewScore()}
}

func (h *Hand) AddCard(c Card) {
	h.Score.AddCard(c)
	h.Cards = append(h.Cards, c)
}

func (h Hand) Copy() Hand {
	return h
}