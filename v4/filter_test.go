package v4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnOnlyNumbers(t *testing.T) {
	s1 := "61.225.412/0001-14aA"

	result := *GetOnlyNumbers(&s1)
	assert.Equal(t, "61225412000114", result)
}

func TestShouldReturnOnlyNumbersOrSpecial(t *testing.T) {
	s1 := "+55 (21) 98765-4321"

	result := *GetOnlyNumbersOrSpecial(&s1, "+")
	assert.Equal(t, "+5521987654321", result)
}

func TestShouldReturnOnlyNumbersOrSpecial1(t *testing.T) {
	s1 := "+55 (21) 98765-4321"

	result := *GetOnlyNumbersOrSpecial(&s1, "+()")
	assert.Equal(t, "+55(21)987654321", result)
}

func TestShouldReturnNilForNilInput(t *testing.T) {
	var s1 string

	result := *GetOnlyNumbers(&s1)
	assert.Equal(t, s1, result)
}

func TestShouldReturnNilForNilInput1(t *testing.T) {
	var s1 string

	result := *GetOnlyNumbersOrSpecial(&s1, "+")
	assert.Equal(t, s1, result)
}
