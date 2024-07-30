package blackjack

const (
	// Card values
	Ace   = "ace"
	Two   = "two"
	Three = "three"
	Four  = "four"
	Five  = "five"
	Six   = "six"
	Seven = "seven"
	Eight = "eight"
	Nine  = "nine"
	Ten   = "ten"
	Jack  = "jack"
	Queen = "queen"
	King  = "king"

	// Actions
	Split = "P"
	Stand = "S"
	Hit   = "H"
	Win   = "W"
)

// ParseCard returns the integer value of a card following Blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case Ace:
		return 11
	case Two:
		return 2
	case Three:
		return 3
	case Four:
		return 4
	case Five:
		return 5
	case Six:
		return 6
	case Seven:
		return 7
	case Eight:
		return 8
	case Nine:
		return 9
	case Ten, Jack, Queen, King:
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == Ace && card1 == card2 {
		return Split
	}

	value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	total := value1 + value2

	dealerValue := ParseCard(dealerCard)

	if total == 21 {
		if dealerValue == 11 || dealerValue == 10 {
			return Stand
		}
		return Win
	}

	if total >= 17 && total <= 20 {
		return Stand
	}

	if total >= 12 && total <= 16 {
		if dealerValue >= 7 {
			return Hit
		}
		return Stand
	}

	return Hit
}
