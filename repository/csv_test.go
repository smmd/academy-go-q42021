package repository

import (
	"testing"

	"github.com/smmd/academy-go-q42021/model"

	"github.com/stretchr/testify/assert"
)

func TestConvertingCSVDataToModelObj(t *testing.T)  {
	expected := model.PokeMonsters{
		[]model.Pokemon{
			{
				ID:   "1",
				Name: "bulbasaur",
			},
			{
				ID:   "2",
				Name: "ivysaur",
			},
			{
				ID:   "3",
				Name: "venusaur",
			},
		},
	}

	actual, _ := GetAllPokeMonsters("fixtures/pokedex_data.csv")

	assert.Equal(t, actual, expected)
}

func TestThrowingErrorFileNoExist(t *testing.T)  {
	_, actual := GetAllPokeMonsters("fixtures/pokedex_data_fail.csv")

	assert.EqualError(t, actual, "open fixtures/pokedex_data_fail.csv: no such file or directory")
}
