package player

import (
    "fmt"
    "strings"

    "github.com/username/poker-game/internal/hand"
)

// Player represents a player in the game of poker.
type Player struct {
    Name string       // The name of the player.
    Hand *hand.Hand   // The hand of cards held by the player.
}

// NewPlayer creates and returns a new player with the given name.
func NewPlayer(name string) *Player {
    return &Player{
        Name: name,
        Hand: hand.NewHand(),
    }
}

// String returns the string representation of the player, including their name and their hand.
func (p *Player) String() string {
    var cards []string
    for _, card := range p.Hand.Cards {
        cards = append(cards, card.String())
    }
    return fmt.Sprintf("%s: %s", p.Name, strings.Join(cards, ", "))
}
