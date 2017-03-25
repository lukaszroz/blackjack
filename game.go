package main

type Game struct {
	Player, dealer Hand
	Dealer         Hand `json:"Dealer"`
	deck           Deck
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

func (g *Game) Hit() {
	g.Player.AddCard(g.deck.Pull())
}

func (g *Game) Stand() {
	for g.dealer.Score.value < 17 {
		g.dealer.AddCard(g.deck.Pull())
	}
	g.Dealer = g.dealer
}
