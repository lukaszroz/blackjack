package main

import "sync"

type Game struct {
	sync.Mutex
	Player, dealer, Dealer          Hand
	deck                            Deck
	ID                              int
	IsFinished, IsTie, HasPlayerWon bool
}

func NewGame() Game {
	g := Game{Player: NewHand(), dealer: NewHand(), deck: NewDeck()}
	g.Player.AddCard(g.deck.Pull())
	g.dealer.AddCard(g.deck.Pull())
	g.Dealer = g.dealer
	g.Player.AddCard(g.deck.Pull())
	g.dealer.AddCard(g.deck.Pull())
	return g
}

// give player another card
// if game is finished, do nothing
func (g *Game) Hit() {
	if g.IsFinished {
		return
	}
	g.Player.AddCard(g.deck.Pull())
	if g.Player.IsBust() {
		g.IsFinished = true
	}
}

const STAND_LIMIT = 17

// resolve dealer hand
// if game is finished, do nothing
func (g *Game) Stand() {
	if g.IsFinished {
		return
	}
	g.IsFinished = true
	for g.dealer.Score.value < STAND_LIMIT {
		g.dealer.AddCard(g.deck.Pull())
	}
	//reveal dealer cards
	g.Dealer = g.dealer
	//determine game outcome
	g.IsTie = g.dealer.Score.value == g.Player.Score.value
	g.HasPlayerWon = g.Player.Score.value > g.dealer.Score.value || g.dealer.IsBust()
}
