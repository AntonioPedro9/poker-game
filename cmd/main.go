package main

import (
	"bufio"
	"fmt"
	"os"
	"poker-game/types"
	"poker-game/utils"
)

const (
	Players   = 5
	FlopCards = 5
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Press Enter to deal and flop, or Ctrl+C to exit...")

		_, err := reader.ReadString('\n')

		if err != nil {
			os.Exit(0)
		}

		dealAndFlop()
	}
}

func dealAndFlop() {
	deck := utils.CreateDeck()
	players := make([]types.Player, Players)
	flop := make([]types.Card, FlopCards)

	utils.ShuffleDeck(deck)

	// deals the players cards
	for i := 0; i < Players*2; i++ {
		currentPlayer := i % Players

		if i < Players {
			players[currentPlayer].FirstCard = utils.DrawCard(deck) // deals the players first card
		} else {
			players[currentPlayer].SecondCard = utils.DrawCard(deck) // deals the players second card
		}
	}

	utils.PrintPlayersHands(players)

	utils.BurnCard(deck) // burn a card before dealing the flop

	// deals the flop cards
	for i := 0; i < FlopCards; i++ {
		if i == 3 || i == 5 {
			utils.BurnCard(deck) // burn a card before dealing the turn and river
		}
		flop[i] = utils.DrawCard(deck)
	}

	utils.PrintFlop(flop)
}
