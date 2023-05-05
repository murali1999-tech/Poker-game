package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/murali1999-tech/poker-game/internal/game"
	"github.com/murali1999-tech/poker-game/internal/hand"
	"github.com/murali1999-tech/poker-game/internal/player"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	p1Wins, p2Wins := 0, 0

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		cards := strings.Fields(input)
		if len(cards) != 10 {
			continue // skip invalid input
		}

		p1Cards, p2Cards := cards[:5], cards[5:]
		p1Hand := hand.New(p1Cards)
		p2Hand := hand.New(p2Cards)

		p1 := player.New(1, p1Hand)
		p2 := player.New(2, p2Hand)

		win, _ := game.Play(p1, p2)
		if win == 1 {
			p1Wins++
		} else {
			p2Wins++
		}
	}

	fmt.Printf("Player 1: %d hands\n", p1Wins)
	fmt.Printf("Player 2: %d hands\n", p2Wins)
}
