package utils

import "poker-game/types"

func DrawCard(deck *types.Deck) types.Card {
	lastIndex := len(deck.Cards) - 1
	card := deck.Cards[lastIndex]
	deck.Cards = deck.Cards[:lastIndex]

	return card
}
