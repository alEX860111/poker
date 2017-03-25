package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCardHeartTwo(t *testing.T) {
	assert := assert.New(t)
	card, err := CreateCard("2H")
	assert.Equal(2, card.Value)
	assert.Equal("H", card.Suit)
	assert.Nil(err)
}

func TestCreateCardClubAce(t *testing.T) {
	assert := assert.New(t)
	card, err := CreateCard("AC")
	assert.Equal(14, card.Value)
	assert.Equal("C", card.Suit)
	assert.Nil(err)
}

func TestCreateCardInvalidCardLength(t *testing.T) {
	assert := assert.New(t)
	_, err := CreateCard("2HH")
	assert.Equal("invalid card length: expected 2, got 3", err.Error())
}

func TestCreateCardInvalidCardValue(t *testing.T) {
	assert := assert.New(t)
	_, err := CreateCard("0H")
	assert.Equal("invalid card value: expected one of \"23456789TJQKA\", got 0", err.Error())
}

func TestCreateCardInvalidCardSuit(t *testing.T) {
	assert := assert.New(t)
	_, err := CreateCard("2X")
	assert.Equal("invalid card suit: expected one of \"CDHS\", got X", err.Error())
}
