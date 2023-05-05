package hand

// Card represents a standard playing card.
type Card struct {
    Rank string // The rank of the card, e.g. "2", "3", ..., "Q", "K", "A".
    Suit string // The suit of the card, e.g. "Spades", "Hearts", "Diamonds", "Clubs".
}

// NewCard creates and returns a new card with the given rank and suit.
func NewCard(rank, suit string) *Card {
    return &Card{
        Rank: rank,
        Suit: suit,
    }
}

// String returns the string representation of the card, e.g. "2 of Spades".
func (c *Card) String() string {
    return c.Rank + " of " + c.Suit
}
