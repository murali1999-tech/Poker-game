package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Card represents a playing card
type Card struct {
	Value int
	Suit  string
}

// Hand represents a poker hand
type Hand struct {
	Cards []Card
	Rank  int
}

// getValueStr returns the string representation of a card value
func getValueStr(value int) string {
	switch value {
	case 14:
		return "A"
	case 13:
		return "K"
	case 12:
		return "Q"
	case 11:
		return "J"
	case 10:
		return "T"
	default:
		return string(value + '0')
	}
}

// NewCard creates a new card from a string representation
func NewCard(cardStr string) Card {
	valueStr := cardStr[:len(cardStr)-1]
	switch valueStr {
	case "A":
		return Card{14, cardStr[len(cardStr)-1:]}
	case "K":
		return Card{13, cardStr[len(cardStr)-1:]}
	case "Q":
		return Card{12, cardStr[len(cardStr)-1:]}
	case "J":
		return Card{11, cardStr[len(cardStr)-1:]}
	case "T":
		return Card{10, cardStr[len(cardStr)-1:]}
	default:
		return Card{int(valueStr[0] - '0'), cardStr[len(cardStr)-1:]}
	}
}

// String returns the string representation of a card
func (c Card) String() string {
	return getValueStr(c.Value) + c.Suit
}

// NewHand creates a new hand from a list of card strings
func NewHand(cardsStr []string) Hand {
	cards := make([]Card, len(cardsStr))
	for i, cardStr := range cardsStr {
		cards[i] = NewCard(cardStr)
	}
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value < cards[j].Value
	})
	rank := getHandRank(cards)
	return Hand{cards, rank}
}

// getHandRank returns the rank of a hand
func getHandRank(cards []Card) int {
	if isRoyalFlush(cards) {
		return 10
	}
	if isStraightFlush(cards) {
		return 9
	}
	if isFourOfAKind(cards) {
		return 8
	}
	if isFullHouse(cards) {
		return 7
	}
	if isFlush(cards) {
		return 6
	}
	if isStraight(cards) {
		return 5
	}
	if isThreeOfAKind(cards) {
		return 4
	}
	if isTwoPairs(cards) {
		return 3
	}
	if isPair(cards) {
		return 2
	}
	return 1
}

// isRoyalFlush checks if a hand is a royal flush
func isRoyalFlush(cards []Card) bool {
	return isStraightFlush(cards) && cards[0].Value == 10
}

// isStraightFlush checks if a hand is a straight flush
func isStraightFlush(cards []Card) bool {
	return isFlush(cards) && isStraight(cards)
}

// isFourOfAKind checks if a hand is a four of a kind
func isFourOfAKind(cards []Card) bool {
	if len(cards) < 4 {
		return false
	}
	for i := 0; i <= len(cards)-4; i++ {
		if cards[i].Value == cards[i+1].Value && cards[i+1].Value == cards[i+2].Value && cards[i+2].Value == cards[i+3].Value {
			return true
		}
	}
	return false
}

// isFullHouse checks if a hand is a full house
func isFullHouse(cards []Card) bool {
	return isThreeOfAKind(cards) && isPair(cards)
}

// isFlush checks if a hand is a flush
func isFlush(cards []Card) bool {
	for i := 1; i < len(cards); i++ {
		if cards[i].Suit != cards[0].Suit {
			return false
		}
	}
	return true
}

// isStraight checks if a hand is a straight
func isStraight(cards []Card) bool {
	for i := 1; i < len(cards); i++ {
		if cards[i].Value != cards[i-1].Value+1 {
			return false
		}
	}
	return true
}

// isThreeOfAKind checks if a hand is a three of a kind
func isThreeOfAKind(cards []Card) bool {
	for i := 0; i < len(cards)-2; i++ {
		if cards[i].Value == cards[i+1].Value && cards[i+1].Value == cards[i+2].Value {
			return true
		}
	}
	return false
}

// isTwoPairs checks if a hand is a two pairs
func isTwoPairs(cards []Card) bool {
	pairCount := 0
	for i := 0; i < len(cards)-1; i++ {
		if cards[i].Value == cards[i+1].Value {
			pairCount++
			i++
		}
	}
	return pairCount == 2
}

// isPair checks if a hand is a pair
func isPair(cards []Card) bool {
	for i := 0; i < len(cards)-1; i++ {
		if cards[i].Value == cards[i+1].Value {
			return true
		}
	}
	return false
}

// String returns the string representation of a hand
func (h Hand) String() string {
	cardStrs := make([]string, len(h.Cards))
	for i, card := range h.Cards {
		cardStrs[i] = card.String()
	}
	return strings.Join(cardStrs, " ") + " (" + getRankStr(h.Rank) + ")"
}

// getRankStr returns the string representation of a hand rank
func getRankStr(rank int) string {
	switch rank {
	case 10:
		return "Royal Flush"
	case 9:
		return "Straight Flush"
	case 8:
		return "Four of a Kind"
	case 7:
		return "Full House"
	case 6:
		return "Flush"
	case 5:
		return "Straight"
	case 4:
		return "Three of a Kind"
	case 3:
		return "Two Pairs"
	case 2:
		return "Pair"
	default:
		return "High Card"
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	player1Wins := 0
	player2Wins := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cards := strings.Split(line, " ")

		player1Hand := NewHand(cards[:5])
		player2Hand := NewHand(cards[5:])

		if player1Hand.Rank > player2Hand.Rank {
			player1Wins++
		} else if player2Hand.Rank > player1Hand.Rank {
			player2Wins++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Player 1:", player1Wins, "hands")
	fmt.Println("Player 2:", player2Wins, "hands")
}
