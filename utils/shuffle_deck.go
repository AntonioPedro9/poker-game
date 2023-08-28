package utils

import (
	"math/rand"
	"poker-game/types"
)

func ShuffleDeck(deck *types.Deck) {
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}
