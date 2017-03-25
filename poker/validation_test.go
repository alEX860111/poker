package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCards(t *testing.T) {
	assert := assert.New(t)
	cards := []Card{Card{2, "H"}, Card{3, "H"}, Card{3, "D"}, Card{4, "H"}, Card{5, "C"}}
	res, err := validateCards(cards)
	assert.True(res)
	assert.Nil(err)
}

func TestValidateCardsValidationError(t *testing.T) {
	assert := assert.New(t)
	cards := []Card{Card{2, "H"}, Card{2, "H"}, Card{3, "D"}, Card{4, "H"}, Card{5, "C"}}
	res, err := validateCards(cards)
	assert.False(res)
	assert.Equal("validation error: each card must be unique", err.Error())
}
