package reader

import (
	"os"

	"github.com/smmd/academy-go-q42021/model"

	"encoding/csv"
)

func GetAllPokeMonsters(filePath string) (model.PokeMonsters, error) {
	lines, err := csvToObject(filePath)

	pokemons := []model.Pokemon{}
	pokeMonsters := model.PokeMonsters{pokemons}

	if err != nil {
		return pokeMonsters, err
	}

	for _, line := range lines {
		pokemon := model.Pokemon{line[0], line[1]}

		pokeMonsters.AddPokemon(pokemon)
	}

	return pokeMonsters, nil
}

func csvToObject(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return lines, nil
}
