package types

type Card struct {
	Rank  int    // numeric value of the card (2...14)
	Suit  string // suit of the card (♣️, ♥️, ♠️, ♦️)
	Label string // readable representation of the card (2...10 and J, Q, K, A)
}
