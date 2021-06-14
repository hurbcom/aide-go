package v4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	assert.Equal(t, 1.2, Round(float64(1.2), 2))
	assert.Equal(t, 1.23, Round(float64(1.23), 2))
	assert.Equal(t, 1.24, Round(float64(1.233), 2))
	assert.Equal(t, 1.24, Round(float64(1.237), 2))
	assert.Equal(t, 1234.56, Round(float64(1234.56), 2))
	assert.Equal(t, 1234.567, Round(float64(1234.567), 3))
	assert.Equal(t, 1234.568, Round(float64(1234.5674), 3))
	assert.Equal(t, 1234.568, Round(float64(1234.5678), 3))
}

func TestRandomInt(t *testing.T) {
	a := RandomInt(1, 9999)
	b := RandomInt(1, 9999)
	c := RandomInt(1, 9999)
	d := RandomInt(1, 9999)
	e := RandomInt(1, 9999)

	assert.NotEqual(t, a, b)
	assert.NotEqual(t, a, c)
	assert.NotEqual(t, a, d)
	assert.NotEqual(t, a, e)
	assert.NotEqual(t, b, c)
	assert.NotEqual(t, b, d)
	assert.NotEqual(t, b, e)
	assert.NotEqual(t, c, d)
	assert.NotEqual(t, c, e)
	assert.NotEqual(t, d, e)
}

func TestFill(t *testing.T) {
	a := struct {
		ID      int
		Name    string
		IsAdmin bool
	}{}

	b := struct {
		Name string
	}{}

	b.Name = "Bobby"

	Fill(&a, b)

	assert.Equal(t, "Bobby", a.Name)
}
