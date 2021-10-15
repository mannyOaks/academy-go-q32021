package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"sync"

	"github.com/mannyOaks/academy-go-q32021/entities"
)

func (ms MovieService) FindMovies(filter string, items int, itemsPerWorker int) ([]entities.Movie, error) {
	file, err := os.Open(os.Getenv("CSV_PATH"))
	if err != nil {
		return nil, err
	}

	file.Seek(0, 0)

	numWorkers := int(math.Ceil(float64(items) / float64(itemsPerWorker)))

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	rowsChan := make(chan []string)
	resultsChan := make(chan *entities.Movie)

	taskFunc := func(id int) {
		fmt.Printf("Worker %d started a job\n", id)
		defer wg.Done()

		finishedJobs := 0
		for j := range rowsChan {
			mov, err := parseMovie(j)
			if err != nil {
				break
			}

			rem := mov.ID % 2

			if filter == "odd" && rem != 0 {
				resultsChan <- mov
				finishedJobs++
			} else if filter == "even" && rem == 0 {
				resultsChan <- mov
				finishedJobs++
			}

			if finishedJobs > itemsPerWorker {
				break
			}
		}
		fmt.Printf("Worker %d finished a job\n", id)

	}

	for w := 0; w < numWorkers; w++ {
		go func(id int) {
			taskFunc(id)
		}(w)
	}

	csvFile := csv.NewReader(file)
	go func() {
		for {
			row, err := csvFile.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				break
			}

			rowsChan <- row
		}
		close(rowsChan)
	}()

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// var movies []entities.Movie
	movies := make([]entities.Movie, 0)
	for mov := range resultsChan {
		if mov == nil {
			return nil, fmt.Errorf("error parsing %v", mov)
		}
		movies = append(movies, *mov)
	}

	if len(movies) > items {
		movies = movies[:items]
	}

	return movies, nil
}

func parseMovie(row []string) (*entities.Movie, error) {
	id, err := strconv.Atoi(row[0])
	if err != nil {
		return nil, err
	}

	popularity, err := strconv.ParseFloat(row[5], 64)
	if err != nil {
		return nil, err
	}

	adult, err := strconv.ParseBool(row[7])
	if err != nil {
		return nil, err
	}

	movie := entities.Movie{
		ID:          id,
		Title:       row[1],
		Overview:    row[2],
		Language:    row[3],
		Poster:      row[4],
		Popularity:  popularity,
		ReleaseDate: row[6],
		Adult:       adult,
	}
	return &movie, nil
}
