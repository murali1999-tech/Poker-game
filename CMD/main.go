package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/poker-game/internal/game"
	"github.com/poker-game/internal/hand"
	"github.com/poker-game/internal/player"
	"github.com/poker-game/pkg/utils"
)

func main() {
	// Create two players
	player1 := player.New("Player 1")
	player2 := player.New("Player 2")

	// Create a new game
	pokerGame := game.New(player1, player2)

	// Read the input from STDIN
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Parse the input into a Hand
		input := strings.TrimSpace(scanner.Text())
		h, err := hand.Parse(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing hand: %v\n", err)
			continue
		}

		// Add the Hand to the game
		pokerGame.AddHand(h)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Play the game
	pokerGame.Play()

	// Print the results
	utils.PrintResults(pokerGame.Results())
}

