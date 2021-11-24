package service

import (
	pokereader "github.com/smmd/academy-go-q42021/repository"

	"github.com/gin-gonic/gin"
	"net/http"
)

const FileName = "repository/files/pokedex_data.csv"

func GetAll(c *gin.Context)  {
	pokeMonsters, err := pokereader.GetAllPokeMonsters(FileName)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, pokeMonsters)
}

func GetOneByID(c *gin.Context) {
	id := c.Param("id")
	pokeMonsters, err := pokereader.GetAllPokeMonsters(FileName)

	if err != nil {
		panic(err)
	}

	for _, poke := range pokeMonsters.Pokemons {
		if poke.ID == id {
			c.IndentedJSON(http.StatusOK, poke)
		}
	}
}
