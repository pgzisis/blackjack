package main

import (
	"github.com/pgzisis/blackjack/game"
	"github.com/pgzisis/deck"
)

func main() {
	d := deck.New(deck.Shuffle)

	game.StartGame(d)
}
