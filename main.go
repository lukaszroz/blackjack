package main

import (
	"fmt"
	"math/rand"
	"time"
)

//manual testing
func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame()
	fmt.Println(g.player)
	fmt.Println(g.dealer)
	for g.player.score.value < 18 {
		g.Hit()
	}
	if g.player.score.value < 22 {
		g.Stand()
	}
	fmt.Println(g.player)
	fmt.Println(g.dealer)
}
