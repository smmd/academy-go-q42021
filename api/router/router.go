package router

import (
	"github.com/smmd/academy-go-q42021/api/service"

	"github.com/gin-gonic/gin"
)

func Route()  {
	router := gin.Default()

	router.GET("/pokemonsters/", service.GetAll)
	router.GET("/pokemonsters/:id", service.GetOneByID)

	router.Run(":3001")
}
