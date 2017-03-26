package main

import (
	"errors"
	"sync"
)

type Game struct {
	sync.RWMutex
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

const game_finished = "Game is finished"

// give player another card
// if game is finished, return an error
func (g *Game) Hit() error {
	if g.IsFinished {
		return errors.New(game_finished)
	}
	g.Player.AddCard(g.deck.Pull())
	if g.Player.IsBust() {
		g.IsFinished = true
	}
	return nil
}

const STAND_LIMIT = 17

// resolve dealer hand
// if game is finished, return an error
func (g *Game) Stand() error {
	if g.IsFinished {
		return errors.New(game_finished)
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
	return nil
}
