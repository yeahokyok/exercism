package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	dealer := ParseCard(dealerCard)
	player := ParseCard(card1) + ParseCard(card2)

	if card1 == "ace" && card2 == "ace" { // Pair of aces
		return "P"
	}

	if player == 21 &&
		!(dealerCard == "ace" || dealerCard == "ten" ||
			dealerCard == "jack" || dealerCard == "queen" ||
			dealerCard == "king") { // Blackjack, dealer does not have Ace/Figure/10
		return "W"
	}

	if player >= 17 { // Value between 17 and 20
		return "S"
	}

	if player <= 16 && player >= 12 { // Value between 12 and 16
		if dealer >= 7 {
			return "H"
		} else {
			return "S"
		}
	}

	if player <= 11 { // Value 11 or less
		return "H"
	}

	// fallback
	return "S"
}
