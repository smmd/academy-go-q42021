package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smmd/academy-go-q42021/model"
	"net/http"
)

type search interface {
	GetAll() (model.PokeMonsters, error)
	GetOneByID(id string) (model.Pokemon, error)
}

type pokeapi interface {
	ConsumeNationalPokedex() error
}

type PokemonsHandler struct {
	searchService search
	apiService pokeapi
}

func NewPokemonsHandler(search search, pokeapi pokeapi) PokemonsHandler {
	return PokemonsHandler {
		search,
		pokeapi,
	}
}

func (ph PokemonsHandler) PokeMonsters(c *gin.Context) {
	pokeMonsters, err := ph.searchService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pokeMonsters)
}

func (ph PokemonsHandler) Pokemon(c *gin.Context) {
	id := c.Param("id")
	pokemon, err := ph.searchService.GetOneByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pokemon)
}

func (ph PokemonsHandler) Pokedex(c *gin.Context) {
	err := ph.apiService.ConsumeNationalPokedex()

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	response := make(map[string]string)
	response["message"] = "OK"

	c.JSON(http.StatusOK, response)
}
