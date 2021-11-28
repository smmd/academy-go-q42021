package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smmd/academy-go-q42021/model"
	"net/http"
)

const FileName = "repository/files/pokedex_data.csv"

type getter interface {
	GetAllPokeMonsters(filePath string) (model.PokeMonsters, error)
}

type SearchService struct {
	repo getter
}

func NewSearchService(repo getter) SearchService {
	return SearchService{repo}
}

func (ss SearchService) GetAll(c *gin.Context) {
	pokeMonsters, err := ss.repo.GetAllPokeMonsters(FileName)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, fmt.Errorf(err.Error()))
	} else {
		c.IndentedJSON(http.StatusOK, pokeMonsters)
	}
}

func (ss SearchService) GetOneByID(c *gin.Context) {
	id := c.Param("id")
	pokeMonsters, err := ss.repo.GetAllPokeMonsters(FileName)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, fmt.Errorf(err.Error()))
	} else {
		for _, poke := range pokeMonsters.Pokemons {
			if poke.ID == id {
				c.IndentedJSON(http.StatusOK, poke)
			}
		}
	}
}
