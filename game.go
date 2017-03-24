package main

type Game struct {
	player, dealer Hand
	deck Deck
}

func NewGame() Game {
	g := Game{NewHand(), NewHand(), NewDeck()}
	g.player.AddCard(g.deck.Pull())
	g.dealer.AddCard(g.deck.Pull())
	g.player.AddCard(g.deck.Pull())
	g.dealer.AddCard(g.deck.Pull())
	return g
}

func (g *Game) Hit() {
	g.player.AddCard(g.deck.Pull())
}

func (g *Game) Stand() {
	for g.dealer.score.value < 17 {
		g.dealer.AddCard(g.deck.Pull())
	}
}