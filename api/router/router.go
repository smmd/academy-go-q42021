package router

import (
	"github.com/smmd/academy-go-q42021/api/service"

	"github.com/gin-gonic/gin"
)

type search interface {
	GetAll(c *gin.Context)
	GetOneByID(c *gin.Context)
}

type PokemonsHandler struct {
	searchService search
}

func NewPokemonsHandler(search search) PokemonsHandler {
	return PokemonsHandler{search}
}

func (ph PokemonsHandler) Route()  {
	router := gin.Default()

	router.GET("/pokemonsters/", ph.searchService.GetAll)
	router.GET("/pokemonsters/:id", ph.searchService.GetOneByID)

	router.GET("/fill-pokedex/", service.ConsumeNationalPokedex)

	router.Run(":3001")
}
