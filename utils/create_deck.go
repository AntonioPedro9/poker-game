package utils

import "poker-game/types"

func CreateDeck() *types.Deck {
	labels := []string{" 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9", "10", " J", " Q", " K", " A"} // space is for visual purposes
	suits := []string{"♣️", "♥️", "♠️", "♦️"}

	var deck types.Deck

	for _, suit := range suits {
		for i, label := range labels {
			card := types.Card{
				Rank:  i + 2,
				Suit:  suit,
				Label: label,
			}
			deck.Cards = append(deck.Cards, card)
		}
	}

	return &deck
}
