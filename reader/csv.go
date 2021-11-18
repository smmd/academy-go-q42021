package reader

import (
	"encoding/csv"
	"github.com/smmd/academy-go-q42021/model"
	"os"
)

func GetPokeMonstersFromFile(filePath string) (*model.PokeMonsters, error) {
	file, err := csvToObject(filePath)
	pokeMonsters := make([]*model.Pokemon, 0)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		pokeMonsters = append(pokeMonsters, &model.Pokemon{
			ID:   line[0],
			Name: line[1],
		})
	}

	return &model.PokeMonsters{
		Pokemon: pokeMonsters,
	}, nil
}

func csvToObject(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	return file, nil
}
