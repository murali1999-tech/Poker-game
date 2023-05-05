package game

import (
	"github.com/poker-game/internal/hand"
	"github.com/poker-game/internal/player"
)

// Game represents a two player poker game
type Game struct {
	players []*player.Player
	hands   []*hand.Hand
	results []*player.Player
}

// New creates a new game with two players
func New(p1, p2 *player.Player) *Game {
	return &Game{
		players: []*player.Player{p1, p2},
		hands:   make([]*hand.Hand, 0),
	}
}

// AddHand adds a hand to the game, alternating between players
func (g *Game) AddHand(h *hand.Hand) {
	g.hands = append(g.hands, h)

	// Alternate between players
	nextPlayer := len(g.hands) % 2
	g.players[nextPlayer].AddHand(h)
}

// Play plays the game and determines the winner
func (g *Game) Play() {
	// Play each hand
	for _, h := range g.hands {
		h.Play()
	}

	// Determine the winner
	p1Score, p2Score := g.players[0].Score(), g.players[1].Score()
	if p1Score > p2Score {
		g.results = []*player.Player{g.players[0]}
	} else if p2Score > p1Score {
		g.results = []*player.Player{g.players[1]}
	} else {
		g.results = []*player.Player{g.players[0], g.players[1]}
	}
}

// Results returns the results of the game
func (g *Game) Results() []*player.Player {
	return g.results
}
