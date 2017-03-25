package main

import (
	"fmt"
	"math/rand"
	"time"
	"encoding/json"
)

//manual testing
func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame()
	m, _ := json.MarshalIndent(g, "", "  ")
	fmt.Println(string(m))
	for g.Player.Score.value < 18 {
		g.Hit()
	}
	if g.Player.Score.value < 22 {
		g.Stand()
	}
	m, _ = json.MarshalIndent(g, "", "  ")
	fmt.Println(string(m))
}
