package main

import (
	"math/rand"
	"time"
)

type Deck []Card

var suits = []string{"Clubs", "Diamonds", "Hearts", "Spades"}

func NewDeck() Deck {
	d := make(Deck, 0, 52)
	for _, s := range suits {
		d = append(d,
			Card{"Ace", s, 11, 1},
			Card{"Two", s, 2, 2},
			Card{"Three", s, 3, 3},
			Card{"Four", s, 4, 4},
			Card{"Five", s, 5, 5},
			Card{"Six", s, 6, 6},
			Card{"Seven", s, 7, 7},
			Card{"Eight", s, 8, 8},
			Card{"Nine", s, 9, 9},
			Card{"Ten", s, 10, 10},
			Card{"Jack", s, 10, 10},
			Card{"Queen", s, 10, 10},
			Card{"King", s, 10, 10},
		)
	}
	return d
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//TODO: empty deck
func (deck *Deck) Pull() Card {
	d := *deck
	i := rand.Intn(len(d))
	last := len(d) - 1
	if i != last {
		d[i], d[last] = d[last], d[i]
	}
	card := d[last]
	*deck = d[:last]
	return card
}