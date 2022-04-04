package game

import (
	"fmt"

	"github.com/pgzisis/deck"
)

type game struct {
	player []deck.Card
	dealer []deck.Card
	cards  []deck.Card
}

func StartGame(cards []deck.Card) {
	g := game{cards: cards}

	for i := 0; i < 2; i++ {
		g.player = append(g.player, g.cards[0])
		g.cards = g.cards[1:]

		g.dealer = append(g.dealer, g.cards[0])
		g.cards = g.cards[1:]
	}

	g.startPlayerTurn()
}

func (g game) startPlayerTurn() {
	score := calculateScore(g.player)

	fmt.Println("Your cards:", g.player)
	fmt.Println("Your score:", score)
	fmt.Printf("Dealer's cards: [%s **HIDDEN CARD**]\n", g.dealer[0].String())

	if score == 21 {
		fmt.Println("Blackjack, you win!")

		return
	} else if score > 21 {
		fmt.Println("You busted!")

		return
	}

	fmt.Println("(h)it or (s)stand?")

	var choice string
	fmt.Scanln(&choice)

	switch {
	case choice == "h":
		g.player = append(g.player, g.cards[0])
		g.cards = g.cards[1:]

		g.startPlayerTurn()
	case choice == "s":
		g.startDealerTurn()
	}
}

func (g game) startDealerTurn() {
	playerScore := calculateScore(g.player)
	dealerScore := calculateScore(g.dealer)

	fmt.Println("Your cards:", g.player)
	fmt.Println("Your score:", playerScore)
	fmt.Println("Dealer's cards:", g.dealer)
	fmt.Println("Dealer's score:", dealerScore)

	switch {
	case dealerScore == 21:
		fmt.Println("Blackjack, house wins!")
	case dealerScore > 21:
		fmt.Println("House busted, you win!")
	case dealerScore < 17:
		g.dealer = append(g.dealer, g.cards[0])
		g.cards = g.cards[1:]

		g.startDealerTurn()
	default:
		switch {
		case playerScore > dealerScore:
			fmt.Println("You win!")
		case playerScore == dealerScore:
			fmt.Println("It's a tie!")
		default:
			fmt.Println("House wins!")
		}
	}

	if dealerScore == 21 {

	} else if dealerScore > 21 {

	}
}

func calculateScore(d []deck.Card) int {
	score := 0

	for _, c := range d {
		switch {
		case c.Rank >= 10:
			score += 10
		case c.Rank == 1:
			if score >= 11 {
				score++
			} else {
				score += 11
			}
		default:
			score += int(c.Rank)
		}
	}

	return score
}
