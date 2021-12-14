package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("4000000000000000", "Carlos Martel", 10, 2024, 575)
	assert.Equal(t, "invalid credit card number", err.Error())
}
