package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPokemon(t *testing.T) {
	poke := NewPokemon("3", "pikachu")
	expected := &Pokemon{
		"3",
		"pikachu",
	}

	actual := poke
	assert.Equal(t, actual, expected)
}
