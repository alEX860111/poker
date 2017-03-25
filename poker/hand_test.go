package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHand(t *testing.T) {
	assert := assert.New(t)
	hand, err := CreateHand("2H 3C JS 2S AD")
	assert.NotNil(hand)
	assert.Nil(err)
}

func TestCreateHandInvalidLength(t *testing.T) {
	assert := assert.New(t)
	_, err := CreateHand("2H 3C JS 2SAD")
	assert.Equal("invalid hand: 2H 3C JS 2SAD", err.Error())
}

func TestCreateHandInvalidCards(t *testing.T) {
	assert := assert.New(t)
	_, err := CreateHand("2H 2H JS 2S AD")
	assert.Equal("validation error: each card must be unique", err.Error())
}
