package poker

import (
	"fmt"
)

func validateCards(cards []Card) (bool, error) {
	m := make(map[Card]int)
	for _, card := range cards {
		m[card] = m[card] + 1
		if m[card] > 1 {
			return false, fmt.Errorf("validation error: each card must be unique")
		}
	}
	return true, nil
}
