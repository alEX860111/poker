package card

import (
	"fmt"
	"strings"
)

// Card a card in a poker game
type Card struct {
	Value int
	Suit  string
}

// CreateCard creates a new card
func CreateCard(s string) (c *Card, err error) {
	if len(s) != 2 {
		return c, fmt.Errorf("invalid card: %s", s)
	}
	value := string(s[0])
	valueAsInt, err := convertValueToInt(value)
	if err != nil {
		return c, err
	}
	suit := string(s[1])
	if !strings.Contains("CDHS", suit) {
		return c, fmt.Errorf("invalid card suit: %s", suit)
	}
	return &Card{valueAsInt, suit}, nil
}

func convertValueToInt(value string) (int, error) {
	if len(value) != 1 {
		return -1, fmt.Errorf("invalid card value: %s", value)
	}
	values := "23456789TJQKA"
	index := strings.Index(values, value)
	if index == -1 {
		return -1, fmt.Errorf("invalid card value: %s", value)
	}
	return index + 2, nil
}
