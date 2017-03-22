package card

import (
	"fmt"
	"strings"
)

// Card a card in a poker game
type Card interface {
	getValue() int
	getSuit() string
}

type card struct {
	value int
	suit  string
}

func (card card) getValue() int {
	return card.value
}

func (card card) getSuit() string {
	return card.suit
}

const allowedValues = "23456789TJQKA"

const allowedSuits = "CDHS"

// CreateCard creates a new card
func CreateCard(s string) (c Card, err error) {
	length := len(s)
	if length != 2 {
		return c, fmt.Errorf("invalid card length: expected 2, got %d", length)
	}

	value := string(s[0])
	index := strings.Index(allowedValues, value)
	if index == -1 {
		return c, fmt.Errorf("invalid card value: expected one of %q, got %s", allowedValues, value)
	}
	valueAsInt := index + 2

	suit := string(s[1])
	if !strings.Contains(allowedSuits, suit) {
		return c, fmt.Errorf("invalid card suit: expected one of %q, got %s", allowedSuits, suit)
	}
	return card{valueAsInt, suit}, nil
}
