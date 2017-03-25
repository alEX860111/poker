package poker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func validateCardsAlwaysFalse([]Card) (bool, error) {
	return false, fmt.Errorf("mock validation error")
}

func validateCardsAlwaysTrue([]Card) (bool, error) {
	return true, nil
}

func TestCreateHand(t *testing.T) {
	assert := assert.New(t)
	_validateCards = validateCardsAlwaysTrue
	hand, err := CreateHand("2H 3C JS 2S AD")
	assert.NotNil(hand)
	assert.Nil(err)
}

func TestCreateHandInvalidLength(t *testing.T) {
	assert := assert.New(t)
	_validateCards = validateCardsAlwaysTrue
	_, err := CreateHand("2H 3C JS 2SAD")
	assert.Equal("invalid hand: 2H 3C JS 2SAD", err.Error())
}

func TestCreateHandCardsValidationError(t *testing.T) {
	assert := assert.New(t)
	_validateCards = validateCardsAlwaysFalse
	_, err := CreateHand("2H 2H JS 2S AD")
	assert.Equal("mock validation error", err.Error())
}
