package utils

import (
	"fmt"
	"poker-game/types"
)

func PrintFlop(flop []types.Card) {
	flopString := "┌────┐┌────┐┌────┐┌────┐┌────┐\n│%s%s││%s%s││%s%s││%s%s││%s%s│\n|    ||    ||    ||    ||    |\n└────┘└────┘└────┘└────┘└────┘\n\n"

	fmt.Println("\nFlop:")
	fmt.Printf(
		flopString,
		flop[0].Label, flop[0].Suit,
		flop[1].Label, flop[1].Suit,
		flop[2].Label, flop[2].Suit,
		flop[3].Label, flop[3].Suit,
		flop[4].Label, flop[4].Suit,
	)
}
