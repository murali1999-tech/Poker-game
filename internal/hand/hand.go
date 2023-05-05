package hand

import (
    "fmt"
    "sort"
    "strings"
)

// Hand represents a player's hand of cards.
type Hand struct {
    Cards []Card
}

// NewHand creates a new hand from the given card strings.
func NewHand(cardsStr string) (*Hand, error) {
    cards := strings.Split(cardsStr, " ")
    if len(cards) != 5 {
        return nil, fmt.Errorf("invalid number of cards: %d", len(cards))
    }

    hand := &Hand{}
    for _, cardStr := range cards {
        card, err := NewCard(cardStr)
        if err != nil {
            return nil, fmt.Errorf("invalid card: %s", cardStr)
        }
        hand.Cards = append(hand.Cards, card)
    }

    sort.Slice(hand.Cards, func(i, j int) bool {
        return hand.Cards[i].Rank < hand.Cards[j].Rank
    })

    return hand, nil
}

// String returns a string representation of the hand.
func (h *Hand) String() string {
    cardStrs := make([]string, len(h.Cards))
    for i, card := range h.Cards {
        cardStrs[i] = card.String()
    }
    return strings.Join(cardStrs, " ")
}

// Evaluate evaluates the hand and returns its rank.
func (h *Hand) Evaluate() (Rank, error) {
    if len(h.Cards) != 5 {
        return 0, fmt.Errorf("invalid number of cards: %d", len(h.Cards))
    }

    if h.IsRoyalFlush() {
        return RoyalFlush, nil
    }
    if rank, ok := h.IsStraightFlush(); ok {
        return rank, nil
    }
    if rank, ok := h.IsFourOfAKind(); ok {
        return rank, nil
    }
    if rank, ok := h.IsFullHouse(); ok {
        return rank, nil
    }
    if h.IsFlush() {
        return Flush, nil
    }
    if rank, ok := h.IsStraight(); ok {
        return rank, nil
    }
    if rank, ok := h.IsThreeOfAKind(); ok {
        return rank, nil
    }
    if rank, ok := h.IsTwoPairs(); ok {
        return rank, nil
    }
    if rank, ok := h.IsPair(); ok {
        return rank, nil
    }

    return HighCard, nil
}

// IsRoyalFlush returns true if the hand is a royal flush.
func (h *Hand) IsRoyalFlush() bool {
    return h.IsFlush() && h.Cards[0].Rank == Ten && h.Cards[4].Rank == Ace
}

// IsStraightFlush returns the rank of the straight flush and true if the hand is a straight flush.
func (h *Hand) IsStraightFlush() (Rank, bool) {
    return h.IsStraightOrFlush(StraightFlush)
}

// IsFourOfAKind returns the rank of the four of a kind and true if the hand is four of a kind.
func (h *Hand) IsFourOfAKind() (Rank, bool) {
    if h.Cards[0].Rank == h.Cards[3].Rank || h.Cards[1].Rank == h.Cards[4].Rank {
        return FourOfAKind, true
    }
    return 0, false
}

// IsFullHouse returns the rank of the full house and true if the hand
