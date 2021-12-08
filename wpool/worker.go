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

type Conditions struct {
	isOdd bool
	limitPerWorker int
	maxItems int
}

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
	csvError := make(chan error, 1)
	channelJobs := make(chan []string, request.ItemsPerWorker)
	channelResult := make (chan *model.Pokemon)

	conditions := &Conditions {
		request.TypeOfJob == "odd",
		request.ItemsPerWorker,
		request.NumberOfItems,
	}

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

	workersNumber := workerCount(request.NumberOfItems, request.ItemsPerWorker)
	wg.Add(workersNumber)

	for i := 0; i < workersNumber; i++ {
		go func() {
			defer wg.Done()
			worker(channelJobs, channelResult, conditions)
		}()
	}

	go func() {
		for {
			rStr, err := csvFileReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				csvError <- err
				break
			}

			channelJobs <- rStr
		}

		close(csvError)
		close(channelJobs)
	}()

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

func workerCount(numberOfItems int, itemsPerWorker int) int {
	count := numberOfItems / itemsPerWorker

	if numberOfItems % itemsPerWorker > 0 {
		return count + 1
	}

	return count
}

func worker(channelJobs <-chan []string, channelResult chan<- *model.Pokemon, conditions *Conditions) {
	countItems := 0

	for {
		job, ok := <-channelJobs

		if !ok {
			return
		}

		if countItems == conditions.limitPerWorker {
			return
		}

		if conditions.maxItems == 0 {
			return
		}

		pokeId, _ := strconv.Atoi(job[0])
		if conditions.isOdd && pokeId%2 != 0 {
			continue
		}

		if !conditions.isOdd && pokeId%2 == 0 {
			continue
		}

		channelResult <- parsePokemon(job)
		conditions.maxItems--
		countItems++
	}
}

func parsePokemon(data []string)  *model.Pokemon {
	return &model.Pokemon{
		ID: data[0],
		Name: data[1],
	}
}
