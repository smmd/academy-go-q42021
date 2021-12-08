package wpool

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/smmd/academy-go-q42021/model"
)

const FileName = "repository/files/pokedex_data.csv"

type Request struct {
	TypeOfJob      string `json:"type" validate:"enum"`
	NumberOfItems  int `json:"items" validate:"required"`
	ItemsPerWorker int `json:"items_per_workers" validate:"required"`
}

type Response struct {
	Value      interface{}
	Err        error
	JobRequest Request
}

type WorkerHandler struct {}

func NewPokemonWorker() WorkerHandler {
	return WorkerHandler{}
}

func (wh WorkerHandler) PokemonWorkerPool(request Request) Response {
	result := make([]*model.Pokemon, 0)
	errs := make(chan error, 1)
	channelJobs := make(chan []string, request.ItemsPerWorker)
	channelResult := make (chan *model.Pokemon)
	isOdd := request.TypeOfJob == "odd"

	file, err := os.Open(FileName)
	if err != nil {
		return Response{
			Value: nil,
			Err: err,
			JobRequest: request,
		}
	}

	defer file.Close()
	csvFileReader := csv.NewReader(file)

	var wg sync.WaitGroup

	workerCount := request.NumberOfItems / request.ItemsPerWorker
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			worker(channelJobs, channelResult, isOdd, request.ItemsPerWorker)
		}()
	}

	go func() {
		for {
			rStr, err := csvFileReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				errs <- err
				break
			}

			channelJobs <- rStr
		}

		close(channelJobs)
		close(errs)
	}()

	//for _, e := range er

	go func() {
		wg.Wait()
		close(channelResult)
	}()

	for r := range channelResult {
		result = append(result, r)
	}

	return Response{
		Value: result,
		Err: nil,
		JobRequest: request,
	}
}

func worker(channelJobs <-chan []string, channelResult chan<- *model.Pokemon, isOdd bool, limit int)  {
	countItems := 0

	for {
		job, ok := <-channelJobs
		if !ok {
			return
		}

		if countItems == limit {
			return
		}

		pokeId, _ := strconv.Atoi(job[0])
		if isOdd && pokeId%2 != 0 {
			continue
		}

		if !isOdd && pokeId%2 == 0 {
			continue
		}

		channelResult <- parsePokemon(job)
		countItems++
	}
}

func parsePokemon(data []string)  *model.Pokemon {
	return &model.Pokemon{
		ID: data[0],
		Name: data[1],
	}
}
