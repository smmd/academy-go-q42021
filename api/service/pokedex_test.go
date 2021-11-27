package service

import (
	"errors"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smmd/academy-go-q42021/model"
)

var pokemonsters = model.PokeMonsters{
	[]model.Pokemon{
		{
			"3",
			"pikachu",
		},
		{
			"28",
			"sandslash",
		},
		{
			"29",
			"nidoran-f",
		},
		{
			"30",
			"nidorina",
		},
	},
}

type mockCsvRepo struct {
	mock.Mock
}

func (ms mockCsvRepo) GetAllPokeMonsters(filePath string) (model.PokeMonsters, error) {
	arg := ms.Called(filePath)
	return arg.Get(0).(model.PokeMonsters), arg.Error(1)
}

func TestSearchService_GetAll(t *testing.T) {
	testCases := []struct {
		name string
		response model.PokeMonsters
		argument string
		code int
		err error
	}{
		{
			"response pokemons properly",
			pokemonsters,
			"repository/files/pokedex_data.csv",
			http.StatusOK,
			nil,
		},
		{
			"error when emtpy response",
			model.PokeMonsters{ []model.Pokemon{}},
			"repository/files/pokedex_data.csv",
			http.StatusInternalServerError,
			errors.New("test error"),
		},
		//TODO: add more cases
	}

	gin.SetMode(gin.TestMode)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mockCsvRepo{}
			mock.On("GetAllPokeMonsters", tc.argument).Return(tc.response, tc.err)

			service := NewSearchService(mock)

			r := gin.Default()
			r.GET("/pokemonsters/", service.GetAll)

			req, _ := http.NewRequest(http.MethodGet, "/pokemonsters/", nil)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.code, w.Code)
		})
	}
}

//TODO: TestSearchService_GetOneByID
