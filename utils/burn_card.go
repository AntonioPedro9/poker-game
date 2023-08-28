package utils

import "poker-game/types"

func BurnCard(deck *types.Deck) {
	lastIndex := len(deck.Cards) - 1
	deck.Cards = deck.Cards[:lastIndex]
}
