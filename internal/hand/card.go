package game

type Card struct {
	Value int
	Suit  string
}

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

func (c Card) String() string {
	return getValueStr(c.Value) + c.Suit
}

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
		return string(value+'0')
	}
}
