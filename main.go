package main

import (
	"fmt"
	"math/rand"
	"time"
)

//manual testing
func main() {
	rand.Seed(time.Now().UnixNano())
	d := NewDeck()
	fmt.Println(len(d))
	fmt.Println(d)
	fmt.Println(d.Pull())
	fmt.Println(d.Pull())
	fmt.Println(d.Pull())
	fmt.Println(d.Pull())
	fmt.Println(len(d))
	fmt.Println(d)
}
