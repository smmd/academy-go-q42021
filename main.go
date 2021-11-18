package main

import (
	"github.com/gin-gonic/gin"
	pokereader "github.com/smmd/academy-go-q42021/reader"
)

func main() {
	router := gin.Default()

	pokeMonsters, err := pokereader.GetPokeMonstersFromFile("reader/fixtures/pokedex_data.csv")

	if err != nil {
		panic(err)
	}

	router.GET("/pokemonsters/", pokeMonsters.GetPokeMonsters)
	router.GET("/pokemonsters/:id", pokeMonsters.SearchNameById)

	router.Run(":3001")
}
