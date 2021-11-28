package service

import (
	"log"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/smmd/academy-go-q42021/model"
	"github.com/smmd/academy-go-q42021/repository"
)

func ConsumeNationalPokedex(c *gin.Context)  {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/1/")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject model.Response
	json.Unmarshal(responseData, &responseObject)

	err = repository.WritePokeMonsters(responseObject, FileName)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, []string {err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, []string {"ok"})
	}
}
