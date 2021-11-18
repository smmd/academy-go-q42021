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

// TODO: figure out how to test this
/*func TestGettingPokemonMonsters(t *testing.T)  {
	gin.SetMode(gin.TestMode)

	mockPokeMonsters := &PokeMonsters{
		[]*Pokemon{
			{
				ID:   "10",
				Name: "jigglypuff",
			},
			{
				ID:   "150",
				Name: "mewtwo",
			},
		},
	}
}*/
