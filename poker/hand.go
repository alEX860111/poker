package poker

import (
	"fmt"
	"strings"
)

// Hand A poker hand consists of 5 cards
type Hand struct {
	cards [5]Card
}

// CreateHand creates a new hand
func CreateHand(s string) (h Hand, err error) {
	cardStrings := strings.Split(s, " ")
	if len(cardStrings) != 5 {
		return h, fmt.Errorf("invalid hand: %s", s)
	}
	cards := [5]Card{}
	for i, cardString := range cardStrings {
		card, err := createCard(cardString)
		if err != nil {
			return h, err
		}
		cards[i] = card
	}

	_, err = validateCards(cards[:])
	if err != nil {
		return h, err
	}
	return Hand{cards}, err
}
