package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Pokemon struct {
	ID   string `json:"pokemon_id"`
	Name string `json:"name"`
}

type PokeMonsters struct {
	Pokemon []*Pokemon
}

func NewPokemon(id string, name string)  *Pokemon{
	p := new(Pokemon)
	p.ID = id
	p.Name = name

	return p
}

func (p Pokemon) GetId() string {
	return p.ID
}

func (p Pokemon) GetName() string {
	return p.Name
}

func (p *PokeMonsters) SearchNameById(c *gin.Context) {
	id := c.Param("id")

	for _, poke := range p.Pokemon {
		if poke.ID == id {
			c.IndentedJSON(http.StatusOK, poke)

            return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func (p *PokeMonsters) GetPokeMonsters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, p)
}
