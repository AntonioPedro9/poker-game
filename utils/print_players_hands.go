package utils

import (
	"fmt"
	"poker-game/types"
)

func PrintPlayersHands(players []types.Player) {
	handString := "\n┌────┐┌────┐\n│%s%s││%s%s│\n|    ||    |\n└────┘└────┘\n"

	for i, player := range players {
		fmt.Printf("\nPlayer %d:", i+1)
		fmt.Printf(
			handString,
			player.FirstCard.Label,
			player.FirstCard.Suit,
			player.SecondCard.Label,
			player.SecondCard.Suit,
		)
	}
}
