package router

import (
	"github.com/gin-gonic/gin"

	"github.com/smmd/academy-go-q42021/api/service"
	"github.com/smmd/academy-go-q42021/controller"
	"github.com/smmd/academy-go-q42021/repository"
)

func Route()  {
	searchService := service.NewSearchService(repository.NewAllPokeMonsters())
	apiService := service.NewWriteService(repository.NewPokeMonstersWriter())
	apiController := controller.NewPokemonsHandler(searchService, apiService)

	router := gin.Default()
	router.GET("/pokemonsters/", apiController.PokeMonsters)
	router.GET("/pokemonsters/:id", apiController.Pokemon)
	router.GET("/fill-pokedex/", apiController.Pokedex)

	router.Run(":3001")
}
