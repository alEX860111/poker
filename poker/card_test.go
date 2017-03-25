package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCardHeartTwo(t *testing.T) {
	assert := assert.New(t)
	card, err := createCard("2H")
	assert.Equal(2, card.Value)
	assert.Equal("H", card.Suit)
	assert.Nil(err)
}

func TestCreateCardClubAce(t *testing.T) {
	assert := assert.New(t)
	card, err := createCard("AC")
	assert.Equal(14, card.Value)
	assert.Equal("C", card.Suit)
	assert.Nil(err)
}

func TestCreateCardInvalidLength(t *testing.T) {
	assert := assert.New(t)
	_, err := createCard("2HH")
	assert.Equal("invalid card length: expected 2, got 3", err.Error())
}

func TestCreateCardInvalidValue(t *testing.T) {
	assert := assert.New(t)
	_, err := createCard("0H")
	assert.Equal("invalid card value: expected one of \"23456789TJQKA\", got 0", err.Error())
}

func TestCreateCardInvalidSuit(t *testing.T) {
	assert := assert.New(t)
	_, err := createCard("2X")
	assert.Equal("invalid card suit: expected one of \"CDHS\", got X", err.Error())
}
