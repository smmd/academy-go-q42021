package main

import (
	pokerouter "github.com/smmd/academy-go-q42021/api/router"
	"github.com/smmd/academy-go-q42021/api/service"
	"github.com/smmd/academy-go-q42021/repository"
)

func main() {
	searchService := service.NewSearchService(repository.NewAllPokeMonsters())
	pokemonsHandler := pokerouter.NewPokemonsHandler(searchService)

	pokemonsHandler.Route()
}
