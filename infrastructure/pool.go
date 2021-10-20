package infrastructure

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/mannyOaks/academy-go-q32021/entities"
)

type MovieWorkerPool struct {
}

func NewMovieWorkerPool() MovieWorkerPool {
	return MovieWorkerPool{}
}

// GetMovies - returns an array of movies read from a csv file concurrently and an error
func (wp MovieWorkerPool) GetMovies(filter string, workerSize, items, maxJobs int) ([]entities.Movie, error) {
	f, err := os.Open(os.Getenv("CSV_PATH"))
	if err != nil {
		return nil, err
	}
	f.Seek(0, 0)

	var wg sync.WaitGroup
	wg.Add(workerSize)

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

			if finishedJobs > maxJobs {
				break
			}
		}
		fmt.Printf("Worker %d finished a job\n", id)
	}

	for w := 0; w < workerSize; w++ {
		go func(id int) {
			taskFunc(id)
		}(w)
	}

	csvFile := csv.NewReader(f)
	go func() {
		for i := 0; i < items; i++ {
			fmt.Println(items, i)
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

	// wait for all goroutines to finish and close the results channel
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// append all the elements from the results channel to an array
	movies := make([]entities.Movie, 0)
	for mov := range resultsChan {
		if mov == nil {
			return nil, fmt.Errorf("error parsing %v", mov)
		}
		movies = append(movies, *mov)
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
