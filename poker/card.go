package poker

import (
	"fmt"
	"strings"
)

// Card A card in a poker game has a suit and a value
type Card struct {
	Value int
	Suit  string
}

const allowedValues = "23456789TJQKA"

const allowedSuits = "CDHS"

func createCard(s string) (c Card, err error) {
	length := len(s)
	if length != 2 {
		return c, fmt.Errorf("invalid card length: expected 2, got %d", length)
	}

	value := string(s[0])
	valueAsInt := strings.Index(allowedValues, value) + 2
	if valueAsInt < 2 || valueAsInt > 14 {
		return c, fmt.Errorf("invalid card value: expected one of %q, got %s", allowedValues, value)
	}

	suit := string(s[1])
	if !strings.Contains(allowedSuits, suit) {
		return c, fmt.Errorf("invalid card suit: expected one of %q, got %s", allowedSuits, suit)
	}

	return Card{valueAsInt, suit}, nil
}
